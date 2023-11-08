package models

type UserApplication struct {
	Name           string      `json:"name"`
	Age            uint        `json:"age"`
	JobId          uint        `json:"jobId"`
	JobApplication Application `json:"jobApplication"`
}

type Application struct {
	Jobname        string `json:"jobName" validate:"required"`
	NoticePeriod   string   `json:"noticePeriod" validate:"required"`
	JobLocations   []uint `json:"location" `
	Technologies   []uint `json:"technologyStack" `
	WorkModes      []uint `json:"workModes"`
	Experience     string   `json:"experience" validate:"required"`
	Qualifications []uint `json:"qualifications"`
	Shifts         []uint `json:"shifts"`
	Jobtypes       string `json:"jobType" validate:"required"`
}
