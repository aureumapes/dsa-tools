package database

import (
	"gorm.io/gorm"
)

type Date struct {
	gorm.Model
	Day   int64
	Month string `gorm:"type:months"`
	Year  int64
}
