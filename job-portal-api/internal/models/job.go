package models

import "gorm.io/gorm"

///////////////////////////////////////////////////////////////////////////////////////////////////
//for validation purpose
type NewCompany struct {
	CompanyName string `json:"companyName" validate:"required"`
	Location    string `json:"location" validate:"required"`
}

type Company struct {
	gorm.Model
	CompanyName string `json:"companyName" gorm:"unique;not null" validate:"required,unique"`
	Location    string `json:"location"`
}

type NewJob struct {
	Title          string `json:"title" validate:"required"`
	MinNPInMonths  uint   `json:"minNPInMonths" validate:"required"`
	MaxNPInMonths  uint   `json:"maxNPInMonths" validate:"required"`
	Budget         string `json:"budget" validate:"required"`
	JobLocations   []uint `json:"jobLocations" validate:"required"`
	Technologies   []uint `json:"technologyStacks" validate:"required"`
	WorkModes      []uint `json:"workModes" validate:"required"`
	MinExp         uint   `json:"minExp" validate:"required"`
	MaxExp         uint   `json:"maxExp" validate:"required"`
	Qualifications []uint `json:"qualifications" validate:"required"`
	Shifts         []uint `json:"shifts" validate:"required"`
	JobTypes       string `json:"jobType" validate:"required"`
	Desc           string `json:"desc" validate:"required"`
}

type JobRespo struct {
	JobId uint `json:"jobId"`
}

type Job struct {
	gorm.Model
	Company        Company         `json:"-" gorm:"ForeignKey:CompanyId"`
	CompanyId      uint            `json:"companyId"`
	Title          string          `json:"title"`
	MinNPInMonths  uint            `json:"minNPInMonths"`
	MaxNPInMonths  uint            `json:"maxNPInMonths"`
	Budget         string          `json:"budget"`
	JobLocations   []Location      `gorm:"many2many:Job_locations;"`
	Technologies   []Technology    `gorm:"many2many:Job_technologies;"`
	WorkModes      []WorkMode      `gorm:"many2many:Job_workmode;"`
	MinExp         uint            `json:"minExp"`
	MaxExp         uint            `json:"maxExp"`
	Qualifications []Qualification `gorm:"many2many:Job_qualifications;"`
	Shifts         []Shift         `gorm:"many2many:Job_shifts;"`
	JobTypes       string          `json:"jobType"`
	Desc           string          `json:"desc"`
}

type Location struct {
	gorm.Model
	City string `json:"city" gorm:"unique;not null"`
}

type Technology struct {
	gorm.Model
	TechName string `json:"techName" gorm:"unique;not null"`
}

type WorkMode struct {
	gorm.Model
	ModeType string `json:"modeType" gorm:"unique;not null"`
}

type Qualification struct {
	gorm.Model
	QualificationReq string `json:"qualReq" gorm:"unique;not null"`
}

type Shift struct {
	gorm.Model
	ShiftType string `json:"shiftType" gorm:"unique;not null"`
}
