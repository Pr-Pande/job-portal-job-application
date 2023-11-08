package handlers
/*
import (
	"bytes"
	"context"
	"encoding/json"
	middlewares "jobportalapi/internal/middlewares"
	"jobportalapi/internal/models"
	"jobportalapi/internal/services"
	"jobportalapi/internal/services/mockmodels"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestUserSignup(t *testing.T) {

	nu := models.NewUser{
		Name:     "Pragalbh Pande",
		Email:    "pragalbh@gmail.com",
		Password: "pass123",
	}

	mockuser := models.User{
		Model:        gorm.Model{ID: 1},
		Name:         "Pragalbh Pande",
		Email:        "pragalbh@gmail.com",
		PasswordHash: "jskskslsms",
	}

	tt := [...]struct {
		name             string
		body             any
		expectedStatus   int
		expectedResponse string
		expectedUser     models.User
		mockUserService  func(m *mockmodels.MockService)
	}{
		{
			name:             "ok",
			body:             nu,
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"jeevan","email":"jeevan@gmail.com"}`,
			expectedUser:     mockuser,
			//set expection inside this field
			mockUserService: func(m *mockmodels.MockService) {
				m.EXPECT().CreateUser(gomock.Any(), gomock.Eq(nu)).
					Times(1).Return(mockuser, nil)
			},
		},
		{

			name: "Fail_NoEmail",
			body: models.NewUser{
				Name:     "testuser",
				Password: "password",
			},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `{"msg":"please provide Name, Email and Password"}`,
			mockUserService: func(m *mockmodels.MockService) {
				m.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Times(1)
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Gomock controller.
			ctrl := gomock.NewController(t)
			//this func give us the mock implementation of the interface
			mockService := mockmodels.NewMockService(ctrl)
			s := services.NewStore(mockService)

			// Apply the mock to the user service.
			tc.mockUserService(mockService)

			router := gin.New()
			h := handler{
				S: s,
			}
			ctx := context.Background()
			traceID := "fake-trace-id"
			ctx = context.WithValue(ctx, middlewares.TraceIdKey, traceID)

			//register endpoints
			router.POST("/signUp", h.UserSignup)
			body, err := json.Marshal(tc.body)
			require.NoError(t, err)
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/userSignup", bytes.NewReader(body))
			require.NoError(t, err)

			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			require.Equal(t, tc.expectedStatus, rec.Code)
			require.Equal(t, tc.expectedResponse, rec.Body.String())
		})
	}
}
*/