package service

import (
	"context"
	"time"

	"github.com/cecepsprd/foodstore-server/model"
	"github.com/cecepsprd/foodstore-server/repository"
	"github.com/cecepsprd/foodstore-server/utils/logger"
)

type CategoryService interface {
	Read(context.Context, model.Paging) ([]model.Category, error)
	Create(ctx context.Context, category model.Category) error
	Update(ctx context.Context, category model.Category) (*model.Category, error)
	Delete(ctx context.Context, id string) error
	ReadByID(ctx context.Context, id string) (*model.Category, error)
}

type Category struct {
	repo           repository.CategoryRepository
	contextTimeout time.Duration
}

func NewCategoryService(repo repository.CategoryRepository, timeout time.Duration) CategoryService {
	return &Category{
		repo:           repo,
		contextTimeout: timeout,
	}
}

func (s *Category) Read(ctx context.Context, req model.Paging) ([]model.Category, error) {
	categories, err := s.repo.Read(ctx, req)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	return categories, nil
}

func (s *Category) Create(ctx context.Context, category model.Category) error {
	err := s.repo.Create(ctx, category)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	return nil
}

func (s *Category) Update(ctx context.Context, Category model.Category) (*model.Category, error) {
	response, err := s.repo.Update(ctx, Category)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	return response, nil
}

func (s *Category) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	return nil
}

func (s *Category) ReadByID(ctx context.Context, id string) (*model.Category, error) {
	category, err := s.repo.ReadByID(ctx, id)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	return category, nil
}
