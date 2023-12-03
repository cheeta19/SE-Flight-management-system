package entity

import "gorm.io/gorm"

type Origin struct {
	gorm.Model
	OriginName string
}
