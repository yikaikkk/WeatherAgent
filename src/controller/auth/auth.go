package auth

import (
	"WeatherAgent/src/model"
	"WeatherAgent/src/service/auth"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	err := auth.Register(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}
