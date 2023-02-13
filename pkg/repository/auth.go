package repository

import (
	"context"
	"errors"

	"github.com/Egor-Tihonov/go-grpc-auth-service.git/pkg/models"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
)

func (p *DBPostgres) CreateUser(ctx context.Context, user *models.User) error {
	_, err := p.Pool.Exec(ctx, "insert into users(id,password,email) values($1,$2,$3)",
		&user.ID, &user.Password, &user.Email)
	if err != nil {
		logrus.Errorf("error adding new user, %w", err)
		return err
	}
	return nil
}

func (p *DBPostgres) GetUser(ctx context.Context, email string) (*models.User, error) {
	user := models.User{}

	err := p.Pool.QueryRow(ctx, "select password,id from users where email=$1", email).Scan(
		&user.Password, &user.ID)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			userDoesntExist := errors.New("Пользователь не существует")
			return nil, userDoesntExist
		}
		return nil, err
	}

	return &user, nil
}
