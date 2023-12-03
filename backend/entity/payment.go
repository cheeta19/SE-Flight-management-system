package entity

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	BankName      string
	AccountNumber string
	Date          time.Time
	Slip          string
	Price         float64

	paymentStatusID *uint
	paymentStatus   paymentStatus `gorm:"foreignKey:paymentStatusID"`

	PaymentTypeID *uint
	PaymentType   PaymentType `gorm:"foreignKey:PaymentTypeID"`

	BookFlightID *uint
	BookFlight   BookFlight `gorm:"foreignKey:BookFlightID"`

	BookHotelID *uint
	BookHotel   BookHotel `gorm:"foreignKey:BookHotelID"`

	MemberID *uint
	Member   Member `gorm:"foreignKey:MemberID"`
}
