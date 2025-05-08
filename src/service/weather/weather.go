package weather

import (
	"WeatherAgent/src/utils/httpService"
	"encoding/json"
	"fmt"
	"time"
)

// GetWeather 实现天气的获取
func GetWeather() {
	weatherAgent := httpService.NewHttpClient("https://api.seniverse.com/v3/weather/now.json?", 10*time.Second)
	res, err := weatherAgent.Get("key=SuNdzfgEMTX_FjaUz&location=beijing&language=zh-Hans&unit=c")
	if err != nil {
		panic(err)
	}
	var result map[string]interface{}
	err = json.Unmarshal(res, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
