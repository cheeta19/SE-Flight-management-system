package entity

import "gorm.io/gorm"

type Banner struct {
	gorm.Model
	Link  string
	Image string
}
