package weather

import (
	"WeatherAgent/src/utils/dbService"
	"WeatherAgent/src/utils/httpService"
	"context"
	"encoding/json"
	"errors"
	"time"
)

// GetWeather 实现天气的获取
func GetWeather() (map[string]interface{}, error) {
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
	//存入redis
	results := result["results"].([]interface{})
	firstResult := results[0].(map[string]interface{})
	lastUpdate := firstResult["last_update"].(string)
	city := firstResult["location"].(map[string]interface{})["name"].(string)
	weatherKey := lastUpdate + city
	ctx := context.Background()
	weatherData, err := json.Marshal(result)
	if err != nil {
		errors.New("redis Marshal error")
		return nil, err
	}
	_, err = dbService.GetRedisHelper().Set(ctx, weatherKey, weatherData, 10*time.Minute).Result()
	if err != nil {
		errors.New("redis set error")
		return nil, err
	}
	return result, err
}
