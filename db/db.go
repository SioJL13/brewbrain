package db

import (
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		panic("Couldnt connect to the Db!")
	}

}
