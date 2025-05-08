package main

import (
	"WeatherAgent/src/global"
	"WeatherAgent/src/routers/auth"
	"WeatherAgent/src/utils/dbService"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//初始化数据库链接
	config, err := global.InitConfig("../config/config.yaml")
	if err != nil {
		panic(err)
	}
	dbService.DbInit(config)
	auth.AuthRouter(r)
	r.Run(":9090")
}
