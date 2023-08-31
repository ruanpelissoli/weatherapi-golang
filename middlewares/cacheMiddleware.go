package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/ruanpelissoli/weatherapi-golang/model"
)

func CacheMiddleware(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		city := c.Param("city")

		val, err := rdb.Get(ctx, city).Result()
		if err != nil || val == "" {
			c.Next()
		} else {
			res := model.Weather{}
			json.Unmarshal([]byte(val), &res)
			c.AbortWithStatusJSON(http.StatusOK, res)
			return
		}
	}
}
