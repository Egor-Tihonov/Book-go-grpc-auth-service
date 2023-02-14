package handlers

import (
	"context"

	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/models"
	pb "github.com/Egor-Tihonov/go-grpc-auth-service/pkg/pb/user"
)

func CreateUser(user *models.UserCreate, uscl pb.UserServiceClient) (*pb.CreateUserResponse, error) {
	res, err := uscl.CreateUser(context.Background(), &pb.CreateUserRequest{
		Id:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Username: user.UserName,
	})
	if err != nil {
		return res, err
	}

	return res, err
}
