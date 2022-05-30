package service

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/cecepsprd/foodstore-server/internal/model"
	"github.com/cecepsprd/foodstore-server/internal/repository"
	"github.com/cecepsprd/foodstore-server/utils"
	"github.com/cecepsprd/foodstore-server/utils/logger"
)

type ProductService interface {
	Read(context.Context, model.ReadProductRequest) (products []model.Product, total int64, err error)
	Create(ctx context.Context, product model.ProductRequest) error
	Update(ctx context.Context, product model.ProductRequest) (*model.Product, error)
	Delete(ctx context.Context, id string) error
	ReadByID(ctx context.Context, id string) (*model.Product, error)
}

type product struct {
	repo           repository.ProductRepository
	contextTimeout time.Duration
}

func NewProductService(repo repository.ProductRepository, timeout time.Duration) ProductService {
	return &product{
		repo:           repo,
		contextTimeout: timeout,
	}
}

func (s *product) Read(ctx context.Context, req model.ReadProductRequest) ([]model.Product, int64, error) {
	products, total, err := s.repo.Read(ctx, req)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, 0, err
	}

	return products, total, nil
}

func (s *product) Create(ctx context.Context, request model.ProductRequest) error {
	var (
		err     error
		product model.Product
	)

	if err := utils.MappingRequest(request, &product); err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	if !reflect.ValueOf(request.CategoryID).IsZero() {
		product.Category, err = s.repo.ReadCategoryByID(ctx, request.CategoryID)
		if err != nil {
			logger.Log.Warn(err.Error())
		}
	}

	if !reflect.ValueOf(request.Tags).IsZero() {
		product.Tags, err = s.repo.ReadTagsByName(ctx, request.Tags)
		if err != nil {
			logger.Log.Warn(err.Error())
		}
	}

	err = s.repo.Create(ctx, product)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	return nil
}

func (s *product) Update(ctx context.Context, request model.ProductRequest) (*model.Product, error) {
	var (
		product model.Product
	)

	err := utils.MappingRequest(request, &product)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	if !reflect.ValueOf(product).IsZero() {
		product.Category, err = s.repo.ReadCategoryByID(ctx, request.CategoryID)
		if err != nil {
			logger.Log.Error(err.Error())
			return nil, errors.New("category doesn't exist")
		}
	}

	if !reflect.ValueOf(request.Tags).IsZero() {
		product.Tags, err = s.repo.ReadTagsByName(ctx, request.Tags)
		if err != nil {
			logger.Log.Warn(err.Error())
		}
	}

	response, err := s.repo.Update(ctx, product)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	return response, nil
}

func (s *product) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	return nil
}

func (s *product) ReadByID(ctx context.Context, id string) (*model.Product, error) {
	product, err := s.repo.ReadByID(ctx, id)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	return product, nil
}
