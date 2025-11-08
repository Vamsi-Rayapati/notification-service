package notification

import "github.com/gin-gonic/gin"

func RegisterRoutes(group *gin.RouterGroup) {
	notificationController := NotificationController{
		service: NotificationService{},
	}
	group.GET("/details", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "active",
		})
	})
	group.POST("/send", notificationController.SendNotification)
}
