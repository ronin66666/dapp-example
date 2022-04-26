package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mars-mec-go/config"
	"mars-mec-go/log"
)

var Db *gorm.DB

func init() {
	cfg := config.GetConfig("./config/config.yml")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.Mysql.Name,
		cfg.Mysql.Password,
		cfg.Mysql.Address,
		cfg.Mysql.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.SFatal(err)
		return
	}
	mysqlDb, err := db.DB()
	if err != nil {
		log.SFatal(err)
		return
	}

	mysqlDb.SetMaxOpenConns(20)
	mysqlDb.SetMaxIdleConns(5)
	Db = db
}
