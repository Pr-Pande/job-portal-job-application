package handlers

import (
	"encoding/json"
	middlewares "job-portal-api/internal/middlewares"
	"job-portal-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func (h *handler) AddJob(c *gin.Context) {

	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		// If the traceId isn't found in the request, log an error and return
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	var nJob models.NewJob
	// Attempt to decode JSON from the request body into the NewUser variable
	err := json.NewDecoder(c.Request.Body).Decode(&nJob)
	if err != nil {
		// If there is an error in decoding, log the error and return
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	//for validation

	validate := validator.New()
	err = validate.Struct(&nJob)
	if err != nil {
		// If validation fails, log the error and return
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "please provide Name and Location"})
		return
	}
	stringCmpnyId := c.Param("company_id")
	compid, err := strconv.ParseUint(stringCmpnyId, 10, 64)
	if err != nil {

		log.Print("conversion string to int error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "error found at conversion.."})
		return

	}
	///////////////////////////////////////////////////////
	//store to database
	job, err := h.service.StoreJob(ctx, nJob, compid)

	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("job table creation problem")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "job table creation failed"})
		return
	}
	// If everything goes right, respond with the created user
	c.JSON(http.StatusOK, job)

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
//fetch a particular company from database using comapny id it will get from seeing the database

func (h *handler) ViewJobByID(c *gin.Context) {
	///context to trace request
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	stringJobId := c.Param("job_Id")
	jobid, err := strconv.ParseUint(stringJobId, 10, 64)
	if err != nil {

		log.Print("conversion string to int error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "error found at conversion.."})
		return

	}
	val, err := h.service.GetJobData(ctx, jobid)
	if err != nil {
		if err != nil {
			log.Error().Err(err).Str("Trace Id", traceId).Msg("not able to hit the database")
			log.Print("company data not found in database %w", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "company record not found"})
			return
		}

	}
	c.JSON(http.StatusOK, val)
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////
// to get all the job details
func (h *handler) ViewAllJobs(c *gin.Context) {
	///context to trace request
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	val, err := h.service.GetAllJobData(ctx)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("not able to hit the database")
		log.Print("job table not present in database")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "job  table not there"})
		return
	}
	c.JSON(http.StatusOK, val)
}

///////////////////////////////////////////////////////////////////////////////////////
// to get all the job details corresponding to the company

func (h *handler) ViewJobByCompid(c *gin.Context) {
	///context to trace request
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	stringcompId := c.Param("company_id")
	compid, err := strconv.ParseUint(stringcompId, 10, 64)
	if err != nil {

		log.Print("conversion string to int error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "error found at conversion.."})
		return

	}
	val, err := h.service.GetJobByCompany(ctx, compid)
	if err != nil {
		if err != nil {
			log.Error().Err(err).Str("Trace Id", traceId).Msg("not able to hit the database")
			log.Print("company data not found in database %w", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "company record not found"})
			return
		}

	}
	c.JSON(http.StatusOK, val)
}

/* func (h *handler) ProcessedJobAppl(c *gin.Context) {

} */