package service

import (
	"context"
	"time"

	"github.com/cecepsprd/foodstore-server/internal/model"
	"github.com/cecepsprd/foodstore-server/internal/repository"
	"github.com/cecepsprd/foodstore-server/utils/logger"
)

type TagService interface {
	Read(context.Context, model.Paging) ([]model.Tag, error)
	Create(ctx context.Context, tag model.Tag) error
	Update(ctx context.Context, tag model.Tag) (*model.Tag, error)
	Delete(ctx context.Context, id string) error
	ReadByID(ctx context.Context, id string) (*model.Tag, error)
}

type tag struct {
	repo           repository.TagRepository
	contextTimeout time.Duration
}

func NewTagService(repo repository.TagRepository, timeout time.Duration) TagService {
	return &tag{
		repo:           repo,
		contextTimeout: timeout,
	}
}

func (s *tag) Read(ctx context.Context, req model.Paging) ([]model.Tag, error) {
	categories, err := s.repo.Read(ctx, req)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	return categories, nil
}

func (s *tag) Create(ctx context.Context, tag model.Tag) error {
	err := s.repo.Create(ctx, tag)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	return nil
}

func (s *tag) Update(ctx context.Context, tag model.Tag) (*model.Tag, error) {
	response, err := s.repo.Update(ctx, tag)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	return response, nil
}

func (s *tag) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	return nil
}

func (s *tag) ReadByID(ctx context.Context, id string) (*model.Tag, error) {
	Tag, err := s.repo.ReadByID(ctx, id)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	return Tag, nil
}
