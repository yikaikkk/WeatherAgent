package v1

import "github.com/gin-gonic/gin"

func v1Router(r *gin.Engine) {
	authGroup := r.Group("/v1")
	{
		authGroup.GET("/getWeather")
	}
}
