package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tag struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name string             `bson:"name"          json:"name" validate:"required,min=3,max=20"`
}
