package handle

type Handle struct {
}

type Global struct {
	Id       string `gorm:"column:id;type:varchar(100); primary key auto_increment"` //ID
	StrValue string `gorm:"column:str_value;type:varchar(100);"`                     //值
	Desc     string `gorm:"column:desc;type:varchar(100);"`                          //描述
}

func (h *Handle) Init() {

}
