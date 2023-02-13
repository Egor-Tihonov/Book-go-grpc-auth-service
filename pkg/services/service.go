package services

import (
	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/pb"
	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/repository"
	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/utils"
)

type Service struct {
	pb.UnimplementedAuthServiceServer
	Rps *repository.DBPostgres
	Jwt *utils.JwtWrapper
}

func New(r *repository.DBPostgres, jwt *utils.JwtWrapper) *Service {
	return &Service{
		Rps: r,
		Jwt: jwt,
	}
}
