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

func TestService_ProcessJobApplication(t *testing.T) {
	type args struct {
		ctx             context.Context
		applicationData []models.UserApplication
	}
	tests := []struct {
		name             string
		args             args
		want             []models.ApplRespo
		wantErr          bool
		mockRepoResponse func() (models.Job, error)
	}{
		// TODO: Add test cases.
		{
			name: "error from mock function",
			args: args{
				ctx: context.Background(),
				applicationData: []models.UserApplication{
					{
						Name:           "Pragalbh Pande",
						Age:            23,
						JobId:          1,
						JobApplication: models.Application{
							Title: "Go developer",
							NoticePeriodInMonths: 1,
							JobLocations:      []uint{1, 2,},
							Technologies:      []uint{1, 2,},
							WorkModes:         []uint{1, 2,},
							Experience:     6,
							Qualifications: []uint{1, 2,},
							Shifts:         []uint{1, 2,},
							Jobtypes:       "Full-time",
						},
					},
					{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := database.NewMockUserRepo(mc)

			mockRepo.EXPECT().GetTheJobData(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()

			s, err := NewService(mockRepo, &auth.Auth{})
			if err != nil {
				t.Errorf("error is initializing the repo layer")
				return
			}

			got, err := s.ProcessJobApplication(tt.args.ctx, tt.args.applicationData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ProcessJobApplication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ProcessJobApplication() = %v, want %v", got, tt.want)
			}
		})
	}
}
