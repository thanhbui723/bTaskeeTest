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
