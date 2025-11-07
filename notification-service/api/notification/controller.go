package notification

import "github.com/gin-gonic/gin"

type NotificationController struct {
	service NotificationService
}

func (nc *NotificationController) SendNotification(c *gin.Context) {
	nc.service.SendNotification()
}
