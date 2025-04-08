package service

import (
	"context"
	"fmt"
	"internet-store/internal/entity"
	"internet-store/internal/repository"
)

type UserService interface {
	UpdateUser(ctx context.Context, user *entity.UpdateUserRequest) error
	GetUserByID(ctx context.Context, id int64) (*entity.User, error)
}

type userService struct {
	r repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r: r}
}

func (s *userService) UpdateUser(ctx context.Context, req *entity.UpdateUserRequest) error {
	if req.ID == 0 {
		return entity.ErrorInvalidParams
	}

	if req.Name == "" {
		return entity.ErrorInvalidParams
	}

	if req.Phone == "" {
		return entity.ErrorInvalidParams
	}

	updateReq := &repository.UpdateUserRequest{
		ID:    req.ID,
		Name:  &req.Name,
		Phone: &req.Phone,
	}

	if req.Email != "" {
		updateReq.Email = &req.Email
	}

	err := s.r.UpdateUser(ctx, updateReq)
	if err != nil {
		return fmt.Errorf("%w: %w", entity.ErrorInternal, err)
	}

	return nil
}

func (s *userService) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	user, err := s.r.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", entity.ErrorUserNotFound, err)
	}
	return user, nil
}
