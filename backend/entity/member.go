package entity

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Firstname      string
	Lastname       string
	Phone          string
	Email          string
	Passport       string
	Password       string
	ProfilePicture string
	Coin           int

	CountryID *uint
	Country   Country `gorm:"foreignKey:CountryID"`

	GenderID *uint
	Gender   Gender `gorm:"foreignKey:GenderID"`

	PrefixID *uint
	Prefix   Prefix `gorm:"foreignKey:PrefixID"`
}
