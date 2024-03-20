package main

import (
	"booking-service/models"
	"booking-service/package/util"
	"booking-service/routers"
	"booking-service/services"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get MongoDB connection string from environment variable
	connectionString := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("DB_NAME")

	// Initialize database connection
	err = models.SetUp(connectionString, dbName)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	services.SetUp()
	router := routers.InitRouter()
	util.Validator = validator.New()

	log.Printf("Server is running at %s\n", "http://localhost:8072")
	err = http.ListenAndServe(":8072", router)
	if err != nil {
		log.Fatal(err)
	}
}
