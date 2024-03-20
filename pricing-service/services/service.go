package services

import (
	"time"
)

type PriceService struct{}

type PriceServiceIface interface {
	GetPriceByDateAndType(date time.Time, jobType string, priceType string, duration int) (int, error)
}

var Price PriceServiceIface

func SetUp() {
	Price = &PriceService{}
}
