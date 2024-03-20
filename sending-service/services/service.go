package services

import "sending-service/models"

type HelperService struct{}

type HelperServiceIface interface {
	CreateHelper(helper *models.Helper) error
}

type AssignmentService struct{}

type AssignmentServiceIface interface {
	CreateAssignment(jobID, helperID string) error
}

var Assignment AssignmentServiceIface
var Helper HelperServiceIface

func SetUp() {
	Helper = &HelperService{}
	Assignment = &AssignmentService{}
}
