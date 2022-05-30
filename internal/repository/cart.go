package repository

import (
	"context"

	"github.com/cecepsprd/foodstore-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CartRepository interface {
	Read(ctx context.Context, userid primitive.ObjectID) ([]model.CartItem, error)
	Update(ctx context.Context, userid primitive.ObjectID, cart model.CartItem) error
	Delete(ctx context.Context, userid primitive.ObjectID) error
	DeleteByID(ctx context.Context, userid primitive.ObjectID, itemid primitive.ObjectID) error
}

type mysqlCartRepository struct {
	db *mongo.Database
}

func NewCartRepository(db *mongo.Database) CartRepository {
	return &mysqlCartRepository{
		db: db,
	}
}

func (repo *mysqlCartRepository) Read(ctx context.Context, userid primitive.ObjectID) (response []model.CartItem, err error) {
	cursor, err := repo.db.Collection("cart_item").
		Find(
			ctx,
			bson.M{"user": userid},
		)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (repo *mysqlCartRepository) Update(ctx context.Context, userid primitive.ObjectID, request model.CartItem) error {
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"user": userid, "product": request.ID}
	update := bson.M{"$set": request}

	_, err := repo.db.Collection("cart_item").UpdateOne(
		ctx, filter, update, opts,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo *mysqlCartRepository) Delete(ctx context.Context, userid primitive.ObjectID) error {
	_, err := repo.db.Collection("cart_item").
		DeleteMany(
			ctx,
			bson.M{"user": userid},
		)
	if err != nil {
		return err
	}

	return nil
}

func (repo *mysqlCartRepository) DeleteByID(ctx context.Context, userid primitive.ObjectID, itemid primitive.ObjectID) error {
	_, err := repo.db.Collection("cart_item").
		DeleteMany(
			ctx,
			bson.M{"user": userid, "product": itemid},
		)
	if err != nil {
		return err
	}

	return nil
}
