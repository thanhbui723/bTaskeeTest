package main

import (
	"log"
	"net/http"
	"pricing-service/models"
	"pricing-service/package/util"
	"pricing-service/routers"
	"pricing-service/services"

	"github.com/go-playground/validator/v10"
)

func main() {
	// Initialize database connection
	err := models.SetUp("mongodb://localhost:27017", "pricing")
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	services.SetUp()
	router := routers.InitRouter()
	util.Validator = validator.New()

	err = http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}
