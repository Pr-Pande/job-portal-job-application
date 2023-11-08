package auth

import (
	"crypto/rsa"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type ctxKey int

const Key ctxKey = 1

type Auth struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

type Authentication interface {
	GenerateToken(claims jwt.RegisteredClaims) (string, error)
	ValidateToken(token string) (jwt.RegisteredClaims, error)
}


func NewAuth(privateKey *rsa.PrivateKey, publickey *rsa.PublicKey) (Authentication, error) {
	if privateKey == nil || publickey == nil {
		return nil, errors.New("private and public key cannot be nil")

	}
	return &Auth{
		privateKey: privateKey,
		publicKey:  publickey,
	}, nil

}

