package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connect() (*gorm.DB, error) {
	// Open a database connection
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {

	db, err := connect()
	if err != nil {
		panic("failed to connect database")
	}

	// CrateStudentTable(db)
	// InsertStudent(db)

	// SingleQuery(db)
	MultipleQuery(db)
}
