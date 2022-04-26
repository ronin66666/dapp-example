package admin

import (
	"fmt"
	"mars-mec-go/log"
	"mars-mec-go/node"
	"mars-mec-go/service"
	"mars-mec-go/sysservice/httpservice"
)

type AdminService struct {
	service.Service
}

// 后台邮件消息协议
type MailData struct {
	Typ    int    `json:"typ"`
	Action string `json:"action"`
	Method string `json:"method"`
	Data   string `json:"data"`
}

type HttpRespone struct {
	Msg string `json:"msg"`
}

func init() {
	webService := &httpservice.HttpService{}
	webService.SetName("MyAdmin")
	node.Setup(webService)
	node.Setup(&AdminService{})
}

func (this *AdminService) OnInit() error {
	log.SDebug("AdminService 初始化")

	//获取系统httpservice服务
	httpervice := node.GetService("MyAdmin").(*httpservice.HttpService)
	corsHeader := httpservice.NewAllowCORSHeader()
	httpervice.SetAllowCORS(corsHeader)
	//新建并设置路由对象
	httpRouter := httpservice.NewHttpHttpRouter()
	httpervice.SetHttpRouter(httpRouter, this.GetEventHandler())

	//GET方法，请求url:http://127.0.0.1:8080/login?nickname=boyce
	//并header中新增key为uid,value为1000的头,则用postman测试返回结果为：
	//head uid:1000, nickname:boyce
	httpRouter.POST("/admin/login", this.OnLogin)

	return nil
}

func (this *AdminService) OnConnected(clientid uint64) {
	log.SDebug("client id %d connected\n", clientid)
}

func (this *AdminService) OnDisconnected(clientid uint64) {
	log.SDebug("client id %d disconnected\n", clientid)
}

// 游戏消息回调
func (this *AdminService) OnLogin(session *httpservice.HttpSession) {
	out := make([]byte, 0)
	in := session.GetBody()

	this.Call("LoginService.RPC_Login", &in, &out)
	session.WriteJsonDone(200, out)
}

func (this *AdminService) RPC_Test(in *string, out *string) error {
	fmt.Println(*in)
	return nil
}
