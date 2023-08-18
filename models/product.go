package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductDescription string
	ProductPrice int16
	SellerId int16
}