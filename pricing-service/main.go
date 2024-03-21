package main

import (
	"log"
	"net/http"
	"os"
	"pricing-service/models"
	"pricing-service/package/util"
	"pricing-service/routers"
	"pricing-service/services"
	"time"

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

	// create data example
	dateTest := time.Unix(int64(1710981711000), 0)
	dateTest = util.RoundToDay(dateTest)

	// create data helper example
	priceExample := &models.Price{
		Date:    dateTest,
		Name:    "Don dep nha",
		Type:    util.NormalTime,
		JobType: util.Cleaning,
		Price:   100000,
	}

	// create data helper example
	priceExample2 := &models.Price{
		Date:    dateTest,
		Name:    "Cham soc tre vao gio binh thuong",
		Type:    util.NormalTime,
		JobType: util.Babysitting,
		Price:   100000,
	}

	err = models.Repository.Price.CreatePrice(priceExample)
	if err != nil {
		log.Fatalf("Create price failed: %v", err)
	}

	err = models.Repository.Price.CreatePrice(priceExample2)
	if err != nil {
		log.Fatalf("Create price failed: %v", err)
	}

	log.Printf("Server is running at %s\n", "http://localhost:8071")
	err = http.ListenAndServe(":8071", router)
	if err != nil {
		log.Fatal(err)
	}
}
