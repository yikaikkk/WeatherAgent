package v1

import (
	"WeatherAgent/src/controller/weather"
	"github.com/gin-gonic/gin"
)

func Addv1Router(r *gin.Engine) {
	authGroup := r.Group("/v1")
	{
		authGroup.GET("/getWeather", weather.WeatherGetHandler)
	}
}
