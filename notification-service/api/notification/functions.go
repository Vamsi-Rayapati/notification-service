package notification

import "encoding/json"

func getEmailKafkaPayload(req SendNotificationRequest) ([]byte, error) {
	return json.Marshal(EmailKafkaPayload{
		Receiver: req.Receiver,
		Subject:  req.Message.Email.Subject,
		Body:     req.Message.Email.Body,
	})
}

func getPushKafkaPayload(req SendNotificationRequest) ([]byte, error) {
	return json.Marshal(EmailKafkaPayload{
		Receiver: req.Receiver,
		Subject:  req.Message.Email.Subject,
		Body:     req.Message.Email.Body,
	})
}
