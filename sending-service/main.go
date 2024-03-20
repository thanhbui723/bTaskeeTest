package main

import (
	"log"
	"net/http"
	"sending-service/models"
	"sending-service/package/util"
	"sending-service/routers"
	"sending-service/services"

	"github.com/go-playground/validator/v10"
)

func main() {
	// Initialize database connection
	err := models.SetUp("mongodb://localhost:27017", "sending")
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	services.SetUp()
	router := routers.InitRouter()
	util.Validator = validator.New()

	helper := &models.Helper{
		Name:    "thanh",
		Skill:   "cleaning",
		Address: "phu nhuan",
		Phone:   2234345,
		Rating:  5,
	}

	err = models.Repository.Helper.CreateHelper(helper)
	if err != nil {
		log.Fatalf("Create helper failed: %v", err)
	}

	err = http.ListenAndServe(":3002", router)
	if err != nil {
		log.Fatal(err)
	}
}
