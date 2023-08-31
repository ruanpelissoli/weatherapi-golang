package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/ruanpelissoli/weatherapi-golang/controller"
	"github.com/ruanpelissoli/weatherapi-golang/database"
	"github.com/ruanpelissoli/weatherapi-golang/middlewares"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB := database.ConnectDB()
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_connection"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	r := gin.Default()

	c := controller.NewController(DB, rdb)

	v1 := r.Group("/api/v1")
	{
		weather := v1.Group("/weather")
		{
			weather.GET(":city", middlewares.CacheMiddleware(rdb), c.GetWeatherByCity)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":5001")
}
