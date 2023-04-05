package db

import (
	"context"

	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/models"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
)

func (p *DBPostgres) CreateUser(ctx context.Context, user *models.User) error {
	_, err := p.Pool.Exec(ctx, "insert into users(id,password,email) values($1,$2,$3)",
		&user.ID, &user.Password, &user.Email)
	if err != nil {
		logrus.Errorf("auth service error: error while creating new user, %w", err)
		return err
	}
	return nil
}

func (p *DBPostgres) GetForUpdatePassword(ctx context.Context, id string) (*models.User, error) {
	user := models.User{}
	err := p.Pool.QueryRow(ctx, "select password,email from users where id=$1", id).Scan(
		&user.Password, &user.Email)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, models.ErrorUserDoesntExist
		}
		logrus.Errorf("auth service error: error get user: %w", err.Error())
		return nil, err
	}
	return &user, nil
}

func (p *DBPostgres) Get(ctx context.Context, email string) (*models.User, error) {
	user := models.User{}
	err := p.Pool.QueryRow(ctx, "select password,id from users where email=$1", email).Scan(
		&user.Password, &user.ID)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, models.ErrorUserDoesntExist
		}
		logrus.Errorf("auth service error: error get user: %w", err.Error())
		return nil, err
	}
	return &user, nil
}

func (p *DBPostgres) Update(ctx context.Context, body *models.UserUpdate) error {
	a, err := p.Pool.Exec(ctx, "update users set password=$1 where id=$2",
		&body.NewPassword, &body.Id)
	if a.RowsAffected() == 0 {
		return models.ErrorUserDoesntExist
	}
	if err != nil {
		logrus.Errorf("error with update user %w", err)
		return err
	}
	return nil
}

func (p *DBPostgres) Delete(ctx context.Context, id string) error {
	a, err := p.Pool.Exec(ctx, "delete from users where id=$1", id)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return models.ErrorUserDoesntExist
		}
		logrus.Errorf("error with delete user %w", err)
		return err
	}
	if a.RowsAffected() == 0 {
		return models.ErrorUserDoesntExist
	}
	return nil
}
