package entity

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	CountryName string
}
