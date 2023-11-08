package services

import (
	"context"
	"fmt"
	"job-portal-api/internal/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

// Authenticate is a method that checks a user's provided email and password against the database.
func (s *Service) UserSignup(ctx context.Context, newUser models.NewUser) (models.User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Info().Err(err).Send()
		return models.User{}, fmt.Errorf("generating password hash: %w", err)
	}

	// We prepare the User record.
	userDetails := models.User{
		Name:         newUser.Name,
		Email:        newUser.Email,
		PasswordHash: string(hashedPass),
	}
	userDetails, err = s.UserRepo.CreateUser(ctx, userDetails)

	if err != nil {
		log.Info().Err(err).Send()
		return models.User{}, err
	}

	// Successfully created the record, return the user.
	return userDetails, nil

}

// Authenticate is a method that checks a user's provided email and password against the database.
func (s *Service) UserLogin(ctx context.Context, userData models.NewUser) (string, error) {

	// We attempt to find the User record where the email
	// matches the provided email.
	var userDetails models.User
	userDetails, err := s.UserRepo.CheckEmail(ctx, userData.Email)
	if err != nil {
		log.Info().Err(err).Send()
		return "", err
	}

	// comparing the password and hashed password
	err = bcrypt.CompareHashAndPassword([]byte(userDetails.PasswordHash), []byte(userData.Password))
	if err != nil {
		log.Info().Err(err).Send()
		return "", err
	}

	// setting up the claims
	claims := jwt.RegisteredClaims{
		Issuer:    "job portal project",
		Subject:   strconv.FormatUint(uint64(userDetails.ID), 10),
		Audience:  jwt.ClaimStrings{"users"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token, err := s.auth.GenerateToken(claims)
	if err != nil {
		log.Info().Err(err).Send()
		return "", err
	}

	return token, nil

}
