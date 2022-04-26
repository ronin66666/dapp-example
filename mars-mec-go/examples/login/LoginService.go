package login

import (
	"encoding/json"
	"mars-mec-go/examples/login/handle"
	"mars-mec-go/log"
	"mars-mec-go/node"
	"mars-mec-go/service"
)

type LoginService struct {
	service.Service
	handler *handle.Handle
}

type RPCRawTestCallBack struct {
}

func init() {
	node.Setup(&LoginService{})
}

func (this *LoginService) OnInit() error {
	log.SDebug("LoginService 初始化")

	this.handler = &handle.Handle{}
	this.handler.Init()

	this.RegRpcListener(this)
	return nil
}

func (this *LoginService) OnNodeConnected(nodeId int) {
	log.SDebug("node id %d is conntected.\n", nodeId)
}

func (this *LoginService) OnNodeDisconnect(nodeId int) {
	log.SDebug("node id %d is disconntected.\n", nodeId)
}

// 远程调用的登录
func (this *LoginService) RPC_Login(in *[]byte, out *[]byte) error {
	var data map[string]interface{}

	err := json.Unmarshal(*in, &data)
	if err != nil {
		return err
	}

	return nil
}
