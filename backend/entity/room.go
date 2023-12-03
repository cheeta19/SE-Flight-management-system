package entity

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Type  string
	Bed   string
	Price int

	HotelID *uint
	Hotel   Hotel `gorm:"foreignKey:HotelID"`
}
