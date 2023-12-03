package entity

import (
	"time"

	"gorm.io/gorm"
)

type BookFlight struct {
	gorm.Model
	FromDate time.Time
	Quantity int
	Price    int

	MemberID *uint
	Member   Member `gorm:"foreignKey:MemberID"`

	FlightID *uint
	Flight   Flight `gorm:"foreignKey:FlightID"`

	CheckID *uint
	Check   Check `gorm:"foreignKey:CheckID"`

	FoodFlightID *uint
	FoodFlight   FoodFlight `gorm:"foreignKey:FoodFlightID"`

	BaggegID *uint
	Baggeg   Baggeg `gorm:"foreignKey:BaggegID"`
}
