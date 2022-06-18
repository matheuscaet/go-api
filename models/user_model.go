package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	User     string             `json:"user,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
}
