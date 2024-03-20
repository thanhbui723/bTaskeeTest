package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Assignment struct {
	ID       primitive.ObjectID `bson:"_id"`
	JobID    primitive.ObjectID `bson:"job_id"`
	HelperID primitive.ObjectID `bson:"helper_id"`
}

type AssignmentRepository struct {
	Collection *mongo.Collection
}

func NewAssignmentRepository(collection *mongo.Collection) *AssignmentRepository {
	return &AssignmentRepository{
		Collection: collection,
	}
}

func (r *AssignmentRepository) CreateAssignment(assignment *Assignment) error {
	assignment.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), assignment)
	if err != nil {
		return err
	}

	return nil
}
