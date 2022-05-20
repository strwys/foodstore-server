package repository

import (
	"context"

	"github.com/cecepsprd/foodstore-server/model"
	"github.com/cecepsprd/foodstore-server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AddressRepository interface {
	Create(ctx context.Context, req model.DeliveryAddress) error
	Read(context.Context) ([]model.DeliveryAddress, error)
	Update(ctx context.Context, req model.DeliveryAddress) (*model.DeliveryAddress, error)
	Delete(ctx context.Context, id string) error
	ReadByID(ctx context.Context, id string) (response model.DeliveryAddress, err error)
}

type mysqlAddressRepository struct {
	db *mongo.Database
}

func NewAddressRepository(db *mongo.Database) AddressRepository {
	return &mysqlAddressRepository{
		db: db,
	}
}

func (repo *mysqlAddressRepository) Read(ctx context.Context) (response []model.DeliveryAddress, err error) {
	cursor, err := repo.db.Collection("delivery_address").
		Find(
			ctx,
			bson.M{},
		)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (repo *mysqlAddressRepository) Create(ctx context.Context, request model.DeliveryAddress) error {
	_, err := repo.db.Collection("delivery_address").
		InsertOne(
			ctx,
			request,
		)
	if err != nil {
		return err
	}

	return nil
}

func (repo *mysqlAddressRepository) Update(ctx context.Context, data model.DeliveryAddress) (*model.DeliveryAddress, error) {
	err := repo.db.Collection("products").
		FindOneAndUpdate(
			ctx,
			bson.M{"_id": data.ID},
			bson.D{{Key: "$set", Value: data.BeforeSave()}},
		).
		Err()

	if err != nil {
		return nil, err
	}

	return &data, err
}

func (repo *mysqlAddressRepository) Delete(ctx context.Context, productID string) error {
	_, err := repo.db.Collection("products").
		DeleteOne(
			ctx,
			bson.M{"_id": utils.ConvertPrimitiveID(productID)},
		)

	if err != nil {
		return err
	}

	return nil
}

func (repo *mysqlAddressRepository) ReadByID(ctx context.Context, requestID string) (model.DeliveryAddress, error) {
	var deliveryAddress model.DeliveryAddress
	err := repo.db.Collection("delivery_address").
		FindOne(
			ctx,
			bson.M{"_id": utils.ConvertPrimitiveID(requestID)},
		).
		Decode(&deliveryAddress)

	if err != nil {
		return model.DeliveryAddress{}, err
	}

	return deliveryAddress, nil
}
