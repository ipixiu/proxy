package main

import (
	"github.com/AlexStocks/getty"
	"github.com/AlexStocks/goext/log"
)

type (
	empty interface{}
)

var (
	CTcpServer  getty.Server
	CMsgHandler *MsgHandler
	CZkConf     *ZkConf
	Log         gxlog.Logger
	gConf       ConfYaml
	gClientMap  map[string]*TcpClient
)
