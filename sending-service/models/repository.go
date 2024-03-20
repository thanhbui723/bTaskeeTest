package models

type AppRepository struct {
	Helper HelperRepositoryIface
}

type HelperRepositoryIface interface {
	CreateHelper(helper *Helper) error
}
