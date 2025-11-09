package notification

import (
	"context"
	"log"
	"slices"

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
	writer := kafkaClient.GetWriter()
	defer writer.Close()

	var msgs []kafka.Message = []kafka.Message{}
	var notifications []database.Notification = []database.Notification{}

	if slices.Contains(req.Channels, "email") {
		emailPayload, err := getEmailKafkaPayload(req)

		if err != nil {
			return nil, errors.InternalServerError("Failed to parse payload")
		}

		msgs = append(msgs, kafka.Message{
			Topic: "email-notifications",
			Key:   []byte(notificationID.String()),
			Value: emailPayload,
		})

		notifications = append(notifications, database.Notification{
			ID:      notificationID,
			UserID:  req.SenderID, // req.SenderID,
			Channel: "email",
			Status:  database.NotificationSending,
		})
	}

	if slices.Contains(req.Channels, "push") {
		emailPayload, err := getEmailKafkaPayload(req)

		if err != nil {
			return nil, errors.InternalServerError("Failed to parse payload")
		}

		msgs = append(msgs, kafka.Message{
			Topic: "push-notifications",
			Key:   []byte(notificationID.String()),
			Value: emailPayload,
		})

		notifications = append(notifications, database.Notification{
			ID:      notificationID,
			UserID:  req.SenderID, // req.SenderID,
			Channel: "push",
			Status:  database.NotificationSending,
		})

	}

	err := writer.WriteMessages(context.Background(), msgs...)

	if err != nil {
		log.Println("failed to write messages:", err)
		return nil, errors.InternalServerError("Failed to send notification")
	}

	db := client.GetMySQLCient().GetDatabase()

	result := db.Create(&notifications)
	if result.Error != nil {
		return nil, errors.InternalServerError("Failed to create product")
	}

	return &SendNotificationResponse{NotificationID: notificationID}, nil

}
