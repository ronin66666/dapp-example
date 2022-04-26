package gateway

import (
	"encoding/json"
	"fmt"
	msgdb "mars-mec-go/examples/common/msg"
	"mars-mec-go/log"
	"mars-mec-go/model"
	"mars-mec-go/network/processor"
	"mars-mec-go/node"
	"mars-mec-go/resjson"
	"mars-mec-go/service"
	"mars-mec-go/sysservice/httpservice"
	"strings"
)

type GatewayService struct {
	service.Service
	jsonProcessor *processor.JsonProcessor
}

func init() {
	webService := &httpservice.HttpService{}
	webService.SetName("MyGateway")
	node.Setup(webService)
	node.Setup(&GatewayService{})
}

func (this *GatewayService) OnInit() error {
	log.SDebug("GatewayService 初始化")

	//获取系统httpservice服务
	httpervice := node.GetService("MyGateway").(*httpservice.HttpService)
	corsHeader := httpservice.NewAllowCORSHeader()
	httpervice.SetAllowCORS(corsHeader)
	//新建并设置路由对象
	httpRouter := httpservice.NewHttpHttpRouter()
	httpervice.SetHttpRouter(httpRouter, this.GetEventHandler())

	httpRouter.GET("/account/getUserEmail/", this.GetUserEmail)
	httpRouter.POST("/account/register", this.Register)  //注册账户
	httpRouter.POST("/account/query", this.QueryAccount) //查询账户

	httpRouter.POST("/card/transfer", this.CardTransfer) //转换游戏资产
	httpRouter.POST("/card/detail", this.CardDetail)     //卡片详情
	return nil
}

func (this *GatewayService) OnConnected(clientid uint64) {
	log.SDebug("client id %d connected\n", clientid)
}

func (this *GatewayService) OnDisconnected(clientid uint64) {
	log.SDebug("client id %d disconnected\n", clientid)
}

// 游戏消息回调
func (this *GatewayService) OnRequest(clientid uint64, msg interface{}) {
	log.SDebug("clientId=", clientid)
	log.SDebug("msg", msg)

	protoMsg := msg.(*msgdb.MESSAGE_REQ)

	log.SDebug(protoMsg.MsgId)
}

// 消息回调
func (this *GatewayService) OnLogin(session *httpservice.HttpSession) {
	fmt.Println("OnLogin")
	in := make(map[string]interface{})
	out := make(map[string]interface{})

	in["account"] = "kries"
	in["passwd"] = "11111"

	this.Call("LoginService.Login", &in, &out)
	fmt.Println(out["reult"])
}

func (this *GatewayService) GetUserEmail(session *httpservice.HttpSession) {
	fmt.Println("测试")

	var out string
	in := "aaaa"

	this.Call("UserService.RPC_OnUserEmail", &in, &out)
	fmt.Println(out)

	session.WriteJsonDone(200, "SUCCESS")
}

// Register 注册账户
func (this *GatewayService) Register(session *httpservice.HttpSession) {
	fmt.Println("gateway...注册")
	resJson := resjson.HttpRes{
		Code: 200,
		Msg:  "SUCCESS",
	}

	var out string

	body := session.GetBody()
	var user model.GameUser
	json.Unmarshal(body, &user)
	//fmt.Println(user)

	this.Call("AccountService.RPC_Register", &user, &out)
	if "SUCCESS" != out {
		log.Error(out)
		if strings.Contains(out, "uni_email") {
			resJson.Code = 499
			resJson.Msg = "email had been registered"
			session.WriteJsonDone(500, resJson)
			return
		}
		resJson.Code = 500
		resJson.Msg = "FAIL"
		session.WriteJsonDone(500, resJson)
		return
	}

	session.WriteJsonDone(200, resJson)
	return
}

// QueryAccount 查询用户资料
func (this *GatewayService) QueryAccount(session *httpservice.HttpSession) {
	fmt.Println("gateway...查询账户")
	resJson := resjson.HttpRes{
		Code: 200,
		Msg:  "SUCCESS",
	}

	//var out string
	body := session.GetBody()
	var user model.GameUser
	json.Unmarshal(body, &user)

	this.Call("AccountService.RPC_USERDATA_QUERY", &user.WltAddr, &user)

	resJson.Data = user
	session.WriteJsonDone(200, resJson)
	return
}

// CardTransfer 转换游戏资产
func (this *GatewayService) CardTransfer(session *httpservice.HttpSession) {
	fmt.Println("gateway...转换游戏资产")

	body := session.GetBody()
	var userCard model.UserCard
	json.Unmarshal(body, &userCard)

	var resJson resjson.HttpRes
	this.Call("CardService.RPC_CardTransfer", &userCard, &resJson)

	session.WriteJsonDone(200, resJson)
}

// CardDetail 卡片详情
func (this *GatewayService) CardDetail(session *httpservice.HttpSession) {
	fmt.Println("gateway...卡片详情")

	body := session.GetBody()
	var userCard model.UserCard
	json.Unmarshal(body, &userCard)

	var resJson resjson.HttpRes
	this.Call("CardService.RPC_CardDetail", &userCard, &resJson)

	session.WriteJsonDone(200, resJson)
}
