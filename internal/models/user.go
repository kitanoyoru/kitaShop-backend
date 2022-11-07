package models

import (
  "time"

  "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
  ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
  Username string `json:"username" bson:"username"`
  Email string `json:"email" bson:"email"`
  Password string `json:"password" bson:"password"`
  RegisteredAt time.Time `json:"registeredAt" bson:"registeredAt"`
  LastVisitAt time.Time `json:"lastVisitAt" bson:"lastVisitAt"`
  ShoppingCart primitive.ObjectID `json:"cart" bson:"cart"`
}
