package services

import (
	"context"
	"errors"
	"job-portal-api/internal/auth"
	"job-portal-api/internal/database"
	"job-portal-api/internal/models"
)

//go:generate mockgen -source serviceInterface.go -destination mockmodels/serviceInterface_mock.go -package mockmodels
type Service struct {
	UserRepo database.UserRepo
	auth     auth.Authentication
}

type UserService interface {
	UserSignup(ctx context.Context, newUser models.NewUser) (models.User, error)
	UserLogin(ctx context.Context, userData models.NewUser) (string, error)

	StoreCompany(ctx context.Context, companyData models.NewCompany) (models.Company, error)
	GetCompanyData(ctx context.Context, companyId uint64) (models.Company, error)
	GetAllCompanyData(ctx context.Context) ([]models.Company, error)

	StoreJob(ctx context.Context, newJob models.NewJob, companyId uint64) (models.JobRespo, error)
	GetJobData(ctx context.Context, jobId uint64) (models.Job, error)
	GetAllJobData(ctx context.Context) ([]models.Job, error)
	GetJobByCompany(ctx context.Context, companyId uint64) ([]models.Job, error)
	ProcessJobApplication(ctx context.Context, applicationData []models.UserApplication) ([]models.ApplRespo, error)
}

func NewService(userRepo database.UserRepo, a auth.Authentication) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("interface cannot be null")
	}
	return &Service{
		UserRepo: userRepo,
		auth:     a,
	}, nil
}
