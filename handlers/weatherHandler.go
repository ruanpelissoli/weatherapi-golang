package handlers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ruanpelissoli/weatherapi-golang/database"
	"github.com/ruanpelissoli/weatherapi-golang/model"
	"github.com/ruanpelissoli/weatherapi-golang/weatherapi"
	"gorm.io/gorm"
)

type WeatherHandler struct {
	DB  *gorm.DB
	rdb *redis.Client
}

func NewWeatherHandler(db *gorm.DB, rdb *redis.Client) *WeatherHandler {
	return &WeatherHandler{
		DB:  db,
		rdb: rdb,
	}
}

func (w WeatherHandler) GetWeatherByCity(city string) (model.Weather, error) {
	weather, err := weatherapi.Call(city)

	if err != nil {
		return model.Weather{}, err
	}

	go w.insertWeather(city, weather)

	return weather, nil
}

func (w WeatherHandler) insertWeather(query string, responseBody model.Weather) {

	wj, _ := json.Marshal(responseBody)
	weatherJson := string(wj)

	w.DB.Create(&database.WeatherEntity{Query: query, Data: weatherJson, CreatedAt: time.Now()})

	// weatherEntity := database.WeatherEntity{Query: query, Data: weatherJson, CreatedAt: time.Now()}

	// c.DB.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "query"}},           // key colume
	// 	DoUpdates: clause.AssignmentColumns([]string{"data"}), // column needed to be updated
	// }).Create(&weatherEntity)

	ctx := context.Background()

	err := w.rdb.Set(ctx, query, weatherJson, 30*time.Minute).Err()
	if err != nil {
		panic(err)
	}
}
