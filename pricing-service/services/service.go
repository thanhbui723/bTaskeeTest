package services

import (
	"pricing-service/models"
	"time"
)

type PriceService struct{}

type PriceServiceIface interface {
	GetPriceByDateAndType(date time.Time, jobType string) (models.Price, error)
}

var Price PriceServiceIface

func SetUp() {
	Price = &PriceService{}
}
