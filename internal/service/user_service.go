package service

import (
	"context"
	"fmt"
	"internet-store/internal/entity"
	"internet-store/internal/repository"
)

type UserService interface {
	UpdateUser(ctx context.Context, user *entity.UpdateUserRequest) error
}

type service struct {
	r repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &service{r: r}
}

func (s *service) UpdateUser(ctx context.Context, req *entity.UpdateUserRequest) error {
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
