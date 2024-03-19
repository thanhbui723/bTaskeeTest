package models

import "time"

type AppRepository struct {
	Price PriceRepositoryIface
}

type PriceRepositoryIface interface {
	GetPriceByDateAndType(date time.Time, jobType string) (Price, error)
	CreatePrice(price Price) error
}
