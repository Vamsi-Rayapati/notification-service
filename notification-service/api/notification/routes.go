package notification

import "github.com/gin-gonic/gin"

func RegisterRoutes(group *gin.RouterGroup) {
	notificationController := NotificationController{
		service: NotificationService{},
	}
	group.POST("/send", notificationController.SendNotification)
}
