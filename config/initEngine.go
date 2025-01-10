package config

import (
	"github.com/gin-gonic/gin"
	"price/internal/api"
)

func InitEngine() (router *gin.Engine) {

	router = gin.Default()

	v0Router := router.Group("/api/v0/")

	v0Router.GET("/price", api.GetItems)
	v0Router.POST("/price", api.AddItems)

	return
}
