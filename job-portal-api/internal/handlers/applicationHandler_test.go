package handlers

import (
	"context"
	"errors"
	"job-portal-api/internal/auth"
	"job-portal-api/internal/middlewares"
	"job-portal-api/internal/models"
	"job-portal-api/internal/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/mock/gomock"
)

func Test_handler_ProcessedJobAppl(t *testing.T) {
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
			name: "missing jwt claims",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   `{"error":"Unauthorized"}`,
		},
		{
			name: "error in decoding json from job application body",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UserService) {
				// Create a valid JSON payload with all required fields
				validPayload := `{
					"name": "Vikram Gupta",
					"age": 24,
					"jobId": 4,  
					"jobApplication": {
						"title": "Software Engineer",
						"noticePeriodInMonths": 3,
						"location": [1, 2],
						"technologyStacks": [1, 2],
						"workModes": [1, 2],
						"experience": 6,
						"qualifications": [1, 2],  
						"shifts": [1, 2], 
						"jobType": "Full-time"
						}
					},
					{
						"name": "Shubh Pande",
						"age": 30,
						"jobId": 4,  
						"jobApplication": {
							"title": "Software Engineer",
							"noticePeriodInMonths": 3,
							"location": [1, 2],
							"technologyStacks": [1, 2],
							"workModes": [1, 2],
							"experience": 7,
							"qualifications": [1, 2],  
							"shifts": [1, 2], 
							"jobType": "Full-time"
						}
					}`
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodPost, "http://test.com:8080", strings.NewReader(validPayload))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)

				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"please provide valid details"}`,
		},
		{
			name: "error in processing job application",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UserService) {
				validPayload := `
				[
					{
						"name": "Vikram Gupta",
						"age": 24,
						"jobId": 4,  
						"jobApplication": 
						{
							"title": "Software Engineer",
							"noticePeriodInMonths": 3,
							"location": [1, 2],
							"technologyStacks": [1, 2],
							"workModes": [1, 2],
							"experience": 6,
							"qualifications": [1, 2],  
							"shifts": [1, 2], 
							"jobType": "Full-time"
						}
					},
					{
						"name": "Shubh Pande",
						"age": 30,
						"jobId": 4,  
						"jobApplication": 
						{
							"title": "Software Engineer",
							"noticePeriodInMonths": 3,
							"location": [1, 2],
							"technologyStacks": [1, 2],
							"workModes": [1, 2],
							"experience": 7,
							"qualifications": [1, 2],  
							"shifts": [1, 2], 
							"jobType": "Full-time"
						}
					}
				]`
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodPost, "http://test.com:8080", strings.NewReader(validPayload))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockUserService(mc)

				ms.EXPECT().ProcessJobApplication(c.Request.Context(), gomock.Any()).Return([]models.ApplRespo{}, errors.New("test service error")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"processing job application failed"}`,
		},
		{
			name: "success in processing job application",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UserService) {
				validPayload := `
				[
					{
						"name": "Vikram Gupta",
						"age": 24,
						"jobId": 4,  
						"jobApplication": 
						{
							"title": "Software Engineer",
							"noticePeriodInMonths": 3,
							"location": [1, 2],
							"technologyStacks": [1, 2],
							"workModes": [1, 2],
							"experience": 6,
							"qualifications": [1, 2],  
							"shifts": [1, 2], 
							"jobType": "Full-time"
						}
					},
					{
						"name": "Shubh Pande",
						"age": 30,
						"jobId": 4,  
						"jobApplication": 
						{
							"title": "Software Engineer",
							"noticePeriodInMonths": 3,
							"location": [1, 2],
							"technologyStacks": [1, 2],
							"workModes": [1, 2],
							"experience": 7,
							"qualifications": [1, 2],  
							"shifts": [1, 2], 
							"jobType": "Full-time"
						}
					}
				]`
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodPost, "http://test.com:8080", strings.NewReader(validPayload))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockUserService(mc)

				ms.EXPECT().ProcessJobApplication(c.Request.Context(), gomock.Any()).Return([]models.ApplRespo{
					{
						Name: "Vikram Gupta",
						JobId: 4,
					},
					{
						Name: "Shubh Pande",
						JobId: 4,
					},
				}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[{"name":"Vikram Gupta","jobId":4},{"name":"Shubh Pande","jobId":4}]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()

			h := &handler{
				service: ms,
			}
			h.ProcessedJobAppl(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}
