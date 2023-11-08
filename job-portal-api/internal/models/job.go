package models

import "gorm.io/gorm"

///////////////////////////////////////////////////////////////////////////////////////////////////
//for validation purpose
type NewCompany struct {
	CompanyName string `json:"Companyname" validate:"required"`
	Location    string `json:"Location" validate:"required"`
}

type Company struct {
	gorm.Model
	CompanyName string `json:"Companyname" gorm:"unique;not null" validate:"required,unique"`
	Location    string `json:"Location"`
}

type NewJob struct {
	CompanyId      uint   `json:"CompanyId"`
	Title          string `json:"Title" validate:"required"`
	MinNP          string `json:"MinNP" validate:"required"`
	MaxNP          string `json:"MaxNP" validate:"required"`
	Budget         string `json:"Budget" validate:"required"`
	JobLocations   []uint `json:"JobLocations" validate:"required"`
	Technologies   []uint `json:"TechnologyStacks" validate:"required"`
	WorkModes      []uint `json:"WorkModes" validate:"required"`
	MinExp         string `json:"MinExp" validate:"required"`
	MaxExp         string `json:"MaxExp" validate:"required"`
	Qualifications []uint `json:"Qualifications" validate:"required"`
	Shifts         []uint `json:"Shifts" validate:"required"`
	JobTypes       string `json:"JobType" validate:"required"`
	Desc           string `json:"Desc" validate:"required"`
}

type JobRespo struct {
	JobId uint `json:"JobId"`
}

type Job struct {
	gorm.Model
	Company        Company         `json:"-" gorm:"ForeignKey:CompanyId"`
	CompanyId      uint            `json:"CompanyId"`
	Title          string          `json:"Title"`
	MinNP          string          `json:"MinNP"`
	MaxNP          string          `json:"MaxNP"`
	Budget         string          `json:"Budget"`
	JobLocations   []Location      `gorm:"many2many:Job_locations;"`
	Technologies   []Technology    `gorm:"many2many:Job_technologies;"`
	WorkModes      []WorkMode      `gorm:"many2many:Job_workmode;"`
	MinExp         string          `json:"MinExp"`
	MaxExp         string          `json:"MaxExp"`
	Qualifications []Qualification `gorm:"many2many:Job_qualifications;"`
	Shifts         []Shift         `gorm:"many2many:Job_shifts;"`
	JobTypes       string          `json:"JobType"`
	Desc           string          `json:"Desc"`
}

type Location struct {
	gorm.Model
	City string `json:"City" gorm:"unique;not null"`
}

type Technology struct {
	gorm.Model
	TechName string `json:"Techname" gorm:"unique;not null"`
}

type WorkMode struct {
	gorm.Model
	ModeType string `json:"Modetype" gorm:"unique;not null"`
}

type Qualification struct {
	gorm.Model
	QualificationReq string `json:"Qualreq" gorm:"unique;not null"`
}

type Shift struct {
	gorm.Model
	ShiftType string `json:"ShiftType" gorm:"unique;not null"`
}
