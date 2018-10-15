package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

import (
	"github.com/goprotobuf/proto"
	"github.com/ipixiu/async-unicast/common"
	"github.com/samuel/go-zookeeper/zk"

	"github.com/AlexStocks/getty"
	"github.com/AlexStocks/goext/database/zookeeper"
	"github.com/AlexStocks/goext/net"
	jerrors "github.com/juju/errors"
)

type ServerInfo struct {
	ip          string `json:"ip""`
	notify_port uint16 `json:"port_msg_notify""`
	ack_port    uint16 `json:"ack_port""`
	flag        uint16 `json:"flag""`
}

func (s *ServerInfo) GetKey() string {
	key := s.ip + strconv.FormatUint(uint64(s.notify_port), 10)
	return key
}

type ServerGroup struct {
	groupID int
	infos   map[string]ServerInfo
}

//////////////////////////////////
// zk keeper
//////////////////////////////////

type ZkConf struct {
	client         *gxzookeeper.Client
	ServerGroups   map[int]ServerGroup
	rwLock         sync.RWMutex
	BrokerGroupNum int

	done chan empty
	wg   sync.WaitGroup
}

func NewZkConf() *ZkConf {
	z := &ZkConf{}
	zkServers := strings.Split(gConf.Zk.Servers, ",")
	zkConn, _, err := zk.Connect(zkServers, time.Second)
	if nil != err {
		panic(fmt.Sprintf("zk.Connect(%#v) = %s", zkServers, err))
	}
	z.client = gxzookeeper.NewClient(zkConn)
	go z.start()
	return z
}

func (z *ZkConf) start() {
	//wg
	z.wg.Add(1)
	defer z.wg.Done()

	//拉取brokers，创建链接；开启该routine，监测brokers变化,动态更新链接
	err := z.LoadServers()
	if nil != err {
		panic(fmt.Sprintf("ZkConf,start GetChildren err: %s", jerrors.ErrorString(err)))
	}

	for {
		z.LoadServers()
		time.Sleep(1e6 * gConf.Zk.IdleMs)
	}
}

func (z *ZkConf) LoadServers() error {
	groups, err := z.client.GetChildren(gConf.Zk.Root)
	if err != nil {
		Log.Error("client.GetChildren(root:%s) = %s", gConf.Zk.Root, jerrors.ErrorString(err))
		return jerrors.Trace(err)
	}

	groupNumStr, err := z.client.Get(gConf.Zk.GroupNumKey)
	if err != nil {
		Log.Error("client.Get(%s) = %s", gConf.Zk.GroupNumKey, jerrors.ErrorString(err))
		return jerrors.Trace(err)
	}
	groupNum, err := strconv.Atoi(string(groupNumStr))
	if err != nil {
		Log.Error("atoi groupNUm err: %s", err.Error())
		return err
	}
	z.setGroupNum(groupNum)

	for _, v := range groups {
		groupID, err := z.GetGroupID(v)
		if err != nil {
			Log.Error("GetGroupID(%#v) = %s", v, jerrors.ErrorString(err))
			continue
		}
		groupPath := gConf.Zk.root + "/" + v
		if gConf.Zk.Root[len(gConf.Zk.Root)-1] == '/' {
			groupPath = gConf.Zk.root + v
		}
		serverNodes, err := z.client.GetChildren(groupPath)
		if err != nil {
			Log.Error("client.GetChildren(path:%s) = %s", groupPath, jerrors.ErrorString(err))
			continue
		}

		nodesMap := make(map[string]ServerInfo)
		for _, node := range serverNodes {
			nodePath := groupPath + "/" + node
			value, err := z.client.Get(nodePath)
			if nil != err {
				Log.Error("client.Get(path:%s) = %s", nodePath, err)
				continue
			}
			if len(value) < 6 {
				continue
			}

			var node ServerInfo
			err = json.Unmarshal([]byte(value), node)
			if err != nil {
				Log.Error("json.Unmarshal(value:%s) = %s", value, err)
				continue
			}
			nodesMap[node.GetKey()] = node
		}
		z.updateNodesMap(groupID, nodesMap)
	}

	return nil
}

