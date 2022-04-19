package service

import (
	"time"

	"github.com/cecepsprd/foodstore-server/repository"
)

type CartService interface {
	// Read(context.Context, model.Paging) ([]model.Tag, error)
	// Create(ctx context.Context, tag model.Tag) error
	// Update(ctx context.Context, req []model.CartItem) (*model.CartItem, error)
	// Delete(ctx context.Context, id string) error
	// ReadByID(ctx context.Context, id string) (*model.Tag, error)
}

type cart struct {
	productRepository repository.ProductRepository
	contextTimeout    time.Duration
}

func NewCartService(productRepository repository.ProductRepository, timeout time.Duration) CartService {
	return &cart{
		productRepository: productRepository,
		contextTimeout:    timeout,
	}
}

// func (s *cart) Update(ctx context.Context, request []model.CartItem) (*model.CartItem, error) {
// 	var (
// 		itemIDs []string
// 	)

// 	for _, cartItem := range request {
// 		itemIDs = append(itemIDs, cartItem.Product.ID.String())
// 	}

// 	products, _, err := s.productRepository.Read(ctx, model.ReadProductRequest{ItemIDs: itemIDs})
// 	if err != nil {
// 		logger.Log.Error(err.Error())
// 		return nil, err
// 	}

// 	return nil, nil
// }

// func (s *tag) Read(ctx context.Context, req model.Paging) ([]model.Tag, error) {
// 	categories, err := s.repo.Read(ctx, req)
// 	if err != nil {
// 		logger.Log.Error(err.Error())
// 		return nil, err
// 	}

// 	return categories, nil
// }
