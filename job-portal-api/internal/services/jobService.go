package services

import (
	"context"
	"job-portal-api/internal/models"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (s *Service) StoreJob(ctx context.Context, jobData models.NewJob, companyId uint64) (models.JobRespo, error) {

	jDetails := models.Job{
		CompanyId: uint(companyId),
		Title:     jobData.Title,
		MinNP:     jobData.MinNP,
		MaxNP:     jobData.MaxNP,
		Budget:    jobData.Budget,
		MinExp:    jobData.MinExp,
		MaxExp:    jobData.MaxExp,
		JobTypes:  jobData.JobTypes,
		Desc:      jobData.Desc,
	}

	for _, v := range jobData.JobLocations {
		tempjDetails := models.Location{
			Model: gorm.Model{
				ID: v,
			},
		}
		jDetails.JobLocations = append(jDetails.JobLocations, tempjDetails)
	}

	for _, v := range jobData.Technologies {
		tempjDetails := models.Technology{
			Model: gorm.Model{
				ID: v,
			},
		}
		jDetails.Technologies = append(jDetails.Technologies, tempjDetails)
	}

	for _, v := range jobData.WorkModes {
		tempjDetails := models.WorkMode{
			Model: gorm.Model{
				ID: v,
			},
		}
		jDetails.WorkModes = append(jDetails.WorkModes, tempjDetails)
	}

	for _, v := range jobData.Qualifications {
		tempjDetails := models.Qualification{
			Model: gorm.Model{
				ID: v,
			},
		}
		jDetails.Qualifications = append(jDetails.Qualifications, tempjDetails)
	}

	for _, v := range jobData.Shifts {
		tempjDetails := models.Shift{
			Model: gorm.Model{
				ID: v,
			},
		}
		jDetails.Shifts = append(jDetails.Shifts, tempjDetails)
	}

	job, err := s.UserRepo.CreateJob(ctx, jDetails)
	if err != nil {
		return models.JobRespo{}, err
	}
	return job, nil

}

// ///////////////////////////////////////////////////////////////////////////////
func (s *Service) GetJobData(ctx context.Context, jobId uint64) (models.Job, error) {
	/* if jobId < uint64(10) {
		return models.Job{}, errors.New("number is less than 10")
	} */
	jobData, err := s.UserRepo.ViewJobDetailsBy(ctx, jobId)
	if err != nil {
		log.Info().Err(err).Send()
		return models.Job{}, err
	}
	return jobData, nil

}

// ////////////////////////////////////////////////////////////////////////////////
func (s *Service) GetAllJobData(ctx context.Context) ([]models.Job, error) {
	jobDatas, err := s.UserRepo.FindAllJobs(ctx)
	if err != nil {
		log.Info().Err(err).Send()
		return nil, err
	}
	return jobDatas, nil
}

// ///////////////////////////////////////////////////////////////////////////////////////
func (s *Service) GetJobByCompany(ctx context.Context, companyId uint64) ([]models.Job, error) {
	jobData, err := s.UserRepo.FindJob(ctx, companyId)
	if err != nil {
		log.Info().Err(err).Send()
		return nil, err
	}
	return jobData, nil

}
