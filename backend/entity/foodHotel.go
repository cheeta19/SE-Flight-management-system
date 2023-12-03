package entity

import "gorm.io/gorm"

type FoodHotel struct {
	gorm.Model
	Price       int
	Description string

	HotelID *uint
	Hotel   Hotel `gorm:"foreignKey:HotelID"`
}
