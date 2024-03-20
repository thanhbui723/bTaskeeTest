package main

import (
	"booking-service/models"
	"booking-service/package/util"
	"booking-service/routers"
	"booking-service/services"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	// Initialize database connection
	err := models.SetUp("mongodb://localhost:27017", "booking")
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	services.SetUp()
	router := routers.InitRouter()
	util.Validator = validator.New()

	err = http.ListenAndServe(":3001", router)
	if err != nil {
		log.Fatal(err)
	}
}
