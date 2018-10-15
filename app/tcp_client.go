package main

import (
	"math/rand"
	"sync"
	"time"
)

import (
	"github.com/AlexStocks/getty"
	log "github.com/AlexStocks/log4go"
	"github.com/pkg/errors"
)

var (
	reqID uint32
	src   = rand.NewSource(time.Now().UnixNano())
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

////////////////////////////////////////////////////////////////////
// echo client
////////////////////////////////////////////////////////////////////

type TcpClient struct {
	lock        sync.RWMutex
	sessions    []*clientSession
	GettyClient getty.Client
}

func (c *TcpClient) IsAvailable() bool {
	if c.SelectSession() == nil {
		return false
	}

	return true
}

func (c *TcpClient) Close() {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.GettyClient != nil {
		for _, s := range c.sessions {
			log.Info("close client session{%s, last active:%s, request number:%d}",
				s.session.Stat(), s.session.GetActive(), s.reqNum)
			s.session.Close()
		}
		c.GettyClient.Close()
		c.GettyClient = nil
		c.sessions = c.sessions[:0]
	}
}

func (c *TcpClient) SelectSession() getty.Session {
	// get route server session
	c.lock.RLock()
	defer c.lock.RUnlock()
	count := len(c.sessions)
	if count == 0 {
		log.Info("client session array is nil...")
		return nil
	}

	return c.sessions[rand.Int31n(int32(count))].session
}

func (c *TcpClient) AddSession(session getty.Session) {
	log.Debug("add session{%s}", session.Stat())
	if session == nil {
		return
	}

	c.lock.Lock()
	c.sessions = append(c.sessions, &clientSession{session: session})
	c.lock.Unlock()
}

func (c *TcpClient) RemoveSession(session getty.Session) {
	if session == nil {
		return
	}

	c.lock.Lock()

	for i, s := range c.sessions {
		if s.session == session {
			c.sessions = append(c.sessions[:i], c.sessions[i+1:]...)
			log.Debug("delete session{%s}, its index{%d}", session.Stat(), i)
			break
		}
	}
	log.Info("after remove session{%s}, left session number:%d", session.Stat(), len(c.sessions))

	c.lock.Unlock()
}

func (c *TcpClient) UpdateSession(session getty.Session) {
	if session == nil {
		return
	}

	c.lock.Lock()

	for i, s := range c.sessions {
		if s.session == session {
			c.sessions[i].reqNum++
			break
		}
	}

	c.lock.Unlock()
}

func (c *TcpClient) GetClientSession(session getty.Session) (clientSession, error) {
	var (
		err         error
		echoSession clientSession
	)

	c.lock.Lock()

	err = errors.New("no seesion")
	for _, s := range c.sessions {
		if s.session == session {
			echoSession = *s
			err = nil
			break
		}
	}

	c.lock.Unlock()

	return echoSession, err
}

func (c *TcpClient) heartbeat(session getty.Session) {

	var pkg Protocol
	if err := session.WritePkg(&pkg, WRITE_TIME_OUT); err != nil {
		log.Warn("session.WritePkg(session{%s}, pkg{%s}) = error{%v}", session.Stat(), pkg, err)
		session.Close()

		c.RemoveSession(session)
	}
}
