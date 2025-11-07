package main

import (
	"fmt"
	"log"

	"github.com/smartbot/notification/api"
	"github.com/smartbot/notification/database"
	"github.com/smartbot/notification/pkg/config"
	"github.com/smartbot/notification/pkg/dbclient"
)

func main() {
	var err error
	config.LoadConfig()

	db, err := dbclient.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v, %v", err, db)
		return
	}

	err = db.AutoMigrate(&database.Notification{})

	if err != nil {
		log.Fatalf("Migration failed: %v", err)

	}
	r := api.RegisterRoutes()
	r.Run(fmt.Sprintf("%s%d", ":", config.Config.Port))
}
