package database

//go:generate go run github.com/dmarkham/enumer -type=Month

import (
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	DateRefer uint
	Time      Date `gorm:"foreignKey:DateRefer"`
	Content   string
}

type Date struct {
	gorm.Model
	Day   int64
	Month string `gorm:"type:months"`
	Year  int64
}

type Jubilee struct {
	gorm.Model
	Day   int64
	Month string `gorm:"type:months"`
	Year  int64
	Name  string
}
