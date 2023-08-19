package models

import "gorm.io/gorm"

type Reward struct {
	gorm.Model
	BuyerId  int16
	SellerId int16
	Coins    int16
	Count    int16
}
