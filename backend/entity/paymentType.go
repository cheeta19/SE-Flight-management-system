package entity

import "gorm.io/gorm"

type PaymentType struct {
	gorm.Model
	TypeName int
}
