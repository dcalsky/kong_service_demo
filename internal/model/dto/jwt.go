package dto

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	AccountId uint
	Email     string
	ExpiredAt time.Time
}

func (m JwtClaims) GetExpirationTime() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{Time: m.ExpiredAt}, nil
}

func (m JwtClaims) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil
}

func (m JwtClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	return nil, nil
}

func (m JwtClaims) GetAudience() (jwt.ClaimStrings, error) {
	return nil, nil
}

func (m JwtClaims) GetIssuer() (string, error) {
	return "", nil
}

func (m JwtClaims) GetSubject() (string, error) {
	return "", nil
}
