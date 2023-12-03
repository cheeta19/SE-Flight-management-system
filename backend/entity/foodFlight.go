package entity

import "gorm.io/gorm"

type FoodFlight struct {
	gorm.Model
	Dessert    string
	Drink      string
	Price      int
	MainCourse string
}
