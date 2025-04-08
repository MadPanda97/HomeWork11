package repository

import (
	"context"
	"database/sql"
	"fmt"
	"internet-store/internal/entity"
	"strings"
)

const (
	nameField     = "name"
	emailField    = "email"
	phoneField    = "phone"
	passwordField = "password"
	balanceField  = "balance"
)

//go:generate mockgen -source=user_repository.go -destination=../mock/user_repository.go -package=mock UserRepository

type UserRepository interface {
	UpdateUser(ctx context.Context, user *UpdateUserRequest) error
	GetUserByID(ctx context.Context, id int64) (*entity.User, error)
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

type userRepo struct {
	db *sql.DB
}

func (r *userRepo) UpdateUser(ctx context.Context, req *UpdateUserRequest) error {
	var (
		queryParams []string
		values      []any
	)

	if req.Phone != nil {
		queryParams = append(queryParams, fmt.Sprintf("%s=$%d", phoneField, len(queryParams)+1))
		values = append(values, *req.Phone)
	}

	if req.Email != nil {
		queryParams = append(queryParams, fmt.Sprintf("%s=$%d", emailField, len(queryParams)+1))
		values = append(values, *req.Email)
	}

	if req.Name != nil {
		queryParams = append(queryParams, fmt.Sprintf("%s=$%d", nameField, len(queryParams)+1))
		values = append(values, *req.Name)
	}

	if req.Password != nil {
		queryParams = append(queryParams, fmt.Sprintf("%s=$%d", passwordField, len(queryParams)+1))
		values = append(values, *req.Password)
	}

	if req.Balance != nil {
		queryParams = append(queryParams, fmt.Sprintf("%s=$%d", balanceField, len(queryParams)+1))
		values = append(values, *req.Balance)
	}

	if len(queryParams) == 0 {
		return nil
	}

	query := fmt.Sprintf("UPDATE users SET %s WHERE id=%d",
		strings.Join(queryParams, ", "),
		req.ID,
	)

	_, err := r.db.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	var user entity.User
	err := r.db.QueryRowContext(ctx,
		"SELECT id, name, email, phone FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Phone)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
