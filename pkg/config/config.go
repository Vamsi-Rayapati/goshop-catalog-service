package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type ApplicationConfig struct {
	Port       int
	DbHost     string
	DbPort     int
	DbUserName string
	DbPassword string
	FaSecret   string
}

var Config *ApplicationConfig

func LoadConfig() *ApplicationConfig {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	Config = &ApplicationConfig{
		Port:       viper.GetInt("PORT"),
		DbHost:     viper.GetString("DB_HOST"),
		DbPort:     viper.GetInt("DB_PORT"),
		DbUserName: viper.GetString("DB_USER_NAME"),
		DbPassword: viper.GetString("DB_PASSWORD"),
		FaSecret:   viper.GetString("FA_SECRET"),
	}

	return Config

}

// ConnectDB initializes the database connection
func ConnectDB() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.DbUserName,
		Config.DbPassword,
		Config.DbHost,
		Config.DbPort,
		"catalog", // DB name
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return err
	}
	log.Println("Database connection successful!")
	return nil
}
