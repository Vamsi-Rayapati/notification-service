package client

import (
	"fmt"
	"log"

	"github.com/smartbot/notification/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLClient struct {
	DB *gorm.DB
}

// var DB *gorm.DB

// ConnectDB initializes the database connection
func (client *MySQLClient) Connect() (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.DbUserName,
		config.Config.DbPassword,
		config.Config.DbHost,
		config.Config.DbPort,
		"notification",
	)

	client.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}
	log.Println("Database connection successful!")
	return client.DB, nil
}

func (client *MySQLClient) GetDatabase() *gorm.DB {
	return client.DB

}

var mysqlClient *MySQLClient

func GetMySQLCient() *MySQLClient {
	if mysqlClient == nil {
		mysqlClient = &MySQLClient{
			DB: nil,
		}
	}
	return mysqlClient
}
