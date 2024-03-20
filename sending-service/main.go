package main

import (
	"log"
	"net/http"
	"os"
	"sending-service/models"
	"sending-service/package/util"
	"sending-service/routers"
	"sending-service/services"

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

	// create data helper example
	helperExample := &models.Helper{
		Name:    "alex",
		Skill:   "cleaning",
		Address: "district 1 TP HCM",
		Phone:   84142234345,
		Rating:  5,
	}

	helperExample2 := &models.Helper{
		Name:    "john",
		Skill:   "babysitting",
		Address: "district 1 TP HCM",
		Phone:   8442234345,
		Rating:  5,
	}

	err = models.Repository.Helper.CreateHelper(helperExample)
	if err != nil {
		log.Fatalf("Create helper failed: %v", err)
	}

	err = models.Repository.Helper.CreateHelper(helperExample2)
	if err != nil {
		log.Fatalf("Create helper failed: %v", err)
	}

	log.Printf("Server is running at %s\n", "http://localhost:8073")
	err = http.ListenAndServe(":8073", router)
	if err != nil {
		log.Fatal(err)
	}
}
