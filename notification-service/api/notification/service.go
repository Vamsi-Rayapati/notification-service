package notification

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/smartbot/notification/database"
	"github.com/smartbot/notification/pkg/client"
	"github.com/smartbot/notification/pkg/errors"
)

func UUIDToBytes(u uuid.UUID) []byte {
	return u[:16]
}

type NotificationService struct {
}

func (ns *NotificationService) SendNotification(req SendNotificationRequest) (*SendNotificationResponse, *errors.ApiError) {
	notificationID := uuid.New()
	kafkaClient := client.GetKafkaClient()
	writer := kafkaClient.GetWriter("email-notifications")
	defer writer.Close()

	bytePayload, err := json.Marshal(req.Payload)
	if err != nil {
		return nil, errors.InternalServerError("Failed to parse payload")
	}

	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(notificationID.String()),
			Value: bytePayload,
		},
	)

	if err != nil {
		log.Println("failed to write messages:", err)
		return nil, errors.InternalServerError("Failed to send notification")
	}

	db := client.GetMySQLCient().GetDatabase()

	newNotif := database.Notification{
		ID:      notificationID,
		UserID:  req.SenderID, // req.SenderID,
		Channel: req.Channel,
		Status:  database.NotificationSending,
	}

	result := db.Create(&newNotif)
	if result.Error != nil {
		return nil, errors.InternalServerError("Failed to create product")
	}

	return &SendNotificationResponse{NotificationID: newNotif.ID}, nil

}
