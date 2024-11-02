package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Student struct {
	ID        int     `gorm:"primaryKey;size:11"`
	FirstName string  `gorm:"size:50"`
	LastName  string  `gorm:"size:50"`
	Age       int     `gorm:"type:smallint"`
	Email     *string // Pointer to allow null values
}

func CrateStudentTable(db *gorm.DB) {
	// Create the table
	db.AutoMigrate(&Student{})
}

func InsertStudent(db *gorm.DB) {

	email := "sdasd@gmail.com"
	var students = []Student{
		{ID: 1, FirstName: "John", LastName: "Doe", Age: 20, Email: nil},
		{ID: 2, FirstName: "Jane", LastName: "Smith", Age: 22, Email: nil},
		{ID: 3, FirstName: "Mike", LastName: "Johnson", Age: 21, Email: nil},
		{ID: 4, FirstName: "Sara", LastName: "Wilson", Age: 23, Email: &email},
	}

	// Insert students
	db.Create(&students)
}

func SingleQuery(db *gorm.DB) {
	// Query a single student
	var student1 Student
	var student2 Student
	var student3 Student

	db.Take(&student1) // SELECT * FROM `students` LIMIT 1
	fmt.Println(student1)

	db.First(&student2) // SELECT * FROM `students` ORDER BY `students`.`id` LIMIT 1
	fmt.Println(student2)

	db.Last(&student3) // SELECT * FROM `students` ORDER BY `students`.`id` DESC LIMIT 1
	fmt.Println(student3)
}
