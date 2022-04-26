package user

import (
	msgdb "mars-mec-go/examples/common/msg"
	"mars-mec-go/log"
	"mars-mec-go/network/processor"
	"mars-mec-go/node"
	"mars-mec-go/service"
)

type UserService struct {
	service.Service
	jsonProcessor *processor.JsonProcessor
}

func init() {
	node.Setup(&UserService{})
}

func (this *UserService) OnInit() error {
	log.SDebug("UserService 初始化")

	return nil
}

func (this *UserService) OnConnected(clientid uint64) {
	log.SDebug("client id %d connected\n", clientid)
}

func (this *UserService) OnDisconnected(clientid uint64) {
	log.SDebug("client id %d disconnected\n", clientid)
}

// 游戏消息回调
func (this *UserService) OnRequest(clientid uint64, msg interface{}) {
	log.SDebug("clientId=", clientid)
	log.SDebug("msg", msg)

	protoMsg := msg.(*msgdb.MESSAGE_REQ)

	log.SDebug(protoMsg.MsgId)
}

// 业务逻辑
func (this *UserService) RPC_OnUserEmail(in *string, out *string) error {
	log.SDebug("in=", *in)

	*out = "dddd"

	return nil
}
