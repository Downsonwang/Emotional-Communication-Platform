package db

import (
	"Gin/pkg/initconf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var (
	Db          *gorm.DB
	err         error
	dbType      string
	dbName      string
	username    string
	password    string
	hostAndPort string
	tablePrefix string
)

// 初始化 gorm mysql .
func init() {
	res, err := initconf.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("[Pkg-db] Failed to get section 'database' : %v", err)
	}
	dbType = res.Key("TYPE").String()
	fmt.Println("datype", dbType)
	dbName = res.Key("NAME").String()
	username = res.Key("USERNAME").String()
	password = res.Key("PASSWORD").String()
	hostAndPort = res.Key("HOSTANDPORT").String()
	tablePrefix = res.Key("TABLE_PREFIX").String()
	//gorm
	Db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", username, password, hostAndPort, dbName))
	if err != nil {
		log.Fatalf("Failed to connect to mysql : %v", err)
	}
	gorm.DefaultTableNameHandler = func(Db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	// 禁用复数xings
	Db.SingularTable(true)
	Db.LogMode(true)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxIdleConns(100)
}

// 关闭Db通道
func CloseDB() {
	defer Db.Close()
}
