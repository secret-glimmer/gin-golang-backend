package main

import (
	"app/config"
	"app/db"
	"app/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Failed to load environment variables from .env file: %v", err)
		return
	}

	cfg := config.GetConfig()

	db, err := db.NewDBConnection(cfg.DatabaseConfig)
	if err != nil {
		log.Printf("Failed to connect database: %v", err)
		return
	}

	router := routes.SetupRouter(db)

	router.Run(":" + cfg.ServerAddress)
}
