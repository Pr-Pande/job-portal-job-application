package database

import (
	"context"
	"errors"
	"job-portal-api/internal/models"
	"github.com/rs/zerolog/log"
)

func (r *Repo) ViewJobDetailsBy(ctx context.Context, jobId uint64) (models.Job, error) {
	var jobData models.Job
	result := r.DB.Where("id = ?", jobId).Find(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Job{}, errors.New("could not create the jobs")
	}
	return jobData, nil
}

func (r *Repo) CreateJob(ctx context.Context, jobData models.Job) (models.JobRespo, error) {
	result := r.DB.Create(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.JobRespo{}, errors.New("could not create the jobs")
	}
	return models.JobRespo{
		JobId: jobData.ID,
	}, nil
}

func (r *Repo) FindAllJobs(ctx context.Context) ([]models.Job, error) {
	var jobDatas []models.Job
	result := r.DB.Find(&jobDatas)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find the jobs")
	}
	return jobDatas, nil

}

func (r *Repo) FindJob(ctx context.Context, companyId uint64) ([]models.Job, error) {
	var jobData []models.Job
	result := r.DB.Where("company_id = ?", companyId).Find(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find the company")
	}
	return jobData, nil
}

func (r *Repo) GetTheJobData(jobid uint) (models.Job, error) {
	var jobData models.Job

	// Preload related data using GORM's Preload method
	result := r.DB.Preload("Company").
		Preload("JobLocations").
		Preload("Technologies").
		Preload("WorkModes").
		Preload("Qualifications").
		Preload("Shifts").
		Where("id = ?", jobid).
		Find(&jobData)

	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Job{}, result.Error
	}

	return jobData, nil
}
