package client

import (
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/smartbot/notification/pkg/utils"
)

type KafkaClient struct {
}

func (client *KafkaClient) GetWriter() *kafka.Writer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"kafka:9092"},
	})

	return writer
}

func (client *KafkaClient) CreateTopics(topics []string) {
	conn, err := kafka.Dial("tcp", "kafka:9092")
	if err != nil {
		log.Fatal("failed to connect to kafka:", err)
	}
	defer conn.Close()
	config := utils.Map(topics, func(topic string) kafka.TopicConfig {
		return kafka.TopicConfig{
			Topic:             topic,
			ReplicationFactor: 2,
			NumPartitions:     3,
		}
	})
	conn.CreateTopics(config...)
	if err != nil {
		log.Fatal("failed to create topic:", err)
	}
	log.Println("Topic created successfully!")

}

var kakaClient *KafkaClient

func GetKafkaClient() *KafkaClient {
	if kakaClient == nil {
		kakaClient = &KafkaClient{}
	}
	return kakaClient
}
