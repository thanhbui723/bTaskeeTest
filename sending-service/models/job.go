package models

import (
	"time"

	"sending-service/package/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
