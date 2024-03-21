package models

type AppRepository struct {
	Helper     HelperRepositoryIface
	Assignment AssignmentRepositoryIface
}

type HelperRepositoryIface interface {
	CreateHelper(helper *Helper) error
	GetHelpers() ([]*Helper, error)
}

type AssignmentRepositoryIface interface {
	CreateAssignment(assignment *Assignment) error
}
