package services

import (
	"github.com/Egor-Tihonov/go-grpc-auth-service.git/pkg/repository"
	"github.com/Egor-Tihonov/go-grpc-auth-service.git/pkg/utils"
)

type Service struct {
	Rps *repository.DBPostgres
	Jwt *utils.JwtWrapper
}

func New(r *repository.DBPostgres, jwt *utils.JwtWrapper) *Service {
	return &Service{
		Rps: r,
		Jwt: jwt,
	}
}
