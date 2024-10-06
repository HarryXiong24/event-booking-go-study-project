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
		panic("could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

func createTables() {

	if DB == nil {
		panic("could not connect to database")
	}

	createEventTable := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER
	)`

	stmt, err := DB.Prepare(createEventTable)
	if err != nil {
		panic("could not create events table")
	}
	defer stmt.Close()

	stmt.Exec()
}
