package handlers

import (
	"encoding/json"
	middlewares "job-portal-api/internal/middlewares"
	"job-portal-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func (h *handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		// If the traceId isn't found in the request, log an error and return
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	var nu models.NewUser

	// Attempt to decode JSON from the request body into the NewUser variable
	err := json.NewDecoder(c.Request.Body).Decode(&nu)
	if err != nil {
		// If there is an error in decoding, log the error and return
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	validate := validator.New()
	err = validate.Struct(nu)
	if err != nil {
		// If validation fails, log the error and return
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "please provide Name, Email and Password"})
		return
	}

	usr, err := h.service.UserSignup(ctx, nu)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("user signup problem")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "user signup failed"})
		return
	}
	// If everything goes right, respond with the created user
	c.JSON(http.StatusOK, usr)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (h *handler) SignIn(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	// Define a new struct for login data
	var usrData models.NewUser
	// Attempt to decode JSON from the request body into the login variable
	err := json.NewDecoder(c.Request.Body).Decode(&usrData)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	// Create a new validator and validate the login variable
	validate := validator.New()
	err = validate.Struct(usrData)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "please provide Email and Password"})
		return
	}

	tkn, err := h.service.UserLogin(ctx, usrData)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "login failed"})
		return
	}

	// If everything goes right, respond with the token
	c.JSON(http.StatusOK, gin.H{
		"token": tkn,
	})
}
