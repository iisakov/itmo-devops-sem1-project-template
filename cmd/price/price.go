package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"price/config"
	_ "price/docs"
)

// @title           price archive [by_Artisan]
// @version         1.0
// @description     price archive позволяет хранить информацию о ценах на товары.

// @contact.name   Artisan
// @contact.url    http://www.by_artisan.io/support
// @contact.email  by@artisan.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {
	gin.SetMode(gin.ReleaseMode)
	router := config.InitEngine()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,                                                // Разрешает все домены, можно настроить по необходимости
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Разрешенные методы
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Разрешенные заголовки
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Разрешить отправку куки
	}))

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
