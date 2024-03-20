package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db         *mongo.Database
	Repository *AppRepository
)

func SetUp(connectionString, dbName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	db = client.Database(dbName)
	Repository = &AppRepository{
		Helper:     NewHelperRepository(db.Collection("helpers")),
		Assignment: NewAssignmentRepository(db.Collection("assignments")),
	}

	return nil
}
