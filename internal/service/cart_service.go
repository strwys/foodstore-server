package service

import (
	"context"
	"time"

	"github.com/strwys/foodstore-server/internal/model"
	"github.com/strwys/foodstore-server/internal/repository"
	"github.com/strwys/foodstore-server/utils/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartService interface {
	Read(ctx context.Context, userid primitive.ObjectID) ([]model.CartItem, error)
	Update(ctx context.Context, userid primitive.ObjectID, req []model.CartItem) error
	DeleteByItemID(ctx context.Context, userid primitive.ObjectID, itemid primitive.ObjectID) error
}

type cart struct {
	cartRepository repository.CartRepository
	contextTimeout time.Duration
}

func NewCartService(cartRepository repository.CartRepository, timeout time.Duration) CartService {
	return &cart{
		cartRepository: cartRepository,
		contextTimeout: timeout,
	}
}

func (s *cart) Read(ctx context.Context, userid primitive.ObjectID) ([]model.CartItem, error) {
	carts, err := s.cartRepository.Read(ctx, userid)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	return carts, nil
}

func (s *cart) Update(ctx context.Context, userid primitive.ObjectID, request []model.CartItem) error {
	for _, cart := range request {
		cart.UserID = userid
		cart.ProductID = cart.ID
		if err := s.cartRepository.Update(ctx, userid, cart); err != nil {
			logger.Log.Error(err.Error())
			return err
		}
	}

	return nil
}

func (s *cart) DeleteByItemID(ctx context.Context, userid primitive.ObjectID, itemid primitive.ObjectID) error {
	if err := s.cartRepository.DeleteByID(ctx, userid, itemid); err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	return nil
}
