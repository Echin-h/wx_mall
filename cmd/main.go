package main

import (
	"fmt"
	"wxapp/config"
	"wxapp/pkg/mysql/db"
	"wxapp/router"
)

func loading() {
	config.LoadConfig()
	db.LoadMysql()
	//db.LoadRedis()
	fmt.Println("loading success...")
}

func main() {
	loading()
	r := router.NewRouter()
	_ = r.Run(config.GetConfig().System.HttpPort)
	fmt.Println("server is running...")
}
