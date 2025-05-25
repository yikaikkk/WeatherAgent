package main

import (
	"WeatherAgent/src/global"
	"WeatherAgent/src/routers/auth"
	v1 "WeatherAgent/src/routers/v1"
	"WeatherAgent/src/utils/dbService"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//初始化数据库链接
	config, err := global.InitConfig("config/config.yml")
	if err != nil {
		panic(err)
	}
	dbService.DbInit(config)
	dbService.InitRedisClient(config)
	auth.AuthRouter(r)
	v1.Addv1Router(r)
	r.Run(":9090")
}
