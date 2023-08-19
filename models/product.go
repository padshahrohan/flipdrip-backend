package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName        string
	ProductDescription string
	ProductPrice       int16
	SellerId           int16
	Tokens             int16
}
