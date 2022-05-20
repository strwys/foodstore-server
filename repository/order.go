package repository

import (
	"context"

	"github.com/cecepsprd/foodstore-server/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	Create(ctx context.Context, order model.Order) error
	StoreOrderItem(ctx context.Context, request []model.OrderItem) error
}

type mysqlOrderRepository struct {
	db *mongo.Database
}

func NewOrderRepository(db *mongo.Database) OrderRepository {
	return &mysqlOrderRepository{
		db: db,
	}
}

func (repo *mysqlOrderRepository) Create(ctx context.Context, order model.Order) error {
	_, err := repo.db.Collection("orders").
		InsertOne(
			ctx,
			order,
		)
	if err != nil {
		return err
	}

	return nil
}

func (repo *mysqlOrderRepository) StoreOrderItem(ctx context.Context, request []model.OrderItem) error {
	orderItems := make([]interface{}, 0)

	for _, v := range request {
		orderItems = append(orderItems, v)
	}

	_, err := repo.db.Collection("order_item").
		InsertMany(
			ctx,
			orderItems,
		)
	if err != nil {
		return err
	}

	return nil
}
