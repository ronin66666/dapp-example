package model

type FILReg struct {
	UserEmail       string `json:"userEmail"`       //邮箱
	UserPassword    string `json:"userPassword"`    //密码
	UserPaypassword string `json:"userPaypassword"` //支付密码
}
