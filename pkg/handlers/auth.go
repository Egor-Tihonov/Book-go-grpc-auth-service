package handlers

import (
	"context"
	"net/http"

	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/models"
	pb "github.com/Egor-Tihonov/go-grpc-auth-service/pkg/pb/auth"
)

func (h *Handler) Registration(ctx context.Context, req *pb.RegistrationRequest) (*pb.Response, error) {
	user := models.User{
		Email:    req.Email,
		Password: req.Password,
	}

	id, err := h.se.Registration(ctx, &user)
	if err != nil {
		return &pb.Response{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, err
	}

	res, err := h.se.UserClient.CreateUser(&models.UserCreate{
		ID:       id,
		Email:    req.Email,
		Name:     req.Name,
		UserName: req.Username,
	})
	if err != nil {
		return &pb.Response{
			Status: res.Status,
			Error:  res.Error,
		}, err
	}

	return &pb.Response{
		Status: http.StatusCreated,
	}, nil
}

func (h *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := h.se.GetUser(ctx, req.Authstring)
	if err != nil {
		return &pb.LoginResponse{
			Response: &pb.Response{
				Status: http.StatusBadGateway,
				Error:  "User not found",
			},
			Token: "",
		}, err
	}

	token, err := h.se.Generate(ctx, req.Password, user)
	if err != nil {
		return &pb.LoginResponse{
			Response: &pb.Response{
				Status: http.StatusBadGateway,
				Error:  "User not found",
			},
			Token: "",
		}, err
	}

	return &pb.LoginResponse{
		Response: &pb.Response{
			Status: http.StatusOK,
		},
		Token: token,
	}, nil
}

func (h *Handler) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.Response, error) {
	body := models.UserUpdate{
		Id:          req.Id,
		OldPassword: req.Oldpassword,
		NewPassword: req.Newpassword,
	}
	err := h.se.UpdatePassword(ctx, &body)
	if err != nil {
		return &pb.Response{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, err
	}

	return &pb.Response{
		Status: http.StatusOK,
	}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.Response, error) {
	err := h.se.DeleteUser(ctx, req.Id)
	if err != nil {
		return &pb.Response{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, err
	}

	return &pb.Response{
		Status: http.StatusOK,
	}, nil
}
