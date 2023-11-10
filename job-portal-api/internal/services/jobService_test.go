package services

import (
	"context"
	"errors"
	"job-portal-api/internal/auth"
	database "job-portal-api/internal/database/mockFiles"
	"job-portal-api/internal/models"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestService_StoreJob(t *testing.T) {
	type args struct {
		ctx       context.Context
		jobData   models.NewJob
		companyId uint64
	}
	tests := []struct {
		name             string
		args             args
		want             models.JobRespo
		wantErr          bool
		mockRepoResponse func() (models.JobRespo, error)
	}{
		// TODO: Add test cases.
		{
			name: "error from repository mock function",
			want: models.JobRespo{},
			args: args{
				ctx: context.Background(),
				jobData: models.NewJob{
					Title:         "random",
					MinNPInMonths: 0,
					MaxNPInMonths: 3,
					Budget:        "1000",
					JobLocations: []uint{
						1, 2,
					},
					Technologies: []uint{
						1, 2,
					},
					WorkModes: []uint{
						1, 2,
					},
					MinExp: 5,
					MaxExp: 10,
					Qualifications: []uint{
						1, 2,
					},
					Shifts: []uint{
						1, 2,
					},
					JobTypes: "no type",
					Desc:     "desc",
				},
				companyId: 1,
			},
			wantErr: true,
			mockRepoResponse: func() (models.JobRespo, error) {
				return models.JobRespo{}, errors.New("test case error from mock function")
			},
		},
		{
			name: "success case",
			want: models.JobRespo{
				JobId: 3,
			},
			args: args{
				ctx: context.Background(),
				jobData: models.NewJob{
					Title:         "random",
					MinNPInMonths: 0,
					MaxNPInMonths: 3,
					Budget:        "1000",
					JobLocations: []uint{
						1, 2,
					},
					Technologies: []uint{
						1, 2,
					},
					WorkModes: []uint{
						1, 2,
					},
					MinExp: 5,
					MaxExp: 10,
					Qualifications: []uint{
						1, 2,
					},
					Shifts: []uint{
						1, 2,
					},
					JobTypes: "no type",
					Desc:     "desc",
				},
				companyId: 1,
			},
			wantErr: false,
			mockRepoResponse: func() (models.JobRespo, error) {
				return models.JobRespo{
					JobId: 3,
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := database.NewMockUserRepo(mc)
			mockRepo.EXPECT().CreateJob(gomock.Any(), gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()

			s, err := NewService(mockRepo, &auth.Auth{})
			if err != nil {
				t.Errorf("error is initializing the repo layer")
				return
			}

			got, err := s.StoreJob(tt.args.ctx, tt.args.jobData, tt.args.companyId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.StoreJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.StoreJob() = %v, want %v", got, tt.want)
			}
		})
	}
}
