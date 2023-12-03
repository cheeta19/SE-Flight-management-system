package entity

import (
	"time"

	"gorm.io/gorm"
)

type BookHotel struct {
	gorm.Model
	DateIn  time.Time
	DateOut time.Time
	Price   int

	MemberID *uint
	Member   Member `gorm:"foreignKey:MemberID"`

	HotelID *uint
	Hotel   Hotel `gorm:"foreignKey:HotelID"`

	FoodHotelID *uint
	FoodHotel   FoodHotel `gorm:"foreignKey:FoodHotelID"`

	RoomID *uint
	Room   Room `gorm:"foreignKey:RoomID"`
}
