package database

import (
	"context"
	"errors"
	"job-portal-api/internal/models"

	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

//go:generate mockgen -source=db.go -destination=db_mock.go -package=database

type UserRepo interface {
	CreateUser(ctx context.Context, userData models.User) (models.User, error)
	CheckEmail(ctx context.Context, email string) (models.User, error)

	CreateCompany(ctx context.Context, companyData models.Company) (models.Company, error)
	ViewCompanies(ctx context.Context) ([]models.Company, error)
	ViewCompanyById(ctx context.Context, companyId uint64) (models.Company, error)

	CreateJob(ctx context.Context, jobData models.Job) (models.JobRespo, error)
	FindJob(ctx context.Context, companyId uint64) ([]models.Job, error)
	FindAllJobs(ctx context.Context) ([]models.Job, error)
	ViewJobDetailsBy(ctx context.Context, jobId uint64) (models.Job, error)
	GetTheJobData(jobid uint) (models.Job, error)
	AutoMigrate() error
}

func NewRepository(db *gorm.DB) (UserRepo, error) {
	if db == nil {
		return nil, errors.New("db cannot be null")
	}
	return &Repo{
		DB: db,
	}, nil
}
