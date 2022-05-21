package repository

import (
	"context"

	"github.com/cecepsprd/foodstore-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	Create(ctx context.Context, order model.Order) error
	StoreOrderItem(ctx context.Context, request []model.OrderItem) ([]model.OrderItem, error)
	Read(ctx context.Context, userid primitive.ObjectID) (response []model.Order, err error)
}

type mysqlOrderRepository struct {
	db *mongo.Database
}

func NewOrderRepository(db *mongo.Database) OrderRepository {
	return &mysqlOrderRepository{
		db: db,
	}
}

func (repo *mysqlOrderRepository) Read(ctx context.Context, userid primitive.ObjectID) (response []model.Order, err error) {
	cursor, err := repo.db.Collection("orders").
		Find(
			ctx,
			bson.M{"user_id": userid},
		)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &response); err != nil {
		return nil, err
	}

	return response, nil
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

func (repo *mysqlOrderRepository) StoreOrderItem(ctx context.Context, items []model.OrderItem) ([]model.OrderItem, error) {
	data := make([]interface{}, 0)

	for _, item := range items {
		data = append(data, item)
	}

	res, err := repo.db.Collection("order_item").
		InsertMany(
			ctx,
			data,
		)

	for i, id := range res.InsertedIDs {
		items[i].ID = id.(primitive.ObjectID)
	}

	if err != nil {
		return nil, err
	}

	return items, nil
}
