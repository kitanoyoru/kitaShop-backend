package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Product   ProductInfo        `json:"product" bson:"product"`
	Promo     PromoInfo          `json:"promo" bson:"promo"`
	Cart      Cart               `json:"cart" bson:"cart"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}

type OrderInfo struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Product ProductInfo        `json:"product" bson:"product"`
	Promo   PromoInfo          `json:"promo" bson:"promo"`
}
