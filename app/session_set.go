package main

import (
	"fmt"
	"net"
	"time"
)

import (
	"github.com/AlexStocks/getty"
	"unicast-async/common"
)

func newServerSession(session getty.Session) error {

	//消息类型回调注册
	//reg handler start
	eventHandler := common.NewMessageHandler()
	eventHandler.RegHandler(common.HANDLER_TYPE_UNICAST_MSG, CMsgHandler)
	//...
	//reg handler stop

	var (
		ok      bool
		tcpConn *net.TCPConn
	)

	if tcpConn, ok = session.Conn().(*net.TCPConn); !ok {
		panic(fmt.Sprintf("%s, session.conn{%#v} is not tcp connection\n", session.Stat(), session.Conn()))
	}

	tcpConn.SetNoDelay(gConf.TcpServer.NoDelay)
	tcpConn.SetKeepAlive(gConf.TcpServer.KeepAlive)
	if gConf.TcpServer.KeepAlive {
		d := time.Duration(gConf.TcpServer.KeepAliveDuration) * time.Second
		tcpConn.SetKeepAlivePeriod(d)
	}
	tcpConn.SetReadBuffer(gConf.TcpServer.ReadBufferSize)
	tcpConn.SetWriteBuffer(gConf.TcpServer.WriteBufferSize)

	session.SetName("proxy session")
	session.SetMaxMsgLen(gConf.TcpServer.MaxMsgLen)
	session.SetPkgHandler(common.NewProxyGettyBuffer())

	session.SetEventListener(eventHandler)
	session.SetRQLen(gConf.TcpServer.ReadBufferSize)
	session.SetWQLen(gConf.TcpServer.WriteQueueSize)
	session.SetReadTimeout(time.Duration(gConf.TcpServer.WriteTimeOut) * time.Second)
	session.SetWriteTimeout(time.Duration(gConf.TcpServer.ReadTimeOut) * time.Second)
	session.SetCronPeriod(gConf.TcpServer.SessionCronPeriod)
	session.SetWaitTime(time.Duration(gConf.TcpServer.SessionTimeOut) * time.Second)
	Log.Debug("app accepts new session:%s\n", session.Stat())

	return nil
}

func newClientSession(session getty.Session) error {
	var (
		ok      bool
		tcpConn *net.TCPConn
	)

	if tcpConn, ok = session.Conn().(*net.TCPConn); !ok {
		panic(fmt.Sprintf("%s, session.conn{%#v} is not tcp connection\n", session.Stat(), session.Conn()))
	}

	tcpConn.SetNoDelay(gConf.TcpServer.NoDelay)
	tcpConn.SetKeepAlive(gConf.TcpServer.KeepAlive)
	if gConf.TcpServer.KeepAlive {
		d := time.Duration(gConf.TcpServer.KeepAliveDuration) * time.Second
		tcpConn.SetKeepAlivePeriod(d)
	}
	tcpConn.SetReadBuffer(gConf.TcpServer.ReadBufferSize)
	tcpConn.SetWriteBuffer(gConf.TcpServer.WriteBufferSize)

	session.SetName("proxy session")
	session.SetMaxMsgLen(gConf.TcpServer.MaxMsgLen)
	session.SetPkgHandler(common.NewProxyGettyBuffer())
	session.SetEventListener(common.NewClientMessageHandler())
	session.SetRQLen(gConf.TcpServer.ReadBufferSize)
	session.SetWQLen(gConf.TcpServer.WriteQueueSize)
	session.SetReadTimeout(time.Duration(gConf.TcpServer.WriteTimeOut) * time.Second)
	session.SetWriteTimeout(time.Duration(gConf.TcpServer.ReadTimeOut) * time.Second)
	session.SetCronPeriod(gConf.TcpServer.SessionCronPeriod)
	session.SetWaitTime(time.Duration(gConf.TcpServer.SessionTimeOut) * time.Second)
	Log.Debug("app accepts new session:%s\n", session.Stat())

	return nil
}
