package service

import (
	"context"
	"time"

	"github.com/cecepsprd/foodstore-server/internal/model"
	"github.com/cecepsprd/foodstore-server/internal/repository"
	"github.com/cecepsprd/foodstore-server/utils"
	"github.com/cecepsprd/foodstore-server/utils/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService interface {
	Create(ctx context.Context, req model.CreateOrderRequest) (*model.Order, error)
	Read(ctx context.Context, userid primitive.ObjectID) ([]model.Order, error)
}

type order struct {
	orderRepository           repository.OrderRepository
	cartRepository            repository.CartRepository
	deliveryAddressRepository repository.AddressRepository
	invoiceRepository         repository.InvoiceRepository
	contextTimeout            time.Duration
}

func NewOrderService(
	orderRepository repository.OrderRepository,
	cartRepository repository.CartRepository,
	deliveryAddressRepository repository.AddressRepository,
	invoiceRepository repository.InvoiceRepository,
	timeout time.Duration) OrderService {
	return &order{
		orderRepository:           orderRepository,
		cartRepository:            cartRepository,
		deliveryAddressRepository: deliveryAddressRepository,
		invoiceRepository:         invoiceRepository,
		contextTimeout:            timeout,
	}
}

func (s *order) Read(ctx context.Context, userid primitive.ObjectID) ([]model.Order, error) {
	orders, err := s.orderRepository.Read(ctx, userid)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	return orders, nil
}

func (s *order) Create(ctx context.Context, req model.CreateOrderRequest) (*model.Order, error) {
	items, err := s.cartRepository.Read(ctx, req.User.ID)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	deliveryAddress, err := s.deliveryAddressRepository.ReadByID(ctx, req.DeliveryAddressID)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	order := model.Order{
		ID:              primitive.NewObjectID(),
		Status:          "waiting_payment",
		DeliveryFee:     utils.ConvertStringToInt(req.DeliveryFee),
		DeliveryAddress: deliveryAddress,
		UserID:          req.User.ID,
	}

	var subtotal int64

	for _, item := range items {
		order.OrderItems = append(order.OrderItems, model.OrderItem{
			Name:      item.Name,
			Price:     int64(item.Price),
			Qty:       item.Qty,
			ProductID: item.ProductID,
			OrderID:   order.ID,
		})

		subtotal += int64(item.Price)
	}

	order.OrderItems, err = s.orderRepository.StoreOrderItem(ctx, order.OrderItems)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	if err = s.orderRepository.Create(ctx, order); err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	if err = s.cartRepository.Delete(ctx, req.User.ID); err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	invoice := model.Invoice{
		Subtotal:        subtotal,
		DeliveryFee:     order.DeliveryFee,
		DeliveryAddress: order.DeliveryAddress,
		Total:           subtotal + order.DeliveryFee,
		PaymentStatus:   "paid",
		User:            req.User,
		OrderID:         order.ID,
	}

	if err = s.invoiceRepository.Create(ctx, invoice); err != nil {
		logger.Log.Error(err.Error())
	}

	return &order, nil
}
