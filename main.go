package main

import (
	"fmt"

	"log"

	"github.com/smartbot/catalog/api"
	"github.com/smartbot/catalog/database"
	"github.com/smartbot/catalog/pkg/config"
	"github.com/smartbot/catalog/pkg/dbclient"
)

func main() {
	var err error
	config.LoadConfig()

	db, err := dbclient.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v, %v", err, db)
		return
	}

	err = db.AutoMigrate(&database.Product{}, &database.Category{}, &database.ProductImages{})

	if err != nil {
		log.Fatalf("Migration failed: %v", err)

	}
	r := api.RegisterRoutes()
	r.Run(fmt.Sprintf("%s%d", ":", config.Config.Port))

}
