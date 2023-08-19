package models

import "gorm.io/gorm"

type LoyaltyPointsDtl struct {
	gorm.Model
	Coins         int16
	StartingRange int16
	EndingRange   int16
}
