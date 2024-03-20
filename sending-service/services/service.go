package services

import "sending-service/models"

type HelperService struct{}

type HelperServiceIface interface {
	CreateHelper(helper *models.Helper) error
}

var Helper HelperServiceIface

func SetUp() {
	Helper = &HelperService{}
}
