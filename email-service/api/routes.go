package api

import (
	"github.com/gin-gonic/gin"
	"github.com/smartbot/notification/api/notification"
)

func RegisterRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	routeGroup := router.Group("/api/v1/notification")
	notification.RegisterRoutes(routeGroup)
	return router
}
