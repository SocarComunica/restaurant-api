package api

import (
	"github.com/gin-gonic/gin"
	"restaurant-api/src/api/config/static"
	"restaurant-api/src/api/dependence"
)

func StartApp() error {
	static.GetConfig()

	r := gin.Default()

	mapRoutes(r)

	return r.Run()
}

func mapRoutes(r *gin.Engine) {
	handlers := dependence.NewWire()

	dishesGroup := r.Group("/dishes")
	dishesGroup.POST("", handlers.DishHandler.Handler)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
