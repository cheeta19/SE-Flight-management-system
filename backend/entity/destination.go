package entity

import "gorm.io/gorm"

type Destination struct {
	gorm.Model
	DestinationName string
}
