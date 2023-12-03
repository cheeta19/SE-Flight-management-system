package entity

import "gorm.io/gorm"

type Reward struct {
	gorm.Model
	Descript    string
	Quantity    int
	CoinAmount  int
	RewardImage string
}
