package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	UserName string
	Role string
	UserPassword string  
	WalletAddress string 
}