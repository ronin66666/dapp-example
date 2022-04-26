package main

import (
	_ "mars-mec-go/examples/admin"
	db "mars-mec-go/examples/db/gorm"
	_ "mars-mec-go/examples/gateway"
	_ "mars-mec-go/examples/login"
	_ "mars-mec-go/examples/user"
	"mars-mec-go/log"
	"mars-mec-go/node"
)

func main() {

	//初始化数据库链接
	err := db.InitMysql()
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	defer db.Close()
	node.Start()

}
