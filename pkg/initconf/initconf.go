package initconf

import (
	"gopkg.in/ini.v1"
	"log"
	"time"
)

var (
	Cfg          *ini.File

	Port         int
	PageSize     int

	RunMode      string
	JwtSecret    string

	ReadTimeout  time.Duration
	WriteTimeout time.Duration

)

func init() {
	var err error

	Cfg, err = ini.Load("conf/app.ini")

	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

// 读取ini配置文件[]
func LoadBase() {

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	res, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("[LoadServer] Failed to get section 'server':%v", err)
	}
	Port = res.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(res.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(res.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	res, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("[LoadApp] Failed to get section 'app':%v", err)
	}
	JwtSecret = res.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = res.Key("PAGE_SIZE").MustInt(10)

}
