package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name      string             `bson:"name"          json:"name"`
	CreatedAt time.Time          `bson:"created_at"    json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"    json:"updated_at"`
}

type CategoryRequest struct {
	Name string `bson:"name"       json:"name" validate:"required,min=3,max=20"`
}

func (c *Category) BeforeSave() Category {
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	return *c
}
