package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/smartbot/notification/pkg/client"
	"github.com/smartbot/notification/pkg/config"
)

func main() {
	var err error
	config.LoadConfig()

	mysql := client.GetMySQLCient()
	db, err := mysql.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v, %v", err, db)
		return
	}

	// err = db.AutoMigrate(&database.Notification{})

	if err != nil {
		log.Fatalf("Migration failed: %v", err)

	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "email-notifications",
		GroupID: "email-group",
	})
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Received message: key=%s value=%s\n", string(msg.Key), string(msg.Value))
	}

}
