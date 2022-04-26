package model

import "time"

type UserCard struct {
	WltAddr   string     `gorm:"wlt_addr" json:"wltAddr"` //钱包ID
	TokenId   int64      `gorm:"token_id" json:"tokenId"` //卡牌TokenId
	CardId    int64      `gorm:"card_id" json:"cardId"`   //卡牌类型ID
	Rarity    int64      `gorm:"rarity" json:"rarity"`    // 稀有度
	Name      string     `gorm:"name" json:"name"`        // 卡牌名称
	Attr1     int64      `gorm:"attr1" json:"attr1"`      //属性值1
	Attr2     int64      `gorm:"attr2" json:"attr2"`      //属性值2
	Attr3     int64      `gorm:"attr3" json:"attr3"`      //属性值3
	CreatedAt *time.Time `gorm:"created_at"`
}
