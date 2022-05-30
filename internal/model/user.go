package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	FullName   string             `bson:"full_name"     json:"full_name"     validate:"required,min=3,max=45"`
	CustomerID int64              `bson:"-"             json:"customer_id"`
	Email      string             `bson:"email"         json:"email"         validate:"required,email"`
	Password   string             `bson:"password"      json:"password"      validate:"required"`
	Role       string             `bson:"role"          json:"role"          validate:"eq=user|eq=admin"`
	Token      string             `bson:"token"         json:"token"`
	CreatedAt  time.Time          `bson:"created_at"    json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"    json:"updated_at"`
}

func (u *User) BeforeSave() error {
	u.Role = "regular"
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}
