package database

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Name    string
	Content string `gorm:"type:OID"`
}
