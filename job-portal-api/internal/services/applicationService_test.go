package services

import (
	"context"
	"job-portal-api/internal/models"
	"reflect"
	"testing"
)

func TestService_ProcessJobApplication(t *testing.T) {
	type args struct {
		ctx             context.Context
		applicationData []models.UserApplication
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		want    []models.ApplRespo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ProcessJobApplication(tt.args.ctx, tt.args.applicationData)
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
