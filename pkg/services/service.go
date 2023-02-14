package services

import (
	pb "github.com/Egor-Tihonov/go-grpc-auth-service/pkg/pb/auth"
	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/repository"
	userSe "github.com/Egor-Tihonov/go-grpc-auth-service/pkg/user"
	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/utils"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	Rps        *repository.DBPostgres
	Jwt        *utils.JwtWrapper
	UserClient *userSe.ServiceClient
}

func New(r *repository.DBPostgres, jwt *utils.JwtWrapper, client *userSe.ServiceClient) *Server {
	return &Server{
		Rps:        r,
		Jwt:        jwt,
		UserClient: client,
	}
}
