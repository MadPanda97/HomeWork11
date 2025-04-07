package service_test

import (
	"context"
	"errors"
	"internet-store/internal/entity"
	"internet-store/internal/mock"
	"internet-store/internal/service"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUser(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		r := mock.NewMockUserRepository(ctrl)

		r.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(nil)

		s := service.NewUserService(r)

		err := s.UpdateUser(context.Background(), &entity.UpdateUserRequest{
			ID:    1,
			Name:  "salam",
			Email: "salam@salam.salam",
			Phone: "123123123",
		})

		assert.Nil(t, err)
	})

	t.Run("invalid id", func(t *testing.T) {
		s := service.NewUserService(nil)

		err := s.UpdateUser(context.Background(), &entity.UpdateUserRequest{
			ID:    0,
			Name:  "salam",
			Email: "salam@salam.salam",
			Phone: "123123123",
		})

		assert.ErrorIs(t, err, entity.ErrorInvalidParams)
	})

	t.Run("invalid name", func(t *testing.T) {
		s := service.NewUserService(nil)

		err := s.UpdateUser(context.Background(), &entity.UpdateUserRequest{
			ID:    1,
			Name:  "",
			Email: "salam@salam.salam",
			Phone: "123123123",
		})

		assert.ErrorIs(t, err, entity.ErrorInvalidParams)
	})

	t.Run("db error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		r := mock.NewMockUserRepository(ctrl)

		r.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(errors.New("db error"))

		s := service.NewUserService(r)

		err := s.UpdateUser(context.Background(), &entity.UpdateUserRequest{
			ID:    1,
			Name:  "salam",
			Email: "salam@salam.salam",
			Phone: "123123123",
		})

		assert.ErrorIs(t, err, entity.ErrorInternal)
	})
}
