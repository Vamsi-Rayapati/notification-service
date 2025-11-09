package notification

import (
	"github.com/google/uuid"
	"github.com/smartbot/notification/database"
)

type PushMessage struct {
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
}

type EmailMessage struct {
	Subject string `json:"subject" validate:"required"`
	Body    string `json:"body" validate:"required"`
}

type Message struct {
	Push  PushMessage       `json:"push" validate:"omitempty"`
	Email EmailMessage      `json:"email" validate:"omitempty"`
	Data  map[string]string `json:"data" validate:"omitempty"`
}

type NotificationReceiver struct {
	UserID uuid.UUID `json:"user_id" validate:"required,uuid4"`
	Email  string    `json:"email" validate:"required,email"`
}

type SendNotificationRequest struct {
	Channels []database.NotificationChannel `json:"channels" validate:"required"`
	SenderID uuid.UUID                      `json:"sender_id" validate:"required,uuid4"`
	Receiver NotificationReceiver           `json:"receiver" validate:"required"`
	Message  Message                        `json:"message" validate:"required"`
}

type SendNotificationResponse struct {
	NotificationID uuid.UUID `json:"notification_id"`
}

type EmailKafkaPayload struct {
	Receiver NotificationReceiver `json:"receiver"`
	Subject  string               `json:"subject"`
	Body     string               `json:"body"`
}

type PushKafkaPayload struct {
	Receiver NotificationReceiver `json:"receiver"`
	Title    string               `json:"title"`
	Body     string               `json:"body"`
	Data     map[string]string    `json:"data"`
}
