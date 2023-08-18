package models

import "gorm.io/gorm"

type AdminApproval struct {
	gorm.Model
	SellerId int16
}