package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
)

const (
	userName = "root"
	//password = "Tt123!@#"
	//ip       = "rm-2vc00bi34120eqkz8uo.mysql.cn-chengdu.rds.aliyuncs.com"
	password = "root"
	ip       = "localhost"
	port     = "3306"
	dbName   = "interstellar_mecha_mec"
	driver   = "mysql"
)

var (
	DB *gorm.DB
)

func Close() {
	DB.Close()
}

func InitMysql() (err error) {
	dsn := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8&parseTime=True&loc=Local"}, "")
	OpenDB, err := gorm.Open(driver, dsn)
	if err != nil {
		return
	}
	DB = OpenDB
	// 全局禁用表名复数， 即将结构体名作为表名
	DB.SingularTable(true) // 如果设置为true,`User`的默认表名为`user` 不加 s ,如不设置true User默认的表为 users,使用`TableName`设置的表名不受影响
	DB.LogMode(true)       // 设置为true之后控制台会输出对应的SQL语句
	return OpenDB.DB().Ping()
}
