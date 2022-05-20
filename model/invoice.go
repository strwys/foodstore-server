package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Invoice struct {
	Subtotal        int64              `bson:"subtotal" json:"subtotal"`
	DeliveryFee     int64              `bson:"delivery_fee" json:"delivery_fee"`
	DeliveryAddress DeliveryAddress    `bson:"delivery_address" json:"delivery_address"`
	Total           int64              `bson:"total" json:"total"`
	PaymentStatus   string             `bson:"payment_status" json:"payment_status"`
	User            User               `bson:"user" json:"user"`
	OrderID         primitive.ObjectID `bson:"order_id" json:"order_id"`
}
