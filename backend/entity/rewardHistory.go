package entity

import "gorm.io/gorm"

type RewardHistory struct {
	gorm.Model
	CoinPay  int
	Quantity int

	MemberID *uint
	Member   Member `gorm:"foreignKey:MemberID"`

	RewardID *uint
	Reward   Reward `gorm:"foreignKey:RewardID"`
}
