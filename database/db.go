package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// var Db, _ = gorm.Open(sqlite.Open("resource/dsa.sqlite"), &gorm.Config{})
var Db, _ = gorm.Open(postgres.Open(fmt.Sprintf("host=localhost user=artur password=%s dbname=dsa port=5432 sslmode=disable TimeZone=Europe/Berlin", os.Getenv("PASSWORT"))), &gorm.Config{})
