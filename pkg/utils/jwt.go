package utils

import (
	"time"

	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/models"
	"github.com/golang-jwt/jwt"
)

type JwtWrapper struct {
	SecretKey       string
	ExpirationHours time.Time
}

func (w *JwtWrapper) GenerateToken(user *models.User) (accessToken string, err error) {
	claims := &models.JWTClaims{
		ID:    user.ID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: w.ExpirationHours.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	accessToken, err = token.SignedString(w.SecretKey)
	if err != nil {
		return
	}

	return
}
