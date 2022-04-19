package repository

import (
	"context"

	"github.com/cecepsprd/foodstore-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(ctx context.Context, user model.User) error
	Read(context.Context) ([]model.User, error)
	CountEmail(ctx context.Context, email string) (int64, error)
	ReadByName(ctx context.Context, name string) (user *model.User, err error)
	ReadByEmail(ctx context.Context, email string) (user *model.User, err error)
}

type mysqlUserRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &mysqlUserRepository{
		db: db,
	}
}

func (repo *mysqlUserRepository) Read(ctx context.Context) (response []model.User, err error) {
	return response, nil
}

func (repo *mysqlUserRepository) ReadByName(ctx context.Context, name string) (*model.User, error) {
	var user model.User
	err := repo.db.
		Collection("user").
		FindOne(
			ctx,
			bson.M{"full_name": name},
		).
		Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *mysqlUserRepository) ReadByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := repo.db.
		Collection("user").
		FindOne(
			ctx,
			bson.M{"email": email},
		).
		Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *mysqlUserRepository) Create(ctx context.Context, user model.User) error {
	if err := user.BeforeSave(); err != nil {
		return err
	}

	_, err := repo.db.Collection("user").
		InsertOne(
			ctx,
			bson.M{
				"full_name":   user.FullName,
				"customer_id": bson.M{"$inc": bson.M{"customer_id": 1}},
				"email":       user.Email,
				"password":    user.Password,
				"role":        user.Role,
				"token":       user.Token,
				"created_at":  user.CreatedAt,
				"updated_at":  user.UpdatedAt,
			},
		)
	if err != nil {
		return err
	}

	return nil
}

func (repo *mysqlUserRepository) CountEmail(ctx context.Context, email string) (int64, error) {
	total, err := repo.db.Collection("user").
		CountDocuments(
			ctx,
			bson.M{"email": email},
		)
	if err != nil {
		return 0, err
	}

	return total, nil
}
