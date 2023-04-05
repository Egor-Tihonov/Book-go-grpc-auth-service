package utils

import (
	"time"

	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/models"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

type JwtWrapper struct {
	SecretKey       string
}

func (w *JwtWrapper) GenerateToken(user *models.User) (accessToken string, err error) {
	claims := &models.JWTClaims{
		ID:    user.ID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	accessToken, err = token.SignedString([]byte(w.SecretKey))
	if err != nil {
		logrus.Errorf("auth service: Failed generate token, %w", err)
		return
	}

	return
}
