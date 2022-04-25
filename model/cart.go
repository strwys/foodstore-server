package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CartItem struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"  json:"_id"`
	Name      string             `bson:"name"           json:"name"`
	Qty       int64              `bson:"quantity"       json:"qty"`
	Price     float64            `bson:"price"          json:"price"`
	ImageURL  string             `bson:"image_url"      json:"image_url"`
	UserID    primitive.ObjectID `bson:"user"           json:"user"`
	ProductID primitive.ObjectID `bson:"product"        json:"product"`
}

type UpdateCartItemRequest struct {
	Items []Items `json:"items"`
}

type Items struct {
	ID       string  `json:"_id"`
	Name     string  `json:"name"`
	Qty      int64   `json:"qty"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"image_url"`
}
