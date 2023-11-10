package services

import (
	"context"
	"errors"
	"job-portal-api/internal/models"
	"sync"

	"github.com/rs/zerolog/log"
)

var cacheMap = make(map[uint]models.Job)

func (s *Service) ProcessJobApplication(ctx context.Context, applicationData []models.UserApplication) ([]models.ApplRespo, error) {
	var wg = new(sync.WaitGroup)
	ch := make(chan models.UserApplication)
	var finalData []models.ApplRespo
	for _, application := range applicationData {
		wg.Add(1)
		go func(appl models.UserApplication) {
			defer wg.Done()
			val, exists := cacheMap[appl.JobId]

			if !exists{
				jobData, err := s.UserRepo.GetTheJobData(appl.JobId)  //fetching job data
				if err != nil {
					log.Error().Err(err).Interface("Job ID", appl.JobId).Send()
					return
				}
				cacheMap[appl.JobId] = jobData
				val = jobData
			}

			check := compareAndCheck(appl, val)
			if check {
				ch <- appl
			}

		}(application)

		/* check, v, err := s.compareAndCheck(v)

		if err != nil {
			return nil, err
		}
		if check {
			finalData = append(finalData, v)
		} */
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for data := range ch{
		respo := models.ApplRespo{
			Name:  data.Name,
			JobId: data.JobId,
		}
		finalData = append(finalData, respo)
	}
	if finalData == nil {
		log.Error().Err(errors.New("no candidates passed the requirments")).Send()
		return nil, errors.New("no candidates passed the requirments")
	}
	return finalData, nil
}

func compareAndCheck(appl models.UserApplication, jobData models.Job) bool {
	if appl.JobApplication.Experience < jobData.MinExp || appl.JobApplication.Experience > jobData.MaxExp {
		log.Error().Err(errors.New("experience requirments not met")).Interface("applicant name", appl.Name).Send()
		return false
	}
	if appl.JobApplication.Jobtypes != jobData.JobTypes {
		log.Error().Err(errors.New("jobtype requirments not met")).Interface("applicant name", appl.Name).Send()
		return false
	}
	if appl.JobApplication.NoticePeriodInMonths > jobData.MaxNPInMonths {
		log.Error().Err(errors.New("notice period requirments not met")).Interface("applicant name", appl.Name).Send()
		return false
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
		return false
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
		return false
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
		return false
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
		return false
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
		return false
	}

	return true
}
