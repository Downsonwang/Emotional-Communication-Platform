package redis

import (
	"Gin/pkg/initconf"

	"log"

	"github.com/go-redis/redis"
)

var (
	Client      *redis.Client
	Type        string
	hostAndPort string
)

type SliceCmd = redis.SliceCmd
type StringStringMapCmd = redis.StringStringMapCmd

func init() {
	res, err := initconf.Cfg.GetSection("redis")
	if err != nil {
		log.Fatalf("[Pkg-redis] Failed to get section 'redis' ：%v", err)
	}
	Type = res.Key("TYPE").String()
	hostAndPort = res.Key("HOSTANDPORT").String()
	Client = redis.NewClient(&redis.Options{
		Addr:     hostAndPort,
		Password: "",
		DB:       0,
		PoolSize: 100,
	})
}

func Close() {
	_ = Client.Close()
}
