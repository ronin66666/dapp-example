package model

type GameUser struct {
	Uid       string `gorm:"uid" json:"uid"`              //用户UID
	WltAddr   string `gorm:"wlt_addr" json:"wltAddr"`     //钱包地址
	Email     string `gorm:"email" json:"email"`          // 邮箱
	Password  string `gorm:"password" json:"password"`    // 密码(只服用)
	InvitAddr string `gorm:"invit_addr" json:"invitAddr"` //邀请码(他人钱包地址)
	FilUid    int    `gorm:"fil_uid" json:"filUid"`       //fil-UID
}
