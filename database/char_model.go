package database

type Char struct {
	Name        string `gorm:"primarykey"`
	TDay        string
	BDay        string
	Type        string
	Description string
	Member      string
}
