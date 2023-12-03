package entity

import (
	"time"

	"gorm.io/gorm"
)

type Check struct {
	gorm.Model
	state string
	Date  time.Time
}
