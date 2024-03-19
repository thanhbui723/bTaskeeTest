package services

import (
	"pricing-service/models"
	"time"
)

func (s *PriceService) GetPriceByDateAndType(date time.Time, jobType string) (models.Price, error) {
	price, err := models.Repository.Price.GetPriceByDateAndType(date, jobType)
	if err != nil {
		return models.Price{}, err
	}
	return price, nil

}