func (z *ZkConf) updateNodesMap(groupID int, nodesMap map[string]ServerInfo) {
	//先加读锁判断是否有变化，有变化再处理
	changed := false
	z.rwLock.RLock()
	if len(nodesMap) == 0 {
		z.rwLock.RUnlock()
		Log.Error("groupID = %d, @nodesMap = %#v", groupID, nodesMap)
		return
	}
	if group, ok := z.ServerGroups[groupID]; ok {
		if len(group.infos) != len(nodesMap) {
			changed = true
		} else {
			for nodeKey, _ := range nodesMap {
				if _, ok := group.infos[nodeKey]; !ok {
					changed = true
					break
				}
			}
		}
	} else {
		changed = false
	}
	z.rwLock.RUnlock()

	//修改本地 server info
	if group, ok := z.ServerGroups[groupID]; ok {
		if changed == true {
			//添加
			for nodeKey, nodeV := range nodesMap {
				if _, ok := group.infos[nodeKey]; !ok {
					Log.Warn("get new brokers %+v", nodeV)

					if _, ok := common.gClientMap[nodeKey]; !ok {
						tcpClient := &common.TcpClient{}
						tcpClient.GettyClient = getty.NewTCPClient(
							getty.WithServerAddress(gxnet.HostAddress(nodeV.ip, int(nodeV.notify_port))),
							getty.WithConnectionNumber(1),
						)
						common.gClientMap[nodeKey] = tcpClient
						tcpClient.GettyClient.RunEventLoop(newClientSession)
					}
				}
			}
			//删除
			for nodeKey, nodeV := range group.infos {
				if _, ok := nodesMap[nodeKey]; !ok {
					Log.Warn("brokers close %+v", nodeV)
					if client, ok := common.gClientMap[nodeKey]; !ok {
						client.Close()
						delete(common.gClientMap, nodeKey)
					}
				}

				delete(group.infos, nodeKey)
			}
		}
	}
}

func (z *ZkConf) setGroupNum(groupNum int) {
	changed := false
	z.rwLock.RLock()
	if z.BrokerGroupNum != groupNum {
		changed = true
	}
	z.rwLock.RUnlock()

	if true == changed {
		z.rwLock.Lock()
		z.BrokerGroupNum = groupNum
		z.rwLock.Unlock()
	}
}

func (z *ZkConf) GetGroupID(groupNode string) (int, error) {
	prefixLen := len(gConf.Zk.groupPrefix)
	if prefixLen >= len(groupNode) {
		return -1, jerrors.New("unexpected groupNode")
	}

	groupIDstr := groupNode[prefixLen:]
	id, err := strconv.Atoi(groupIDstr)
	if err != nil {
		return -1, err
	}

	return id, err
}

func (z *ZkConf) NotifyBroker(userID uint64, hash uint64) error {
	z.rwLock.RLock()
	defer z.rwLock.RUnlock()
	hashID := hash % uint64(z.BrokerGroupNum)
	for gID, group := range z.ServerGroups {
		if uint64(gID) != hashID {
			continue
		}
		if len(group.infos) == 0 {
			Log.Error("group %d empty", gID)
			continue
		}
		index := int(hash % uint64(len(group.infos)))
		i := 0
		flag := false
		for nodeKey, _ := range group.infos {
			if i == index {
				if client, ok := common.gClientMap[nodeKey]; ok {
					notifyProto := unicast.UnicasrNotify{}
					notifyProto.UserId = &userID
					var pkg common.Protocol
					pkg.Head.Type = common.HANDLER_TYPE_UNICAST_NOTIFY
					proto.Unmarshal(pkg.Data, &notifyProto)
					pkg.Head.Len = uint16(common.PROTOCAL_HEAD_LEN + len(pkg.Data))

					session := client.SelectSession()
					if session != nil {
						session.WritePkg(pkg, common.WRITE_TIME_OUT)
						flag = true
					} else {
						Log.Error("client %s has not valid session", nodeKey)
					}

				}
				break
			}
			i++
			if i > index {
				break
			}
		}
		if flag == false {
			for nodeKey, _ := range group.infos {
				if client, ok := common.gClientMap[nodeKey]; ok {
					notifyProto := unicast.UnicasrNotify{}
					notifyProto.UserId = &userID
					var pkg common.Protocol
					pkg.Head.Type = common.HANDLER_TYPE_UNICAST_NOTIFY
					proto.Unmarshal(pkg.Data, &notifyProto)
					pkg.Head.Len = uint16(common.PROTOCAL_HEAD_LEN + len(pkg.Data))

					session := client.SelectSession()
					if session != nil {
						session.WritePkg(pkg, common.WRITE_TIME_OUT)
						flag = true
						break
					}
				}

			}
		}

		if flag == false {
			Log.Error("group %d has no valid broker", gID)
		}

	}

	return nil
}

func (z *ZkConf) Stop() {
	close(z.done)
	z.wg.Wait()
	Log.Warn("ZkConf stop")
}
