package service

import (
	"context"
	"reflect"
	"time"

	"github.com/cecepsprd/foodstore-server/constans"
	"github.com/cecepsprd/foodstore-server/model"
	"github.com/cecepsprd/foodstore-server/repository"
	"github.com/cecepsprd/foodstore-server/utils"
	"github.com/cecepsprd/foodstore-server/utils/logger"
)

type UserService interface {
	Read(ctx context.Context) (users []model.User, err error)
	ReadByUsername(ctx context.Context, username string) (user *model.User, err error)
	ReadByEmail(ctx context.Context, email string) (user *model.User, err error)
	Create(ctx context.Context, request model.User) error
}

type userService struct {
	repo           repository.UserRepository
	contextTimeout time.Duration
}

func NewUserService(urepo repository.UserRepository, timeout time.Duration) UserService {
	return &userService{
		repo:           urepo,
		contextTimeout: timeout,
	}
}

func (s *userService) ReadByUsername(ctx context.Context, name string) (*model.User, error) {
	user, err := s.repo.ReadByName(ctx, name)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	return user, nil
}

func (s *userService) ReadByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := s.repo.ReadByEmail(ctx, email)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	return user, nil
}

func (s *userService) Read(ctx context.Context) ([]model.User, error) {
	users, err := s.repo.Read(ctx)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	return users, nil
}

func (s *userService) Create(ctx context.Context, request model.User) error {
	emailCounted, err := s.repo.CountEmail(ctx, request.Email)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	if !reflect.ValueOf(emailCounted).IsZero() {
		return constans.ErrConflict
	}

	request.Password, err = utils.HashPassword(request.Password)
	if err != nil {
		return err
	}

	err = s.repo.Create(ctx, request)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	return nil
}
