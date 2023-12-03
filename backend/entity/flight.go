package entity

import (
	"time"

	"gorm.io/gorm"
)

type Flight struct {
	gorm.Model
	Depart     string
	Startdate  time.Time
	Arrivedate time.Time
	Price      int

	CompanyID *uint
	Company   Company `gorm:"foreignKey:CompanyID"`

	OriginID *uint
	Origin   Origin `gorm:"foreignKey:OriginID"`

	DestinationID *uint
	Destination   Destination `gorm:"foreignKey:DestinationID"`
}
