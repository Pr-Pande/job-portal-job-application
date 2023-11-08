package services

import (
	"context"
	"job-portal-api/internal/models"
	"sync"
)

func (s *Service) ProcessJobApplication(ctx context.Context, applicationData []models.UserApplication) ([]models.UserApplication, error) {
	var wg = new(sync.WaitGroup)
	var finalData []models.UserApplication
	for _, v := range applicationData {
		wg.Add(1)
		go func(v models.UserApplication) {
			defer wg.Done()
			check, v, err := s.compareAndCheck(v)
			if err != nil {
				return
			}
			if check {
				finalData = append(finalData, v)
			}
		}(v)

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

var cacheMap = make(map[uint]models.Job)

func (s *Service) compareAndCheck(applicationDetails models.UserApplication) (bool, models.UserApplication, error) {
	val, exists := cacheMap[applicationDetails.JobId]
	if !exists {
		jobData, err := s.UserRepo.GetTheJobData(applicationDetails.JobId)
		if err != nil {
			return false, models.UserApplication{}, err
		}
		cacheMap[applicationDetails.JobId] = jobData
		val = jobData
	}
	if applicationDetails.JobApplication.Experience < val.MinExp || applicationDetails.JobApplication.Experience > val.MaxExp {
		return false, models.UserApplication{}, nil
	}
	if applicationDetails.JobApplication.Jobtypes != val.JobTypes {
		return false, models.UserApplication{}, nil
	}
	if applicationDetails.JobApplication.NoticePeriod < val.MinNP || applicationDetails.JobApplication.NoticePeriod > val.MaxNP {
		return false, models.UserApplication{}, nil
	}
	count := 0
	for _, v := range applicationDetails.JobApplication.JobLocations {
		for _, v1 := range val.JobLocations {
			if v == v1.ID {
				count++
			}
		}
	}
	if count == 0 {
		return false, models.UserApplication{}, nil
	}

	count = 0
	for _, v := range applicationDetails.JobApplication.Technologies {
		for _, v1 := range val.Technologies {
			if v == v1.ID {
				count++
			}
		}
	}
	if count == 0 {
		return false, models.UserApplication{}, nil
	}

	count = 0
	for _, v := range applicationDetails.JobApplication.WorkModes {
		for _, v1 := range val.WorkModes {
			if v == v1.ID {
				count++
			}
		}
	}
	if count == 0 {
		return false, models.UserApplication{}, nil
	}

	count = 0
	for _, v := range applicationDetails.JobApplication.Qualifications {
		for _, v1 := range val.Qualifications {
			if v == v1.ID {
				count++
			}
		}
	}
	if count == 0 {
		return false, models.UserApplication{}, nil
	}

	count = 0
	for _, v := range applicationDetails.JobApplication.Shifts {
		for _, v1 := range val.Shifts {
			if v == v1.ID {
				count++
			}
		}
	}
	if count == 0 {
		return false, models.UserApplication{}, nil
	}

	return true, applicationDetails, nil
}
