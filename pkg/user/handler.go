package user

import (
	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/config"
	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/models"
	pb "github.com/Egor-Tihonov/go-grpc-auth-service/pkg/pb/user"
	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/user/handlers"
)

func RegisterHandlers(conf config.Config) *ServiceClient {
	svc := ServiceClient{
		Cliet: InitUserClient(&conf),
	}
	return &svc
}

func (s *ServiceClient) CreateUser(user *models.UserCreate) (*pb.CreateUserResponse, error) {
	return handlers.CreateUser(user, s.Cliet)
}
