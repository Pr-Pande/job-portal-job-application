package middlewares

import (
	"errors"
	"job-portal-api/internal/auth"
)

type Mid struct {
	a auth.Authentication
}

func NewMid(a auth.Authentication) (Mid, error) {
	if a == nil {
		return Mid{}, errors.New("auth can't be nil")
	}

	return Mid{a: a}, nil
}
