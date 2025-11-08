package client

import "github.com/segmentio/kafka-go"

type KafkaClient struct {
}

func (client *KafkaClient) GetWriter() *kafka.Writer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "test-topic",
	})

	return writer
}

var kakaClient *KafkaClient

func GetKafkaClient() *KafkaClient {
	if kakaClient == nil {
		kakaClient = &KafkaClient{}
	}
	return kakaClient
}
