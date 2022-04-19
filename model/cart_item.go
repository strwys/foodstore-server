package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CartItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"  json:"_id"`
	Name     string             `bson:"name"           json:"name"`
	Quantity string             `bson:"quantity"       json:"quantity"`
	Price    float64            `bson:"price"          json:"price"`
	ImageURL string             `bson:"image_url"      json:"image_url"`
	User     User               `bson:"user"           json:"user"`
	Product  Product            `bson:"product"        json:"product"`
}

type UpdateCartItemRequest struct {
	Name     string  `bson:"name"           json:"name"`
	Quantity string  `bson:"quantity"       json:"quantity"`
	Price    float64 `bson:"price"          json:"price"`
	ImageURL string  `bson:"image_url"      json:"image_url"`
	User     User    `bson:"user"           json:"user"`
	Product  Product `bson:"product"        json:"product"`
}
