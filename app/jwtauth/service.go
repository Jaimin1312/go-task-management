package jwtauth

import (
	"task-management/consts"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	Validatejwt(token string) (*Claims, error)
	CreateJWTToken(userID string, tokenExpiration time.Duration, JWTKey string) (*JWTToken, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Validatejwt(token string) (*Claims, error) {
	claims, err := validatejwt(token, consts.JwtKey)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (s *service) CreateJWTToken(userID string, tokenExpiration time.Duration, JWTKey string) (*JWTToken, error) {
	expirationTime := time.Now().Add(tokenExpiration * time.Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(JWTKey))
	if err != nil {
		return nil, err
	}
	return &JWTToken{
		Value:     tokenString,
		ExpiresAt: expirationTime,
	}, nil
}
