package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruanpelissoli/weatherapi-golang/handlers"
)

type RainController struct {
	Controller
	handler *handlers.RainHandler
}

func AddRainController(r *gin.Engine) {
	c := &RainController{
		Controller: Controller{},
		handler:    handlers.NewRainHandler(),
	}

	weather := r.Group("/rain")
	{
		weather.GET("/", c.IsRaining)
	}
}

// IsRaining godoc
//
//	@Summary		Get rain status
//	@Description	get rain status
//	@Tags			rain
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/rain	[get]
func (c *RainController) IsRaining(ctx *gin.Context) {

	city := ctx.Param("city")

	weather, err := c.handler.IsRaining(city)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, weather)
}
