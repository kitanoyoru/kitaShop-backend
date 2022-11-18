package repository

import (
	"context"
	"errors"

	models "github.com/kitanoyoru/kitaShop-backend/internal/models/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PromocodeRepo struct {
	coll *mongo.Collection
}

func NewPromocodeRepo(db *mongo.Database) *PromocodeRepo {
	return &PromocodeRepo{
		coll: db.Collection(promocodeCollectionName),
	}
}

func (r *PromocodeRepo) Create(ctx context.Context, promocode models.Promo) (primitive.ObjectID, error) {
	res, err := r.coll.InsertOne(ctx, promocode)
	if err != nil {
		return primitive.NewObjectID(), err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (r *PromocodeRepo) GetByCode(ctx context.Context, code string) (models.Promo, error) {
	var promo models.Promo
	if err := r.coll.FindOne(ctx, bson.D{{"code", code}}).Decode(&promo); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Promo{}, models.PromoNotFound
		}

		return models.Promo{}, err
	}

	return promo, nil
}
