package util

type PriceType string

const (
	PeakTime    PriceType = "peak"
	NormalTime  PriceType = "normal"
	HolidayTime PriceType = "holiday"
)

type JobType string

const (
	Cleaning    JobType = "cleaning"
	Babysitting JobType = "babysitting"
)

type JobStatus string

const (
	Pending   JobStatus = "pending"
	Assigned  JobStatus = "assigned"
	Completed JobStatus = "completed"
	Cancelled JobStatus = "cancelled"
)
