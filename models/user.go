package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt  time.Time          `bson:"created_at,omitempty"`
	FullName   string             `json:"full_name" bson:"full_name" validate:"required"`
	Username   string             `json:"username" bson:"username" validate:"required,unique"`
	Email      string             `json:"email" bson:"email" validate:"required,email,unique"`
	Password   string             `bson:"password" validate:"required"`
	IsDisabled bool               `bson:"is_disabled"`
}
