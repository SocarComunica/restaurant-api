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
	dishesGroup.GET("/", handlers.GetAllDishHandler.Handler)
	dishesGroup.GET("/:dishID", handlers.GetDishHandler.Handler)

	ordersGroup := r.Group("/orders")
	ordersGroup.GET("/queued", handlers.GetQueuedOrdersQueueHandler.Handler)
	ordersGroup.GET("/in-progress", handlers.GetInProgressOrdersQueueHandler.Handler)
	ordersGroup.GET("/finished", handlers.GetFinishedOrdersQueueHandler.Handler)
	ordersGroup.POST("/new-random-order", handlers.CreateNewRandomOrderHandler.Handler)
	ordersGroup.PUT("/queues/:queue", handlers.UpdateOrdersQueuesHandler.Handler)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
