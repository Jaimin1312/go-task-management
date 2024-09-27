package jwtauth

import (
	"task-management/apperror"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims - a struct that will be encoded to JWT
type Claims struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}

// JWTToken - JWT Token
type JWTToken struct {
	Value     string
	ExpiresAt time.Time
}

func validatejwt(tokenStr string, jwtKey string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, apperror.ErrUnauthorized.Customize("the JWT Token is invalid").LogWithLocation()
	}

	return claims, nil
}
