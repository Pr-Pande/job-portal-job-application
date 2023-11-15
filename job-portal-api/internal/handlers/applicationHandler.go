package handlers

import (
	"encoding/json"
	"job-portal-api/internal/auth"
	"job-portal-api/internal/middlewares"
	"job-portal-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func (h *handler) ProcessedJobAppl(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceId).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	var applicationData []models.UserApplication

	err := json.NewDecoder(c.Request.Body).Decode(&applicationData)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "please provide valid details"})
		return
	}

	filteredData, err := h.service.ProcessJobApplication(ctx, applicationData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "processing job application failed",
		})
		return
	}

	c.JSON(http.StatusOK, filteredData)

}
