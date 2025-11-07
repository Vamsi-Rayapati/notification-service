package notification

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smartbot/notification/pkg/validator"
)

type NotificationController struct {
	service NotificationService
}

func (nc *NotificationController) SendNotification(c *gin.Context) {
	var req SendNotificationRequest
	err := validator.ValidateBody(c, &req)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	res, err := nc.service.SendNotification(req)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, res)

}
