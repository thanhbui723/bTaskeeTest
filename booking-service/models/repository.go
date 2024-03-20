package models

type AppRepository struct {
	Job JobRepositoryIface
}

type JobRepositoryIface interface {
	CreateJob(job *Job) error
}
