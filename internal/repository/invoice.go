package repository

import (
	"context"

	"github.com/cecepsprd/foodstore-server/internal/model"
	"github.com/cecepsprd/foodstore-server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InvoiceRepository interface {
	Create(ctx context.Context, order model.Invoice) error
	Read(ctx context.Context, orderid string) (model.Invoice, error)
}

type mysqlInvoiceRepository struct {
	db *mongo.Database
}

func NewInvoiceRepository(db *mongo.Database) InvoiceRepository {
	return &mysqlInvoiceRepository{
		db: db,
	}
}

func (repo *mysqlInvoiceRepository) Create(ctx context.Context, invoice model.Invoice) error {
	_, err := repo.db.Collection("invoice").
		InsertOne(
			ctx,
			invoice,
		)
	if err != nil {
		return err
	}

	return nil
}

func (repo *mysqlInvoiceRepository) Read(ctx context.Context, orderid string) (model.Invoice, error) {
	var invoice model.Invoice
	err := repo.db.Collection("invoice").
		FindOne(
			ctx,
			bson.M{"order_id": utils.ConvertPrimitiveID(orderid)},
		).Decode(&invoice)

	if err != nil {
		return model.Invoice{}, err
	}

	return invoice, nil
}
