package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	routeGroup := router.Group("/api/v1/notification")

	routeGroup.GET("/details2", func(ctx *gin.Context) {
		ctx.Status(200)
	})
	// notification.RegisterRoutes(routeGroup)
	return router
}
