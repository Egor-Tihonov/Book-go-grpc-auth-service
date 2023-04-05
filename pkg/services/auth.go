package services

import (
	"context"

	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/models"
	"github.com/sirupsen/logrus"

	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/utils"
	"github.com/google/uuid"
)

func (s *Server) GetUser(ctx context.Context, id string) (*models.User, error) {
	return s.rps.Get(ctx, id)
}

func (s *Server) Registration(ctx context.Context, user *models.User) (string, error) {
	if !s.Validate(user) {
		return "", models.ErrorEmptyUser
	}

	newID := uuid.New().String()
	user.ID = newID
	hashpassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	user.Password = hashpassword

	if err = s.rps.CreateUser(ctx, user); err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return newID, nil
}

func (s *Server) Generate(ctx context.Context, password string, user *models.User) (string, error) {
	err := utils.CheckPassword(password, user.Password)
	if err != nil {
		return "", models.ErrorIncorrectPassword
	}

	token, err := s.Jwt.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Server) UpdatePassword(ctx context.Context, body *models.UserUpdate) error {
	user, err := s.rps.GetForUpdatePassword(ctx, body.Id)
	if err != nil {
		return err
	}

	err = utils.CheckPassword(body.OldPassword, user.Password)
	if err != nil {
		return models.ErrorIncorrectPassword
	}

	newpassw, err := utils.HashPassword(body.NewPassword)
	if err != nil {
		logrus.Errorf("auth service error: cannot hash password, %w", err)
	}

	body.NewPassword = newpassw
	err = s.rps.Update(ctx, body)
	if err != nil {
		return err
	}

	return err
}

func (s *Server) DeleteUser(ctx context.Context, id string) error {
	return s.rps.Delete(ctx, id)
}

func (s *Server) Validate(user *models.User) bool {
	if user.Email == "" {
		return false
	}

	if user.Password == "" {
		return false
	}
	return true
}
