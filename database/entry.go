package database

import "gorm.io/gorm"

type Entry struct {
	gorm.Model
	DateRefer uint
	Time      Date `gorm:"foreignKey:DateRefer"`
	Content   string
}
