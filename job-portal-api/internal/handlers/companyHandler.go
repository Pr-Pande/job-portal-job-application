package handlers

import (
	"encoding/json"
	"job-portal-api/internal/auth"
	middlewares "job-portal-api/internal/middlewares"

	"job-portal-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

// create comapny table in database
func (h *handler) AddCompany(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		// If the traceId isn't found in the request, log an error and return
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceId).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	var nCompany models.NewCompany

	// Attempt to decode JSON from the request body into the NewUser variable
	err := json.NewDecoder(c.Request.Body).Decode(&nCompany)
	if err != nil {
		// If there is an error in decoding, log the error and return
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	//for validation
	validate := validator.New()
	err = validate.Struct(&nCompany)
	if err != nil {
		// If validation fails, log the error and return
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "please provide Name and Location"})
		return
	}

	//store to database
	cmpny, err := h.service.StoreCompany(ctx, nCompany)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("database creation is not happening")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "company creation in database failed"})
		return

	}
	// If everything goes right, respond with the created user
	c.JSON(http.StatusOK, cmpny)

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//fetch a particular company from database using comapny id it will get from seeing the database

func (h *handler) ViewCompany(c *gin.Context) {
	///context to trace request
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	stringCmpnyId := c.Param("company_id")
	compid, err := strconv.ParseUint(stringCmpnyId, 10, 64)
	if err != nil {

		log.Print("conversion string to int error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "error found at conversion.."})
		return

	}
	val, err := h.service.GetCompanyData(ctx, compid)
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
// to get all the company details
func (h *handler) ViewAllCompanies(c *gin.Context) {
	///context to trace request
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	val, err := h.service.GetAllCompanyData(ctx)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("not able to hit the database")
		log.Print("company table not present in database")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "Company table not there"})
		return
	}
	c.JSON(http.StatusOK, val)
}
