package services

import (
	"context"
	"errors"
	"job-portal-api/internal/models"
	"sync"

	"github.com/rs/zerolog/log"
)

func (s *Service) ProcessJobApplication(ctx context.Context, applicationData []models.UserApplication) ([]models.ApplRespo, error) {
	var wg = new(sync.WaitGroup)
	var finalData []models.ApplRespo
	for _, application := range applicationData {
		wg.Add(1)
		go func(appl models.UserApplication) {
			defer wg.Done()
			jobData, err := s.UserRepo.GetTheJobData(appl.JobId)  //fetching job data
			if err != nil {
				log.Error().Err(err).Interface("Job ID", appl.JobId).Send()
				return
			}

			if appl.JobApplication.Experience < jobData.MinExp || appl.JobApplication.Experience > jobData.MaxExp {
				log.Error().Err(errors.New("experience requirments not met")).Interface("applicant name", appl.Name).Send()
				return
			}
			if appl.JobApplication.Jobtypes != jobData.JobTypes {
				log.Error().Err(errors.New("jobtype requirments not met")).Interface("applicant name", appl.Name).Send()
				return
			}
			if appl.JobApplication.NoticePeriodInMonths > jobData.MaxNPInMonths {
				log.Error().Err(errors.New("notice period requirments not met")).Interface("applicant name", appl.Name).Send()
				return
			}
			count := 0
			for _, v := range appl.JobApplication.JobLocations {
				for _, v1 := range jobData.JobLocations {
					if v == v1.ID {
						count++
					}
				}
			}
			if count == 0 {
				log.Error().Err(errors.New("location requirments not met")).Interface("applicant name", appl.Name).Send()
				return
			}

			count = 0
			for _, v := range appl.JobApplication.Technologies {
				for _, v1 := range jobData.Technologies {
					if v == v1.ID {
						count++
					}
				}
			}
			if count < (len(appl.JobApplication.Technologies)/2) {
				log.Error().Err(errors.New("technology requirments not met")).Interface("applicant name", appl.Name).Send()
				return
			}

			count = 0
			for _, v := range appl.JobApplication.WorkModes {
				for _, v1 := range jobData.WorkModes {
					if v == v1.ID {
						count++
					}
				}
			}
			if count == 0 {
				log.Error().Err(errors.New("work mode requirments not met")).Interface("applicant name", appl.Name).Send()
				return
			}

			count = 0
			for _, v := range appl.JobApplication.Qualifications {
				for _, v1 := range jobData.Qualifications {
					if v == v1.ID {
						count++
					}
				}
			}
			if count == 0 {
				log.Error().Err(errors.New("qualification requirments not met")).Interface("applicant name", appl.Name).Send()
				return
			}

			count = 0
			for _, v := range appl.JobApplication.Shifts {
				for _, v1 := range jobData.Shifts {
					if v == v1.ID {
						count++
					}
				}
			}
			if count == 0 {
				log.Error().Err(errors.New("shift requirments not met")).Interface("applicant name", appl.Name).Send()
				return
			}
			respo := models.ApplRespo{
				Name:  appl.Name,
				JobId: appl.JobId,
			}

			finalData = append(finalData, respo)
		}(application)

		/* check, v, err := s.compareAndCheck(v)

		if err != nil {
			return nil, err
		}
		if check {
			finalData = append(finalData, v)
		} */
	}
	wg.Wait()
	return finalData, nil
}
