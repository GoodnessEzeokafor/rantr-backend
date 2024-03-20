package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	// Email    string             `bson:"email" binding:"required"`
	Username string `bson:"username" binding:"required,min=3,max=20"` // Set min and max length
	Password string `bson:"password" binding:"required"`
	Bio      string `bson:"bio"` // Changed to optional
	Avatar   string `bson:"avatar"`
}
