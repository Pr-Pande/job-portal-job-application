package handlers

import (
	"context"
	"errors"
	"job-portal-api/internal/middlewares"
	"job-portal-api/internal/models"
	"job-portal-api/internal/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

func Test_handler_ViewJobByID(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.UserService)
		expectedStatusCode int
		expectedResponse   string
	}{
		// TODO: Add test cases.
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "invalid job id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "job_Id", Value: "abc"})
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "error from service mock function",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "job_Id", Value: "123"})
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockUserService(mc)

				ms.EXPECT().GetJobData(c.Request.Context(), gomock.Any()).Return(models.Job{}, errors.New("test service error")).AnyTimes()
				return c, rr, ms
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Test service error"}`,
		},
		{
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "job_Id", Value: "123"})
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockUserService(mc)

				ms.EXPECT().GetJobData(c.Request.Context(), gomock.Any()).Return(models.Job{}, nil).AnyTimes()
				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"companyId":0,"title":"","minNPInMonths":0,"maxNPInMonths":0,"budget":"","JobLocations":null,"Technologies":null,"WorkModes":null,"minExp":0,"maxExp":0,"Qualifications":null,"Shifts":null,"jobType":"","desc":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()

			h := &handler{
				service: ms,
			}
			h.ViewJobByID(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_AddJob(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.UserService)
		expectedStatusCode int
		expectedResponse   string
	}{
		// TODO: Add test cases.
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "error in decoding json from job request body",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{
					"title" : "senior web developer",
					"minNPInMonths" : 3,
					"maxNPInMonths" : 8,
					"budget" : "$50000"
					"jobLocations" : [
						1,
						2
						],
					"technologyStacks":[
						1,
						2
					],
					"workModes":[
						1,
						2
					],
					"minExp": 1,
					"maxExp": 5,
					"qualifications": [
						1,
						2
					],
					"shifts": [
						1,
						2
					],
					"jobType": "Full-time",
					"desc": "Web designers primarily focus on the visual and user experience aspects of web development. They create mockups, wireframes, and prototypes to communicate design concepts, working closely with web developers to implement designs and maintain a consistent user interface."
				}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "job_Id", Value: "123"})
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "validation error in NewJob struct",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{
					"companyId" : 1,
					"title" : "senior web developer",
					"minNPInMonths" : 3,
					"maxNPInMonths" : 8
				}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "job_Id", Value: "123"})
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusNotFound,
			expectedResponse:   `{"error":"please provide required field data"}`,
		},
		{
			name: "invalid company id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UserService) {
				// Create a valid JSON payload with all required fields
				validPayload := `{
					"title": "Software Engineer",
					"minNPInMonths": 3,
					"maxNPInMonths": 12,
					"budget": "$80000",
					"jobLocations": [1, 2],
					"technologyStacks": [1, 2],
					"workModes": [1, 2],
					"minExp": 2,
					"maxExp": 5,
					"qualifications": [1, 2],
					"shifts": [1, 2],
					"jobType": "Full-time",
					"desc": "Exciting software engineering opportunity"
				}`
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodPost, "http://test.com:8080", strings.NewReader(validPayload))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				
				// Set up the context with a non-numeric company ID
				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "abc"})
				c.Request = httpRequest
		
				return c, rr, nil // Mock implementation
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"error found at conversion"}`,
		},
		{
			name: "error in job storage",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UserService) {
				validPayload := `{
					"title": "Software Engineer",
					"minNPInMonths": 3,
					"maxNPInMonths": 12,
					"budget": "$80000",
					"jobLocations": [1, 2],
					"technologyStacks": [1, 2],
					"workModes": [1, 2],
					"minExp": 2,
					"maxExp": 5,
					"qualifications": [1, 2],
					"shifts": [1, 2],
					"jobType": "Full-time",
					"desc": "Exciting software engineering opportunity"
				}`
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodPost, "http://test.com:8080", strings.NewReader(validPayload))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				
				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "123"})
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockUserService(mc)

				ms.EXPECT().StoreJob(c.Request.Context(), gomock.Any(), gomock.Any()).Return(models.JobRespo{}, errors.New("test service error")).AnyTimes()
				return c, rr, ms
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"job table creation failed"}`,
		},
		{
			name: "successful job addition",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UserService) {
				validPayload := `{
					"title": "Software Engineer",
					"minNPInMonths": 3,
					"maxNPInMonths": 12,
					"budget": "$80000",
					"jobLocations": [1, 2],
					"technologyStacks": [1, 2],
					"workModes": [1, 2],
					"minExp": 2,
					"maxExp": 5,
					"qualifications": [1, 2],
					"shifts": [1, 2],
					"jobType": "Full-time",
					"desc": "Exciting software engineering opportunity"
				}`
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodPost, "http://test.com:8080", strings.NewReader(validPayload))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				
				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "123"})
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockUserService(mc)

				ms.EXPECT().StoreJob(c.Request.Context(), gomock.Any(), gomock.Any()).Return(models.JobRespo{}, nil).AnyTimes()
				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"jobId":0}`,
		},


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()

			h := &handler{
				service: ms,
			}

			h.AddJob(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}
