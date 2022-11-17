package models

import (
  "time"

  "go.mongodb.org/mongo-driver/bson/primitive"
)

type Promo struct {
  ID primitive.ObjectID `json:"id" bson:"_id"`
  Code string `json:"code" bson:"code"`
  Percent uint16 `json:"percent" bson:"percent"`
  CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

type PromoInfo struct {
  ID primitive.ObjectID `json:"id" bson:"_id"`
  Code string `json:"code" bson:"code"`
  Percent uint16 `json:"percent" bson:"percent"`
}
