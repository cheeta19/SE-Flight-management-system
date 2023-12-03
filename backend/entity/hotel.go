package entity

import "gorm.io/gorm"

type Hotel struct {
	gorm.Model
	Name         string
	Type         string
	Guest        int
	Price        int
	Location     string
	Descript     string
	HotelClass   string
	Image        string
	Service      string
	NumberOfRoom int

	HotelTypeID *uint
	HotelType   HotelType `gorm:"foreignKey:HotelTypeID"`
}
