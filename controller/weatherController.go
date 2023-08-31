package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/ruanpelissoli/weatherapi-golang/database"
	"github.com/ruanpelissoli/weatherapi-golang/model"
	"github.com/ruanpelissoli/weatherapi-golang/weatherapi"
	"gorm.io/gorm"
)

type Controller struct {
	DB  *gorm.DB
	rdb *redis.Client
}

func NewController(db *gorm.DB, rdb *redis.Client) *Controller {
	return &Controller{db, rdb}
}

// GetWeatherByCity godoc
//
//	@Summary		Get weather by city
//	@Description	get weather by city
//	@Tags			weather
//	@Accept			json
//	@Produce		json
//	@Param			city	path	string true	"City"
//	@Success		200	{object}	model.Weather
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/weather/{city} [get]
func (c *Controller) GetWeatherByCity(ctx *gin.Context) {

	city := ctx.Param("city")

	weather, err := weatherapi.Call(city)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	go insertWeather(city, weather, c)

	ctx.JSON(http.StatusOK, weather)
}

func insertWeather(query string, responseBody model.Weather, c *Controller) {

	w, _ := json.Marshal(responseBody)
	weatherJson := string(w)

	c.DB.Create(&database.WeatherEntity{Query: query, Data: weatherJson, CreatedAt: time.Now()})

	ctx := context.Background()

	err := c.rdb.Set(ctx, query, weatherJson, 30*time.Minute).Err()
	if err != nil {
		panic(err)
	}
}
