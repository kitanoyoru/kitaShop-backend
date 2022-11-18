package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`

	RegisteredAt time.Time `json:"registeredAt" bson:"registeredAt"`
	LastVisitAt  time.Time `json:"lastVisitAt" bson:"lastVisitAt"`

	ShoppingCart Cart `json:"cart" bson:"cart"`
}

type UserInfo struct {
	ID       primitive.ObjectID `json:"id" bson"id"`
	Username string             `json:"username" bson:"username"`
}
