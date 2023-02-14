package user

import (
	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/config"
	pb "github.com/Egor-Tihonov/go-grpc-auth-service/pkg/pb/user"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Cliet pb.UserServiceClient
}

func InitUserClient(conf *config.Config) pb.UserServiceClient {
	cc, err := grpc.Dial(conf.UserService, grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("Could not connect: %w", err)
	}

	return pb.NewUserServiceClient(cc)
}
