package entity

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Descript string
	Date     time.Time
	Rating   string

	MemberID *uint
	Member   Member `gorm:"foreignKey:MemberID"`

	HotelID *uint
	Hotel   Hotel `gorm:"foreignKey:HotelID"`
}
