package database

type Adventure struct {
	Number string `gorm:"primarykey"`
	Name   string
	Heroes string
	Time   string
}
