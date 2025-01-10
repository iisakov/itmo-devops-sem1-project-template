package main

import (
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

// @host      0.0.0.0:8080
// @BasePath  /
func main() {
	router := config.InitEngine()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
