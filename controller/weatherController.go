package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/ruanpelissoli/weatherapi-golang/handlers"
	"github.com/ruanpelissoli/weatherapi-golang/middlewares"
	"gorm.io/gorm"
)

type WeatherController struct {
	Controller
	handler *handlers.WeatherHandler
}

func AddWeatherController(db *gorm.DB, rdb *redis.Client, r *gin.Engine) {
	c := &WeatherController{
		Controller: Controller{},
		handler:    handlers.NewWeatherHandler(db, rdb),
	}

	weather := r.Group("/weather")
	{
		weather.GET(":city", middlewares.CacheMiddleware(rdb), c.getWeatherByCity)
	}
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
func (c *WeatherController) getWeatherByCity(ctx *gin.Context) {

	city := ctx.Param("city")

	weather, err := c.handler.GetWeatherByCity(city)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, weather)
}
