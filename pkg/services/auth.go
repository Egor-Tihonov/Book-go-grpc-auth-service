package services

import (
	"context"
	"net/http"

	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/models"
	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/pb"
	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/utils"
	"github.com/google/uuid"
)

func (s *Service) Registration(ctx context.Context, req *pb.RegistrationRequest) (*pb.Response, error) {
	newID := uuid.New().String()

	user := models.User{
		ID:    newID,
		Email: req.Email,
	}

	hashpassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return &pb.Response{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	user.Password = hashpassword

	if err = s.Rps.CreateUser(ctx, &user); err != nil {
		return &pb.Response{
			Status: http.StatusNotFound,
			Error:  "E-mail already exist",
		}, nil
	}

	return &pb.Response{
		Status: http.StatusCreated,
	}, nil
}

func (s *Service) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.Rps.GetUser(ctx, req.Authstring)
	if err != nil {
		return &pb.LoginResponse{
			Response: &pb.Response{
				Status: http.StatusBadGateway,
				Error:  "User not found",
			},
			Token: "",
		}, nil
	}

	err = utils.CheckPassword(req.Password, user.Password)
	if err != nil {
		return &pb.LoginResponse{
			Response: &pb.Response{
				Status: http.StatusBadGateway,
				Error:  "Password is incorrect",
			},
			Token: "",
		}, nil
	}

	token, err := s.Jwt.GenerateToken(user)
	if err != nil {
		return &pb.LoginResponse{
			Response: &pb.Response{
				Status: http.StatusBadGateway,
				Error:  "Token doesnt generate",
			},
			Token: "",
		}, nil
	}

	return &pb.LoginResponse{
		Response: &pb.Response{
			Status: http.StatusOK,
		},
		Token: token,
	}, nil
}
