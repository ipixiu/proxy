package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path"
	"syscall"
)

import (
	"github.com/AlexStocks/getty"
	"github.com/AlexStocks/goext/log"
	"github.com/AlexStocks/goext/time"
	"github.com/AlexStocks/log4go"
)

const (
	APP_CONF_FILE     = "APP_CONF_FILE"
	APP_LOG_CONF_FILE = "APP_LOG_CONF_FILE"
)

var (
	pprofPath = "/debug/pprof/"

	usageStr = `
Usage: log-kafka [options]
Go runtime version %s
Server Options:
    -c, --config <file>              Configuration file path
    -l, --log <file>                 Log configuration file
    -k, --kafka_log <file>           Kafka Log configuration file
    -t, --http_log <file>            Http Log configuration file
Common Options:
    -h, --help                       Show this message
    -v, --version                    Show version
`
)

func initLog(logConf string) {
	Log = gxlog.NewLoggerWithConfFile(logConf)
	Log.SetAsDefaultLogger()
}

func initSignal() {
	var (
		// signal.Notify的ch信道是阻塞的(signal.Notify不会阻塞发送信号), 需要设置缓冲
		signals = make(chan os.Signal, 1)
	)
	// It is not possible to block SIGKILL or syscall.SIGSTOP
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case sig := <-signals:
			Log.Info("get signal %s", sig.String())
			switch sig {
			case syscall.SIGHUP:
				// reload()
			default:
				go gxtime.Future(gConf.Core.FailFastTimeout, func() {
					Log.Warn("app exit now by force...")
					os.Exit(1)
				})
				//close
				CTcpServer.Close()
				CMsgHandler.Stop()
				CZkServerKeeper.Stop()
				log4go.Close()
				Log.Warn("app exit now...")
				Log.Close()
				return
			}
		}
	}
}

func initIDCreater(workID uint64) {
	var err error
	CIDCreater, err = NewIDCreater(workID)
	if err != nil {
		panic(fmt.Sprintf("initIDCreater, NewIDCreater err: %s", err.Error()))
	}
}

func initMsgDealer() {
	CMsgHandler = NewMsgHandler()
}

func initZkServerKeeper() {
	CZkServerKeeper = NewZkServerKeeper()
}

func initTcpServer() {
	addr := net.JoinHostPort(gConf.TcpServer.localIP, gConf.TcpServer.localPort)
	CTcpServer = getty.NewTCPServer(
		getty.WithLocalAddress(addr),
	)

	CTcpServer.RunEventLoop(newServerSession)
}

func main() {
	var (
		showVersion bool
		configFile  string
		err         error
		logConf     string
	)

	SetVersion(Version)

	/////////////////////////////////////////////////
	// conf
	/////////////////////////////////////////////////

	flag.BoolVar(&showVersion, "v", false, "Print version information.")
	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.StringVar(&configFile, "c", "", "Configuration file path.")
	flag.StringVar(&configFile, "config", "", "Configuration file path.")
	flag.StringVar(&logConf, "l", "", "Logger configuration file.")
	flag.StringVar(&logConf, "log", "", "Logger configuration file.")

	flag.Usage = usage
	flag.Parse()

	// Show version and exit
	if showVersion {
		PrintVersion()
		os.Exit(0)
	}

	// config
	if configFile == "" {
		configFile = os.Getenv(APP_CONF_FILE)
		if configFile == "" {
			panic("can not get configFile!")
		}
	}
	if path.Ext(configFile) != ".yml" {
		panic(fmt.Sprintf("application configure file name{%v} suffix must be .yml", configFile))
	}
	gConf, err = LoadConfYaml(configFile)
	if err != nil {
		log.Printf("Load yaml config file error: '%v'", err)
		return
	}
	fmt.Printf("config: %+v\n", gConf)

	if logConf == "" {
		logConf = os.Getenv(APP_LOG_CONF_FILE)
		if logConf == "" {
			panic("can not get logConf!")
		}
	}
	initLog(logConf)

	// 顺序不能变动
	initZkServerKeeper()
	initIDCreater(uint64(gConf.Core.ServerID))
	initMsgDealer()

	//
	initTcpServer()

	// signal
	initSignal()
}
