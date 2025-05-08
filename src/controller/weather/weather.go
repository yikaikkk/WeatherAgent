package weather

import (
	"WeatherAgent/src/service/weather"
	"github.com/gin-gonic/gin"
)

func WeatherGetHandler(c *gin.Context) {
	res, err := weather.GetWeather()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": res})
}
