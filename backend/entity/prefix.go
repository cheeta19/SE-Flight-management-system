package entity

import "gorm.io/gorm"

type Prefix struct {
	gorm.Model
	PrefixType string
}
