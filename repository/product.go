package repository

import (
	"context"
	"reflect"

	"github.com/cecepsprd/foodstore-server/model"
	"github.com/cecepsprd/foodstore-server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"
)

type ProductRepository interface {
	Read(context.Context, model.ReadProductRequest) (products []model.Product, total int64, err error)
	Create(ctx context.Context, product model.Product) error
	Update(ctx context.Context, product model.Product) (*model.Product, error)
	ReadByID(ctx context.Context, id string) (*model.Product, error)
	Delete(ctx context.Context, id string) error
	ReadCategoryByID(ctx context.Context, id string) (model.Category, error)
	ReadTagsByName(ctx context.Context, tagName []string) (tags []model.Tag, err error)
}

type mysqlProductRepository struct {
	db *mongo.Database
}

func NewProductRepository(db *mongo.Database) ProductRepository {
	return &mysqlProductRepository{
		db: db,
	}
}

func (repo *mysqlProductRepository) Read(ctx context.Context, req model.ReadProductRequest) (response []model.Product, total int64, err error) {
	g, ctx := errgroup.WithContext(ctx)

	opt := new(options.FindOptions)
	filter := bson.M{}

	g.Go(func() error {
		if !reflect.ValueOf(req.Keyword).IsZero() {
			filter["name"] = bson.M{"$regex": primitive.Regex{Pattern: req.Keyword, Options: "i"}}
		}

		if !reflect.ValueOf(req.Category).IsZero() {
			filter["category"] = utils.ConvertPrimitiveID(req.Category)
		}

		if len(req.Tags) > 0 {
			filter["tags.name"] = bson.M{"$in": req.Tags}
		}

		cursor, err := repo.db.Collection("products").
			Find(
				ctx,
				filter,
				opt.SetSkip(req.Offset),
				opt.SetLimit(req.Limit),
			)
		if err != nil {
			return err
		}

		if err = cursor.All(ctx, &response); err != nil {
			return err
		}

		return nil
	})

	g.Go(func() error {
		total, err = repo.db.Collection("products").CountDocuments(ctx, bson.M{})
		if err != nil {
			return err
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, 0, err
	}

	return response, total, nil
}

func (repo *mysqlProductRepository) Create(ctx context.Context, product model.Product) error {
	_, err := repo.db.Collection("products").
		InsertOne(
			ctx,
			product.BeforeSave(),
		)
	if err != nil {
		return err
	}

	return nil
}

func (repo *mysqlProductRepository) Update(ctx context.Context, product model.Product) (*model.Product, error) {
	err := repo.db.Collection("products").
		FindOneAndUpdate(
			ctx,
			bson.M{"_id": product.ID},
			bson.D{{Key: "$set", Value: product.BeforeSave()}},
		).
		Err()

	if err != nil {
		return nil, err
	}

	return &product, err
}

func (repo *mysqlProductRepository) Delete(ctx context.Context, productID string) error {
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

func (repo *mysqlProductRepository) ReadByID(ctx context.Context, id string) (*model.Product, error) {
	var product model.Product
	err := repo.db.Collection("products").
		FindOne(
			ctx,
			bson.M{"_id": utils.ConvertPrimitiveID(id)},
		).
		Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *mysqlProductRepository) ReadCategoryByID(ctx context.Context, id string) (model.Category, error) {
	var category model.Category
	err := repo.db.Collection("category").FindOne(
		ctx,
		bson.M{"_id": utils.ConvertPrimitiveID(id)},
	).Decode(&category)

	if err != nil {
		return category, err
	}

	return category, nil
}

func (repo *mysqlProductRepository) ReadTagsByName(ctx context.Context, tagName []string) (tags []model.Tag, err error) {
	cursor, err := repo.db.Collection("tag").Find(
		ctx,
		bson.M{"name": bson.M{"$in": tagName}},
	)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &tags); err != nil {
		return nil, err
	}

	return tags, err
}
