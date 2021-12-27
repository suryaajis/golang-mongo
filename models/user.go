package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     *string            `json:"name" validate:"required"`
	Email    *string            `json:"email" validate:"required,email"`
	Password string             `json:"password" validate:"required,min=6"`
	Phone    string             `bson:"phone" json:"phone"`
}
