package dbclient

import (
	"fmt"
	"log"

	"github.com/smartbot/catalog/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB initializes the database connection
func Connect() (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.DbUserName,
		config.Config.DbPassword,
		config.Config.DbHost,
		config.Config.DbPort,
		"catalog",
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}
	log.Println("Database connection successful!")
	return DB, nil
}

func GetCient() *gorm.DB {
	return DB
}
