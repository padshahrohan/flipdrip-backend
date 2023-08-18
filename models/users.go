package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	UserName string
	UserType string
	Userusername string
	UserPassword string  
	WalletAddress string 
}