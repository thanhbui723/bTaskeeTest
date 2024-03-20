package services

import (
	"booking-service/models"
)

type JobService struct{}

type JobServiceIface interface {
	CreateJob(job *models.Job) (*models.Job, error)
	GetJobByID(jobID string) (*models.Job, error)
}

var Job JobServiceIface

func SetUp() {
	Job = &JobService{}
}
