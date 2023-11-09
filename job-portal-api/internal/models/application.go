package models

type UserApplication struct {
	Name           string      `json:"name"`
	Age            uint        `json:"age"`
	JobId          uint        `json:"jobId"`
	JobApplication Application `json:"jobApplication"`
}

type Application struct {
	Title                string `json:"title" validate:"required"`
	MinNPInMonths  uint   `json:"minNPInMonths" validate:"required"`
	MaxNPInMonths  uint   `json:"maxNPInMonths" validate:"required"`
	JobLocations         []uint `json:"location" `
	Technologies         []uint `json:"technologyStacks" `
	WorkModes            []uint `json:"workModes"`
	Experience           uint   `json:"experience" validate:"required"`
	Qualifications       []uint `json:"qualifications"`
	Shifts               []uint `json:"shifts"`
	Jobtypes             string `json:"jobType" validate:"required"`
}

type ApplRespo struct {
	Name           string      `json:"name"`
	JobId          uint        `json:"jobId"`
}
