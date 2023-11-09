package handlers

import (
	"errors"
	"fmt"
	"job-portal-api/internal/auth"
	middlewares "job-portal-api/internal/middlewares"
	"job-portal-api/internal/services"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service services.UserService
}

type UserHandler interface {
	SignIn(c *gin.Context)
	SignUp(c *gin.Context)

	AddCompany(c *gin.Context)
	ViewCompany(c *gin.Context)
	ViewAllCompanies(c *gin.Context)

	ViewJobByID(c *gin.Context)
	ViewAllJobs(c *gin.Context)
	ViewJobByCompid(c *gin.Context)
	AddJob(c *gin.Context)
	ProcessedJobAppl(c *gin.Context)
}

func NewHandler(s services.UserService) (UserHandler, error) {
	if s == nil {
		return nil, errors.New("service interface cannot be null")
	}
	return &handler{
		service: s,
	}, nil
}

func SetupAPI(a auth.Authentication, ser services.UserService) *gin.Engine {

	// Create a new Gin engine; Gin is a HTTP web framework written in Go
	r := gin.New()
	m, err := middlewares.NewMid(a)
	if err != nil {
		log.Panic().Msg("middlewares not set up")

	}

	h, err := NewHandler(ser)
	if err != nil {
		log.Panic().Msg("handlers not setup")
	}

	r.Use(m.Log(), gin.Recovery())
	//Users
	r.GET("/api/check", check)
	r.POST("/api/register", h.SignUp)
	r.POST("/api/login", h.SignIn)

	//Company
	r.POST("/api/companies", m.Authenticate(h.AddCompany))
	r.GET("/api/companies", m.Authenticate(h.ViewAllCompanies))
	r.GET("/api/companies/:company_id", m.Authenticate(h.ViewCompany))

	//Job
	r.POST("/api/companies/:company_id/jobs", m.Authenticate(h.AddJob))
	r.GET("/api/jobs/:company_id", m.Authenticate(h.ViewJobByCompid))
	r.GET("/api/jobs", m.Authenticate(h.ViewAllJobs))
	r.GET("/api/job/:job_Id", m.Authenticate(h.ViewJobByID))
	r.POST("/api/job/application", m.Authenticate(h.ProcessedJobAppl))

	return r

}

func check(c *gin.Context) {
	//handle panic using recovery function when happening in separate goroutine
	//go func() {
	//	panic("some kind of panic")
	//}()
	time.Sleep(time.Second * 3)
	select {
	case <-c.Request.Context().Done():
		fmt.Println("user not there")
		return
	default:
		c.JSON(http.StatusOK, gin.H{"msg": "statusOk"})

	}

}
