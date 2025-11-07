package notification

import (
	"log"

	"github.com/google/uuid"
	"github.com/smartbot/notification/database"
	"github.com/smartbot/notification/pkg/dbclient"
	"github.com/smartbot/notification/pkg/errors"
)

func UUIDToBytes(u uuid.UUID) []byte {
	return u[:16]
}

type NotificationService struct {
}

func (ns *NotificationService) SendNotification(req SendNotificationRequest) (*SendNotificationResponse, *errors.ApiError) {
	db := dbclient.GetCient()

	newNotif := database.Notification{
		ID:      uuid.New(),
		UserID:  req.SenderID, // req.SenderID,
		Channel: req.Channel,
		Status:  database.NotificationSending,
	}

	result := db.Create(&newNotif)
	if result.Error != nil {
		log.Fatal("Failed to create notif", result.Error.Error())
		return nil, errors.InternalServerError("Failed to create product")
	}

	return &SendNotificationResponse{NotificationID: newNotif.ID}, nil

}
