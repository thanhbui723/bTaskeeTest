package services

import (
	"booking-service/models"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrJobNotFound = errors.New("job not found")
)

func (s *JobService) CreateJob(job *models.Job) (*models.Job, error) {
	return &models.Job{}, nil
}

func (s *JobService) GetJobByID(jobID string) (*models.Job, error) {
	// convert job_id string to objectID
	objID, err := primitive.ObjectIDFromHex(jobID)
	if err != nil {
		return nil, err
	}
	fmt.Println("objID: ", objID)

	job, err := models.Repository.Job.GetJobByID(objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			fmt.Println("error: ", err)
			return nil, ErrJobNotFound
		}

		return nil, err
	}

	return job, nil
}
