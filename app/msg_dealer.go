package main

import (
	"strconv"
	"sync"
	"time"
)

import (
	jerrors "github.com/juju/errors"

	"github.com/AlexStocks/getty"
	"github.com/gomodule/redigo/redis"
	"github.com/goprotobuf/proto"
)

const (
	NOTIFY_CHANNEL_SIZE = 10240
	MSG_CHANNEL_SIZE    = 10240
	MSG_CHANNEL_NUM     = 1024
)

const (
	startTime = time.Parse("2006-01-02 03:04:05 AM", "2018-10-01 00:00:00 AM").Unix()
)

type MsgHandler struct {
	Pi     *redis.Pool
	Xiu    *redis.Pool
	notify chan Notify
	MsgCh  []chan Protocol

	done    chan empty
	wg      sync.WaitGroup
	handNum int
}

func NewMsgHandler() *MsgHandler {
	h := &MsgHandler{
		notify: make(chan Notify, NOTIFY_CHANNEL_SIZE),
	    h.MsgCh : make([]chan Protocol, MSG_CHANNEL_NUM),
	    done : make(chan empty),
	}

	for i := 0; i < MSG_CHANNEL_NUM; i++ {
		h.MsgCh[i] = make(chan Protocol, MSG_CHANNEL_SIZE)
		// h.MsgCh = append(h.MsgCh, make(chan Protocol, MSG_CHANNEL_SIZE))
	}

	h.Pi = &redis.Pool{
		MaxIdle:   gConf.Pi.MaxIdle,
		MaxActive: gConf.Pi.MaxActive,
		Dial: func() (redis.Conn, error) {
			options := redis.DialDatabase(0)
			c, err := redis.Dial("tcp", gConf.Pi.Host, options)
			if nil != err {
				panic(err.Error())
			}
			return c, err
		},
	}
	h.Xiu = &redis.Pool{
		MaxIdle:   gConf.Xiu.MaxIdle,
		MaxActive: gConf.Xiu.MaxActive,
		Dial: func() (redis.Conn, error) {
			options := redis.DialDatabase(0)
			c, err := redis.Dial("tcp", gConf.Xiu.Host, options)
			if nil != err {
				panic(err.Error())
			}
			return c, err
		},
	}

	go h.start()

	return h
}

func (m *MsgHandler) Handle(session getty.Session, pkg *common.Protocol) error {
	//这里不需要保证原子操作
	m.handNum++
	index := m.handNum % len(m.MsgCh)
	m.MsgCh[index] <- pkg

	return nil
}

func (m *MsgHandler) start() {
	Log.Info("msgHandler start...")

	go m.dealNotify()

	for i := 0; i < MSG_CHANNEL_NUM; i++ {
		go m.dealMsg(i)
	}
}

func (m *MsgHandler) dealMsg(index int) {
	m.wg.Add(1)
	defer m.wg.Done()

	for {
		select {
		case <-m.done:
			Log.Warn("grouting dealMsg %d, stopped", index)
			return

		case msg := <-m.MsgCh[index]:
			m.SaveMsg(msg.Data)
		}
	}
}

func (m *MsgHandler) dealNotify() {
	m.wg.Add(1)
	defer m.wg.Done()

	for {
		select {
		case <-m.done:
			log.Warn("notify goroutine exit.")
			return

		case notifyData := <-m.notify:
			CZkServerKeeper.NotifyBroker(*notifyData.UserId, notifyData.GetUserId())
		}

	}
}

func (m *MsgHandler) SaveMsg(msgData []byte) error {
	var (
		err   error
		msgID uint64
	)

	msg := &UnicastMsg{}
	err = proto.Unmarshal(msgData, msg)
	if nil != err {
		Log.Error("SaveMsg, unmarshal err, %s", err.Error())
		return err
	}

	//xiu save
	xiuConn := m.Xiu.Get()
	if nil != xiuConn.Err() {
		Log.Error("get xiu connection error %s", xiuConn.Err())
		return xiuConn.Err()
	}
	xiuKey := pxInfo.MsgDbPrefix + strconv.FormatUint(msg.UserId/pxInfo.MsgDbNum, 10) // msg_db_11
	_, err = xiuConn.Do("HSET", xiuKey, msg.GetMsgID(), msgData)
	if nil != err {
		Log.Error("xiu conn save msg %+v, got err %s", msg, err)
		return jerrors.Trace(err)
	}

	//pi save
	piConn := m.Pi.Get()
	if nil != piConn.Err() {
		Log.Error("get pi connection error %s", piConn.Err())
		return piConn.Err()
	}
	piKey := pxInfo.MsgIDListPrefix + strconv.FormatUint(msg.UserId, 10)
	score := time.Now().Unix() - startTime
	_, err = piConn.Do("ZADD", piKey, score, msgID)
	if nil != err {
		Log.Error("pi conn save msg index err %s", err)
		return err
	}

	// 通知brokers，非阻塞
	notify := Notify(UserId:msg.UserId)
	select {
	case m.notify <- notify:
	default:
		Log.Error("notify channel full")
	}

	return nil
}

func (m *MsgHandler) Service() string {
	return "MsgHandler"
}

func (m *MsgHandler) Version() string {
	return "v1.0"
}

func (m *MsgHandler) Stop() {
	close(m.done)
	m.Pi.Close()
	m.Xiu.Close()

	m.wg.Wait()
	Log.Warn("MsgHandler stop")
}
