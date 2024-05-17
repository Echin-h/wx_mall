package main

import (
	"fmt"
	"wxapp/config"
	"wxapp/pkg/log"
	"wxapp/pkg/mysql/db"
	"wxapp/pkg/redis"
	"wxapp/router"
)

func loading() {
	config.LoadConfig()
	db.LoadMysql()
	redis.CacheLoading()
	log.LoadLog()
	fmt.Println("loading success...")
}

func main() {
	loading()
	r := router.NewRouter()
	_ = r.Run(config.GetConfig().System.HttpPort)
	fmt.Println("server is running...")
}
