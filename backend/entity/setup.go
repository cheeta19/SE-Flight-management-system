package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("AirportDB.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	database.AutoMigrate(
		&Baggeg{},
		&Banner{},
		&BookFlight{},
		&BookHotel{},
		&Check{},
		&Company{},
		&Country{},
		&Destination{},
		&Flight{},
		&FoodFlight{},
		&FoodHotel{},
		&Gender{},
		&Hotel{},
		&HotelType{},
		&Member{},
		&Origin{},
		&Payment{},
		&paymentStatus{},
		&PaymentType{},
		&Prefix{},
		&Review{},
		&Reward{},
		&RewardHistory{},
		&Room{},
	)
	db = database
}
