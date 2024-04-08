package repository

import (
	"context"

	"github.com/strwys/foodstore-server/internal/model"
	"github.com/strwys/foodstore-server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CategoryRepository interface {
	Read(context.Context, model.Paging) ([]model.Category, error)
	Create(ctx context.Context, category model.Category) error
	Update(ctx context.Context, category model.Category) (*model.Category, error)
	Delete(ctx context.Context, id string) error
	ReadByID(ctx context.Context, id string) (*model.Category, error)
}

type mysqlCategoryRepository struct {
	db *mongo.Database
}

func NewCategoryRepository(db *mongo.Database) CategoryRepository {
	return &mysqlCategoryRepository{
		db: db,
	}
}

func (repo *mysqlCategoryRepository) Read(ctx context.Context, req model.Paging) (response []model.Category, err error) {
	opt := new(options.FindOptions)
	cursor, err := repo.db.Collection("category").
		Find(
			ctx,
			bson.M{},
			opt.SetSkip(req.Offset),
			opt.SetLimit(req.Limit),
		)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (repo *mysqlCategoryRepository) Create(ctx context.Context, category model.Category) error {
	_, err := repo.db.Collection("category").
		InsertOne(
			ctx,
			category,
		)
	if err != nil {
		return err
	}

	return nil
}

func (repo *mysqlCategoryRepository) Update(ctx context.Context, category model.Category) (*model.Category, error) {
	err := repo.db.Collection("category").
		FindOneAndUpdate(
			ctx,
			bson.M{"_id": category.ID},
			bson.D{{Key: "$set", Value: category}},
		).
		Err()

	if err != nil {
		return nil, err
	}

	return &category, err
}

func (repo *mysqlCategoryRepository) Delete(ctx context.Context, categoryID string) error {
	_, err := repo.db.Collection("category").
		DeleteOne(
			ctx,
			bson.M{"_id": utils.ConvertPrimitiveID(categoryID)},
		)

	if err != nil {
		return err
	}

	return nil
}

func (repo *mysqlCategoryRepository) ReadByID(ctx context.Context, id string) (*model.Category, error) {
	var Category model.Category
	err := repo.db.Collection("category").
		FindOne(
			ctx,
			bson.M{"_id": utils.ConvertPrimitiveID(id)},
		).
		Decode(&Category)
	if err != nil {
		return nil, err
	}

	return &Category, nil
}
