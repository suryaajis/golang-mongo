package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string            `json:"name" validate:"required"`
	Email      *string            `json:"email" validate:"required,email"`
	Password   *string            `json:"password" validate:"required,min=6"`
	Phone      *string            `json:"phone" validate:"required"`
	Token      *string            `json:"token"`
	Role       *string            `json:"role" validate:"required,eq=Admin|eq=User"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	User_id    string             `json:"user_id"`
}
