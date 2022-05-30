package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID              primitive.ObjectID `bson:"_id" json:"_id"`
	OrderNumber     int64              `bson:"order_number" json:"order_number"`
	Status          string             `bson:"status" json:"status"`
	DeliveryFee     int64              `bson:"delivery_fee" json:"delivery_fee"`
	DeliveryAddress DeliveryAddress    `bson:"delivery_address" json:"delivery_address"`
	OrderItems      []OrderItem        `bson:"order_item" json:"order_item"`
	UserID          primitive.ObjectID `bson:"user_id" json:"user_id"`
}

type OrderItem struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name      string             `bson:"name" json:"name"`
	Price     int64              `bson:"price" json:"price"`
	Qty       int64              `bson:"quantity" json:"qty"`
	ProductID primitive.ObjectID `bson:"product_id" json:"product_id"`
	OrderID   primitive.ObjectID `bson:"order_id" json:"order_id"`
}

type CreateOrderRequest struct {
	DeliveryFee       string `json:"delivery_fee"`
	DeliveryAddressID string `json:"delivery_address"`
	User              User   `json:"-" validate:"-"`
}
