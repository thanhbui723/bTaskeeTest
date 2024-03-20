package services

import (
	"sending-service/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *AssignmentService) CreateAssignment(jobID, helperID string) error {
	jobObjectID, err := primitive.ObjectIDFromHex(jobID)
	if err != nil {
		return err
	}

	helperObjectID, err := primitive.ObjectIDFromHex(helperID)
	if err != nil {
		return err
	}

	assignment := &models.Assignment{
		JobID:    jobObjectID,
		HelperID: helperObjectID,
	}

	err = models.Repository.Assignment.CreateAssignment(assignment)
	if err != nil {
		return err
	}

	return nil
}
