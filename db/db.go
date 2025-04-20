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

	createTables()
}

func createTables() {
	createBrewsTable := `
	CREATE TABLE IF NOT EXISTS brews (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		coffee_name TEXT,
		coffee_type TEXT,
		coffee_grams DECIMAL NOT NULL,
		grind_size INTEGER NOT NULL,
		water_grams DECIMAL NOT NULL,
		brewing_method TEXT NOT NULL,
		brew_time INTEGER NOT NULL,
		extraction_time INTEGER, 
		water_temp DECIMAL,
		grinder_type TEXT,
		created_at DATETIME NOT NULL
	)
	`

	tx := DB.Exec(createBrewsTable)
	if tx.Error != nil {
		panic("cant create brews table: " + tx.Error.Error())
	}
}
