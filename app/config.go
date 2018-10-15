package main

import (
	"io/ioutil"
)

import (
	"gopkg.in/yaml.v2"
)

// ConfYaml is config structure.
type ConfYaml struct {
	Core      SectionCore      `yaml:"core" json:"core,omitempty"`
	Pi        SectionPi        `yaml:"pi" json:"pi,omitempty"`
	Xiu       SectionXiu       `yaml:"xiu" json:"xiu,omitempty"`
	Zk        SectionZk        `yaml:"zk" json:"zk,omitempty"`
	TcpServer SectionTcpServer `yaml:"tcpserver" json:"tcp_server,omitempty"`
}

// SectionCore is sub section of config.
type SectionCore struct {
	ServerID        uint16 `yaml:"server_int_id" json:"server_id,omitempty"`
	FailFastTimeout int    `yaml:"fail_fast_time_out" json:"fail_fast_timeout,omitempty"`
}

type SectionTcpServer struct {
	SessionNum        int    `yaml:"session_num" json:"session_num,omitempty"`
	localIP           string `yaml:"local_ip" json:"local_ip,omitempty"`
	localPort         string `yaml:"local_port" json:"local_port,omitempty"`
	NoDelay           bool   `yaml:"no_delay" json:"no_delay,omitempty"`
	KeepAlive         bool   `yaml:"keep_alive" json:"keep_alive,omitempty"`
	KeepAliveDuration int    `yaml:"keep_alive_duration" json:"keep_alive_duration,omitempty"`
	ReadBufferSize    int    `yaml:"read_buffer_size" json:"read_buffer_size,omitempty"`
	WriteBufferSize   int    `yaml:"write_buffer_size" json:"write_buffer_size,omitempty"`
	MaxMsgLen         int    `yaml:"max_msg_len" json:"max_msg_len,omitempty"`
	ReadQueueSize     int    `yaml:"read_queue_size" json:"read_queue_size,omitempty"`
	WriteQueueSize    int    `yaml:"write_queue_size" json:"write_queue_size,omitempty"`
	ReadTimeOut       int    `yaml:"read_time_out" json:"read_time_out,omitempty"`
	WriteTimeOut      int    `yaml:"write_time_out" json:"write_time_out,omitempty"`
	SessionCronPeriod int    `yaml: "session_cron_period"`
	SessionTimeOut    int    `yaml:"session_time_out" json:"session_time_out,omitempty"`
}

//zk
type SectionZk struct {
	Servers       string `json:"servers,omitempty"`
	Root          string `json:"root,omitempty"`
	GroupPrefix   string `json:"group_prefix,omitempty"`
	NodePrefix    string `json:"node_prefix,omitempty"`
	GroupNumKey   string `json:"group_num_key,omitempty"`
	PixiuInfoPath string `json:"pixiu_info_path,omitempty"`
	IdleMs        int    `json:"idle_ms,omitempty"`

	MsgDbPrefix     string `json:"msg_db_prefix,omitempty"`
	MsgDbNum        int    `json:"msg_db_num,omitempty"`
	MsgIDListPrefix string `json:"msg_id_list_prefix,omitempty"`
}

//Pi
type SectionPi struct {
	Host      string `json:"host,omitempty"`
	MaxIdle   int    `json:"max_idle,omitempty"`
	MaxActive int    `json:"max_active,omitempty"`
}

//Xiu
type SectionXiu struct {
	Host      string `json:"host,omitempty"`
	MaxIdle   int    `json:"max_idle,omitempty"`
	MaxActive int    `json:"max_active,omitempty"`
}

// LoadConfYaml provide load yml config.
func LoadConfYaml(confPath string) (ConfYaml, error) {
	var config ConfYaml

	configFile, err := ioutil.ReadFile(confPath)

	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(configFile, &config)

	if err != nil {
		return config, err
	}

	return config, nil
}
