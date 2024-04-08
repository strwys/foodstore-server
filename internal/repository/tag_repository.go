package repository

import (
	"context"

	"github.com/strwys/foodstore-server/internal/model"
	"github.com/strwys/foodstore-server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TagRepository interface {
	Read(context.Context, model.Paging) ([]model.Tag, error)
	Create(ctx context.Context, tag model.Tag) error
	Update(ctx context.Context, tag model.Tag) (*model.Tag, error)
	ReadByID(ctx context.Context, id string) (*model.Tag, error)
	Delete(ctx context.Context, id string) error
}

type mysqlTagRepository struct {
	db *mongo.Database
}

func NewTagRepository(db *mongo.Database) TagRepository {
	return &mysqlTagRepository{
		db: db,
	}
}

func (repo *mysqlTagRepository) Read(ctx context.Context, req model.Paging) (response []model.Tag, err error) {
	opt := new(options.FindOptions)
	cursor, err := repo.db.Collection("Tag").
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

func (repo *mysqlTagRepository) Create(ctx context.Context, tag model.Tag) error {
	_, err := repo.db.Collection("tag").
		InsertOne(
			ctx,
			tag,
		)
	if err != nil {
		return err
	}

	return nil
}

func (repo *mysqlTagRepository) Update(ctx context.Context, tag model.Tag) (*model.Tag, error) {
	err := repo.db.Collection("Tag").
		FindOneAndUpdate(
			ctx,
			bson.M{"_id": tag.ID},
			bson.D{{Key: "$set", Value: tag}},
		).
		Err()

	if err != nil {
		return nil, err
	}

	return &tag, err
}

func (repo *mysqlTagRepository) Delete(ctx context.Context, tagID string) error {
	_, err := repo.db.Collection("tag").
		DeleteOne(
			ctx,
			bson.M{"_id": utils.ConvertPrimitiveID(tagID)},
		)

	if err != nil {
		return err
	}

	return nil
}

func (repo *mysqlTagRepository) ReadByID(ctx context.Context, id string) (*model.Tag, error) {
	var tag model.Tag
	err := repo.db.Collection("tag").
		FindOne(
			ctx,
			bson.M{"_id": utils.ConvertPrimitiveID(id)},
		).
		Decode(&tag)
	if err != nil {
		return nil, err
	}

	return &tag, nil
}
