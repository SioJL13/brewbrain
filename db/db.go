package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Couldnt connect to the Db!")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createBrewsTable := `
	CREATE TABLE IF NOT EXISTS brews (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		coffeeName TEXT,
		coffeeType TEXT,
		coffeeGrams DECIMAL NOT NULL,
		grindSize INTEGER NOT NULL,
		waterGrams DECIMAL NOT NULL,
		brewingMethod TEXT NOT NULL,
		brewTime INTEGER NOT NULL,
		extractionTime INTEGER, 
		waterTemp DECIMAL,
		grinderType TEXT,
		createdAt DATETIME NOT NULL
	)
	`

	_, err := DB.Exec(createBrewsTable)
	if err != nil {
		panic("cant create brews table: " + err.Error())
	}
}
