package models

import "gorm.io/gorm"

type WalletTransactionData struct {
	gorm.Model
	FromWalletAddress string
	ToWalletAddress   string
	Coins             int16
}
