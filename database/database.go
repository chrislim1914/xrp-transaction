package database

import (
	"fmt"
	"log"

	"github.com/chrislim1914/xrp-transaction/bootstrap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *bootstrap.Config) (*gorm.DB, error) {
	var err error
	dsn := connectionStr(config)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
		return DB, err
	}
	return DB, nil
}

func connectionStr(config *bootstrap.Config) string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	return dsn
}
