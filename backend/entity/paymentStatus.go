package entity

import "gorm.io/gorm"

type paymentStatus struct {
	gorm.Model
	Status string
}
