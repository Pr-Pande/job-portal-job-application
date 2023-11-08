package database

/* import (
	"context"
	"errors"
	"job-portal-api/internal/models"
	"sync"

	"github.com/rs/zerolog/log"
)

func (r *Repo) ApplicationFilter(ctx context.Context, applicationData []models.UserApplication) ([]models.Application, error) {
	var job models.Job
	var applicant []models.Application
	ch := make(chan *models.Application)
	wg := &sync.WaitGroup{}
	for _, data := range applicationData {
		tx := r.DB.WithContext(ctx).Where("ID = ?", data.JobId)
		err := tx.First(&job).Error
		if err != nil {
			log.Error().Err(err).Str("Error", "Job Id not found in db").Send()
			return nil, err
		}
		wg.Add(1)
		go func(job models.Job, data models.UserApplication) {
			defer wg.Done()
			// for _, q := range job.Qualifications {
			// 	if q == data.Qualifications {
			// 		break
			// 	}
			// }
			if !(job.MinExp <= data.Ex && data.Experience <= job.Experience) {
				log.Error().Err(errors.New("experience requirments not met")).Send()
				return
			} else if data.Budget > job.Budget {
				log.Error().Err(errors.New("salary requirments not met")).Send()
				return
			} else if !(job.Max_NP >= data.Max_NP && job.Min_NP >= data.Min_NP) {
				log.Error().Err(errors.New("notice periode requirments not met")).Send()
				return
			} else if job.WorkMode != data.WorkMode {
				log.Error().Err(errors.New("work mode requirments not met")).Send()
				return
			} else if job.Shift != data.Shift {
				log.Error().Err(errors.New("shift requirments not met")).Send()
				return
			}
			// for _,s := range job.Stack {
			// 	for _,as := range data.Stack{

			// 	}
			// }
			ch <- data
		}(job, data)

		go func() {
			wg.Wait()
			close(ch)
		}()

		for appl := range ch {
			applicant = append(applicant, &models.ApplicantRespo{
				Id:    appl.ID,
				Name:  appl.Name,
				JobId: appl.JobId,
			})
		}
	}

	return applicant, nil
} */
