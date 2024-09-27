package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}

type RegisterRequest struct {
	Name     string `json:"name" example:"test user"`
	Email    string `json:"email" example:"test@gmail.com"`
	Password string `json:"password" example:"test"`
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
