package models

import "gorm.io/gorm"

type Trans struct {
	gorm.Model
	BuyerId   int16
	SellerId  int16
	ProductId int16
}
