package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		fmt.Println(err)
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

	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	)`

	stmt, err := DB.Prepare(createUsersTable)
	if err != nil {
		fmt.Println(err)
		panic("could not create users table")
	}
	defer stmt.Close()
	stmt.Exec()

	createEventTable := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	stmt, err = DB.Prepare(createEventTable)
	if err != nil {
		fmt.Println(err)
		panic("could not create events table")
	}
	defer stmt.Close()
	stmt.Exec()
}
