package models

import (
	"booking-service/package/util"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Job struct {
	ID          primitive.ObjectID `bson:"_id"`
	Date        time.Time          `bson:"date"`
	Description string             `bson:"description"`
	Type        util.JobType       `bson:"type"`
	Status      util.JobStatus     `bson:"status"`
	Price       int                `bson:"price"`
	Address     string             `bson:"address"`
	Duration    int                `bson:"duration"`
	CreateAt    int                `bson:"create_at"`
	UpdateAt    int                `bson:"update_at"`
	DeleteAt    int                `bson:"delete_at"`
}

type JobRepository struct {
	Collection *mongo.Collection
}

func NewJobRepository(collection *mongo.Collection) *JobRepository {
	return &JobRepository{
		Collection: collection,
	}
}

func (r *JobRepository) CreateJob(job *Job) error {
	currentTime := int(time.Now().UnixMilli())
	job.CreateAt = currentTime
	job.UpdateAt = currentTime
	job.ID = primitive.NewObjectID()

	_, err := r.Collection.InsertOne(context.Background(), job)
	if err != nil {
		return err
	}

	return nil
}
