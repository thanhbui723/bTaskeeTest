package services

import (
	"log"
	"pricing-service/models"
	"pricing-service/package/util"

	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetPriceByDateAndType(t *testing.T) {
	// Initialize database connection
	err := models.SetUp("mongodb://localhost:27019", "pricing")
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	SetUp()

	var (
		expectedPrice = 300000
	)

	Convey("Given a PriceService", t, func() {
		Convey("When a price is found for the given date, job type, and price type", func() {
			price := expectedPrice // Replace with logic to fetch price from repository

			date := time.Date(2024, 3, 21, 0, 0, 0, 0, time.UTC)
			jobType := string(util.Cleaning)
			priceType := string(util.NormalTime)
			duration := 3

			totalCost, err := Price.GetPriceByDateAndType(date, jobType, priceType, duration)

			Convey("Then the total cost is calculated correctly", func() {
				So(totalCost, ShouldEqual, price)
				So(err, ShouldBeNil)
			})
		})

	})
}
