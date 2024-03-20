package models

import (
	"context"
	"pricing-service/package/util"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Price struct {
	ID      primitive.ObjectID `bson:"_id"`
	Date    time.Time          `bson:"date"`
	Name    string             `bson:"name"`
	Type    util.PriceType     `bson:"type"`
	JobType util.JobType       `bson:"job_type"`
	Price   int                `bson:"price"` // price per hour
}

type PriceRepository struct {
	Collection *mongo.Collection
}

func NewPriceRepository(collection *mongo.Collection) *PriceRepository {
	return &PriceRepository{
		Collection: collection,
	}
}

func (r *PriceRepository) GetPriceByDateAndType(date time.Time, jobType string, priceType string) (*Price, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var price *Price
	filter := bson.M{
		"date":     date,
		"type":     priceType,
		"job_type": jobType,
	}

	err := r.Collection.FindOne(ctx, filter).Decode(&price)
	if err != nil {
		return nil, err
	}

	return price, nil
}

func (r *PriceRepository) CreatePrice(price *Price) error {
	_, err := r.Collection.InsertOne(context.Background(), price)
	if err != nil {
		return err
	}

	return nil
}
