package notification

import (
	"github.com/google/uuid"
	"github.com/smartbot/notification/database"
)

type SendNotificationRequest struct {
	Channel  database.NotificationChannel `json:"channel" validate:"required,oneof=email push"`
	SenderID uuid.UUID                    `json:"sender_id" validate:"required,uuid4"`
	Payload  map[string]interface{}       `json:"payload"`
}

type SendNotificationResponse struct {
	NotificationID uuid.UUID `json:"notification_id"`
}
