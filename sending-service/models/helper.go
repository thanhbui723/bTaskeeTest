package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Helper struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Skill    string             `bson:"skill"`
	Address  string             `bson:"address"`
	Phone    int                `bson:"phone"`
	Rating   int                `bson:"rating"`
	CreateAt int                `bson:"create_at"`
	UpdateAt int                `bson:"update_at"`
	DeleteAt int                `bson:"delete_at"`
}

type HelperRepository struct {
	Collection *mongo.Collection
}

func NewHelperRepository(collection *mongo.Collection) *HelperRepository {
	return &HelperRepository{
		Collection: collection,
	}
}

func (r *HelperRepository) CreateHelper(helper *Helper) error {
	currentTime := int(time.Now().UnixMilli())
	helper.CreateAt = currentTime
	helper.UpdateAt = currentTime
	helper.ID = primitive.NewObjectID()

	_, err := r.Collection.InsertOne(context.Background(), helper)
	if err != nil {
		return err
	}

	return nil
}

func (r *HelperRepository) GetHelpers() ([]*Helper, error) {
    var helpers []*Helper

    cursor, err := r.Collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var helper *Helper
        if err := cursor.Decode(&helper); err != nil {
            return nil, err
        }
        helpers = append(helpers, helper)
    }
    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return helpers, nil
}