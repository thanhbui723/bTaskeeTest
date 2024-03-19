package main

import (
	"log"
	"pricing-service/models"
	"pricing-service/routers"
	"pricing-service/services"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// Initialize database connection
	err := models.InitDB("mongodb://localhost:27017", "pricing")
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	services.SetUp()
	routers.InitRouter()

	// Create a price repository
	priceRepo := models.NewPriceRepository("prices")

	// Create a sample price
	price := models.Price{
		ID:    primitive.NewObjectID(),
		Date:  time.Now(),
		Name:  "Sample Price",
		Type:  "Sample Type",
		Price: 100,
	}

	// Create a sample price
	price2 := models.Price{
		ID:    primitive.NewObjectID(),
		Date:  time.Now(),
		Name:  "Sample Price",
		Type:  "Sample Type",
		Price: 100,
	}

	// Insert the sample price into the database
	err = priceRepo.CreatePrice(price)
	if err != nil {
		log.Fatalf("Error creating price: %v", err)
	}

	err = priceRepo.CreatePrice(price2)
	if err != nil {
		log.Fatalf("Error creating price: %v", err)
	}

	log.Println("Price created successfully")
}
