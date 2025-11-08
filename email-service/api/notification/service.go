package notification

import (
	"context"
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

	kafkaClient := client.GetKafkaClient()
	writer := kafkaClient.GetWriter()
	defer writer.Close()

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("key-A"),
			Value: []byte("Hello from Go!"),
		},
	)

	if err != nil {
		log.Println("failed to write messages:", err)
		return nil, errors.InternalServerError("Failed to send notification")
	}

	db := client.GetMySQLCient().GetDatabase()

	newNotif := database.Notification{
		ID:      uuid.New(),
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
