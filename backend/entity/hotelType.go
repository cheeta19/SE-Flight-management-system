package entity

import "gorm.io/gorm"

type HotelType struct {
	gorm.Model
	Type string
}
