package models

import "errors"

var (
  UserNotFound = errors.New("user not found")
  ProductNotFound = errors.New("product not found")
  CartNotFound = errors.New("cart not found")
  PromoNotFound = errors.New("promo not found")

  UserAlreadyExists = errors.New("user already exists")
  ProductAlreadyExists = errors.New("product already exists")
  CartAlreadyExists = errors.New("cart already exists")
  PromoAlreadyExists = errors.New("promot already exists")
)
