package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"  json:"_id"`
	Name        string             `bson:"name"           json:"name"        validate:"required,min=3,max=45"`
	Description string             `bson:"description"    json:"description" validate:"required"`
	Price       float32            `bson:"price"          json:"price"`
	ImageURL    string             `bson:"image_url"      json:"image_url"`
	CreatedAt   time.Time          `bson:"created_at"     json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"     json:"updated_at"`
	Category    Category           `bson:"-"              json:"category"`
	Tags        []Tag              `bson:"-"              json:"tags"`
}

type ReadProductRequest struct {
	Limit    int64    `json:"limit"`
	Offset   int64    `json:"skip"`
	Keyword  string   `json:"q"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
	ItemIDs  []string `json:"-"`
}

type ProductRequest struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name        string             `bson:"name"          json:"name"        validate:"required,min=3,max=45"`
	Description string             `bson:"description"   json:"description" validate:"required"`
	Price       float32            `bson:"price"         json:"price"`
	ImageURL    string             `bson:"image_url"     json:"image_url"`
	CategoryID  string             `bson:"category_id"   json:"category_id"`
	Tags        []string           `bson:"tags"          json:"-"`
}

func (p *Product) BeforeSave() Product {
	p.ImageURL = "images/default.jpg"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return *p
}
