package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AppRepository struct {
	Job JobRepositoryIface
}

type JobRepositoryIface interface {
	CreateJob(job *Job) (*Job, error)
	GetJobByID(jobID primitive.ObjectID) (*Job, error)
}
