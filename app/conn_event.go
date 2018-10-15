package main

import (
	"github.com/AlexStocks/getty"
	log "github.com/AlexStocks/log4go"
)

import (
	"errors"
	"sync"
	"time"
)

//网络消息回调
const (
	HANDLER_TYPE_NIL            = 0
	HANDLER_TYPE_HEART_BEAT     = 1000
	HANDLER_TYPE_UNICAST_MSG    = 1001
	HANDLER_TYPE_UNICAST_NOTIFY = 1002
)

//tcp server or client prama
const (
	WRITE_TIME_OUT  = time.Second * 1
	READ_TIME_OUT   = time.Second * 1
	HEART_BEAT_TIME = time.Second * 3
	MAX_SESSION_NUM = 100000
)

//common session
type clientSession struct {
	session getty.Session
	reqNum  uint64
}

//base handler
type PackageHandler interface {
	Handle(getty.Session, *Protocol) error
}

//heart beat msg handler
type HeartbeatHandler struct{}

func (h *HeartbeatHandler) Handle(session getty.Session, pkg *Protocol) error {
	if pkg.Head.Type == HANDLER_TYPE_HEART_BEAT {
		log.Error("HeartbeatHandler err, type = %d", pkg.Head.Type)
		return errors.New("type err")
	}

	log.Debug("HeartbeatHandler, get hearbeat from %s", session.RemoteAddr())
	//return session.WritePkg(pkg, WRITE_TIME_OUT)
	return nil
}

//tcp server event handler
type MessageHandler struct {
	handlers map[uint16]PackageHandler
	rwlock   sync.RWMutex
	//sessionMap map[getty.Session]*clientSession
}

func NewMessageHandler() *MessageHandler {
	handlers := make(map[uint16]PackageHandler)

	//注册回调,心跳统一处理
	handlers[HANDLER_TYPE_HEART_BEAT] = &HeartbeatHandler{}

	//return &MessageHandler{sessionMap: make(map[getty.Session]*clientSession),
	//handlers: handlers}
	return nil
}

func (h *MessageHandler) RegHandler(hType uint16, handler PackageHandler) error {
	h.rwlock.RLock()
	defer h.rwlock.Unlock()
	h.handlers[hType] = handler
	return nil
}

func (h *MessageHandler) OnOpen(session getty.Session) error {
	clientKey := session.RemoteAddr()
	if v, ok := gClientMap[clientKey]; ok {
		v.AddSession(session)
		return nil
	} else {
		log.Error("session OnOpen, not found tcpclient, addr: %s", clientKey)
	}

	return nil
}

func (h *MessageHandler) OnError(session getty.Session, err error) {
	log.Info("session{%s} got error{%v}, will be closed.", session.Stat(), err)
	clientKey := session.RemoteAddr()
	if v, ok := gClientMap[clientKey]; ok {
		v.RemoveSession(session)
	} else {
		log.Error("session OnError, not found tcpclient, addr: %s", clientKey)
	}
}

func (h *MessageHandler) OnClose(session getty.Session) {
	log.Info("session{%s} is closing......", session.Stat())
	clientKey := session.RemoteAddr()
	if v, ok := gClientMap[clientKey]; ok {
		v.RemoveSession(session)
	} else {
		log.Error("session OnClose, not found tcpclient, addr: %s", clientKey)
	}
}

func (h *MessageHandler) OnMessage(session getty.Session, pkg interface{}) {
	p, ok := pkg.(*Protocol)
	if !ok {
		log.Error("illegal packge{%#v}", pkg)
		return
	}

	handler, ok := h.handlers[p.Head.Type]
	if !ok {
		log.Error("illegal command{%d}", p.Head.Type)
		return
	}

	handler.Handle(session, p)
}

func (h *MessageHandler) OnCron(session getty.Session) {
	var (
		active time.Time
	)

	clientKey := session.RemoteAddr()
	if _, ok := gClientMap[clientKey]; ok {
		active = session.GetActive()
		if (HEART_BEAT_TIME * 3).Nanoseconds() < time.Since(active).Nanoseconds() {
			session.Close()
			log.Error("session=%s timeout=%s, addr=%s",
				session.Stat(), time.Since(active).String(), clientKey)
		}
	} else {
		log.Error("session OnClose, not found tcpclient, addr: %s", clientKey)
	}

}

//tcp client event handler
type ClientHeartbeatHandler struct{}

func (h *ClientHeartbeatHandler) Handle(session getty.Session, pkg *Protocol) error {
	if pkg.Head.Type == HANDLER_TYPE_HEART_BEAT {
		log.Error("ClientHeartbeatHandler, handler err, type = %d", pkg.Head.Type)
		return errors.New("type err")
	}

	log.Debug("*ClientHeartbeatHandler, get heart from %s", session.RemoteAddr())
	//return session.WritePkg(pkg, WRITE_TIME_OUT)
	return nil
}

type ClientMessageHandler struct {
	handlers   map[uint16]PackageHandler
	rwlock     sync.RWMutex
	sessionMap map[getty.Session]*clientSession
}

func NewClientMessageHandler() *ClientMessageHandler {
	handlers := make(map[uint16]PackageHandler)
	handlers[HANDLER_TYPE_HEART_BEAT] = &HeartbeatHandler{}

	return &ClientMessageHandler{sessionMap: make(map[getty.Session]*clientSession), handlers: handlers}
}

func (c ClientMessageHandler) OnOpen(session getty.Session) error {
	log.Info("got session:%s", session.Stat())
	c.rwlock.Lock()
	c.sessionMap[session] = &clientSession{session: session}
	c.rwlock.Unlock()
	return nil
}

func (c ClientMessageHandler) OnError(session getty.Session, err error) {
	log.Info("session{%s} got error{%v}, will be closed.", session.Stat(), err)
	c.rwlock.Lock()
	delete(c.sessionMap, session)
	c.rwlock.Unlock()
}

func (c ClientMessageHandler) OnClose(session getty.Session) {
	log.Info("session{%s} is closing......", session.Stat())
	c.rwlock.Lock()
	delete(c.sessionMap, session)
	c.rwlock.Unlock()
}

func (c ClientMessageHandler) OnMessage(session getty.Session, pkg interface{}) {
	p, ok := pkg.(*Protocol)
	if !ok {
		log.Error("illegal packge{%#v}", pkg)
		return
	}

	handler, ok := c.handlers[p.Head.Type]
	if !ok {
		log.Error("illegal command{%d}", p.Head.Type)
		return
	}

	handler.Handle(session, p)

	c.rwlock.Lock()
	if _, ok := c.sessionMap[session]; ok {
		c.sessionMap[session].reqNum++
		session.UpdateActive()
	}
	c.rwlock.Unlock()
}

func (c ClientMessageHandler) OnCron(session getty.Session) {
	var (
		flag   bool
		active time.Time
	)
	c.rwlock.RLock()
	if _, ok := c.sessionMap[session]; ok {
		active = session.GetActive()
		if (HEART_BEAT_TIME * 3).Nanoseconds() < time.Since(active).Nanoseconds() {
			flag = true
			log.Warn("session{%s} timeout{%s}, reqNum{%d}",
				session.Stat(), time.Since(active).String(), c.sessionMap[session].reqNum)
		}
	}
	c.rwlock.RUnlock()
	if flag {
		c.rwlock.Lock()
		delete(c.sessionMap, session)
		c.rwlock.Unlock()
		session.Close()
	}
}
