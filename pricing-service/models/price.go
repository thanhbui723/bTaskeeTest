package models

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Price struct {
	ID    primitive.ObjectID `bson:"_id"`
	Date  time.Time          `bson:"date"`
	Name  string             `bson:"name"`
	Type  string             `bson:"type"`
	Price int                `bson:"price"`
}

type PriceRepository struct {
	Collection *mongo.Collection
}

func NewPriceRepository(collectionName string) *PriceRepository {
	return &PriceRepository{
		Collection: db.Collection(collectionName),
	}
}

func (r *PriceRepository) GetPriceByDateAndType(date time.Time, jobType string) (Price, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var price Price
	filter := bson.M{
		"date": date,
		"type": jobType,
	}

	err := r.Collection.FindOne(ctx, filter).Decode(&price)
	if err != nil {
		return Price{}, err
	}

	return price, nil
}

func (r *PriceRepository) CreatePrice(price Price) error {
	_, err := r.Collection.InsertOne(context.Background(), price)
	if err != nil {
		return errors.New("Error creating price: " + err.Error())
	}

	return nil
}
