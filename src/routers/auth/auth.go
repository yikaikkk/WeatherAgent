package auth

import (
	"WeatherAgent/src/controller/auth"
	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login")
		authGroup.POST("/register", auth.RegisterHandler)
	}
}
