package services

import (
	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/db"
	userSe "github.com/Egor-Tihonov/go-grpc-auth-service/pkg/user"
	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/utils"
)

type Server struct {
	rps        *db.DBPostgres
	Jwt        *utils.JwtWrapper
	UserClient *userSe.ServiceClient
}

func New(r *db.DBPostgres, jwt *utils.JwtWrapper, client *userSe.ServiceClient) *Server {
	return &Server{
		rps:        r,
		Jwt:        jwt,
		UserClient: client,
	}
}
