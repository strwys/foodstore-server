package service

import (
	"context"
	"time"

	"github.com/strwys/foodstore-server/internal/model"
	"github.com/strwys/foodstore-server/internal/repository"
	"github.com/strwys/foodstore-server/utils/logger"
)

type AddressService interface {
	Create(ctx context.Context, req model.DeliveryAddress) error
	Read(context.Context) (response []model.DeliveryAddress, err error)
	Update(ctx context.Context, req model.DeliveryAddress) (*model.DeliveryAddress, error)
	Delete(ctx context.Context, id string) error
}

type address struct {
	repo           repository.AddressRepository
	contextTimeout time.Duration
}

func NewAddressService(repo repository.AddressRepository, timeout time.Duration) AddressService {
	return &address{
		repo:           repo,
		contextTimeout: timeout,
	}
}

func (s *address) Read(ctx context.Context) ([]model.DeliveryAddress, error) {
	deliveryAddresss, err := s.repo.Read(ctx)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	return deliveryAddresss, nil
}

func (s *address) Create(ctx context.Context, request model.DeliveryAddress) error {
	err := s.repo.Create(ctx, request)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	return nil
}

func (s *address) Update(ctx context.Context, request model.DeliveryAddress) (*model.DeliveryAddress, error) {
	response, err := s.repo.Update(ctx, request)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	return response, nil
}

func (s *address) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	return nil
}
