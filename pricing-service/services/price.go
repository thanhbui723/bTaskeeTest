package services

import (
	"errors"
	"pricing-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrPriceNotFound = errors.New("price not found")
)

func (s *PriceService) GetPriceByDateAndType(date time.Time, jobType string, priceType string, duration int) (int, error) {
	price, err := models.Repository.Price.GetPriceByDateAndType(date, jobType, priceType)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return -1, ErrPriceNotFound
		}

		return -1, err
	}

	totalCost := price.Price * duration
	return totalCost, nil

}
