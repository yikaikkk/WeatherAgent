package main

import (
	"WeatherAgent/src/routers/auth"
	v1 "WeatherAgent/src/routers/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//初始化数据库链接
	//config, err := global.InitConfig("../config/config.yaml")
	//if err != nil {
	//	panic(err)
	//}
	//dbService.DbInit(config)
	auth.AuthRouter(r)
	v1.Addv1Router(r)
	r.Run(":9090")
}
