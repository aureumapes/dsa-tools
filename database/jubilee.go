package database

import "gorm.io/gorm"

type Jubilee struct {
	gorm.Model
	Day   int64
	Month string `gorm:"type:months"`
	Year  int64
	Name  string
}
