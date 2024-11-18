package main

import (
	"encoding/json"
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

func CreateStudentTable(db *gorm.DB) {
	// Create the table
	err := db.AutoMigrate(&Student{})
	if err != nil {
		return
	}
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

	// Query a single student by primary key
	var student4 Student
	db.Take(&student4, 2) // SELECT * FROM `students` WHERE `students`.`id` = 2
	fmt.Println(student4)

	// Query a single student by condition
	var student5 Student
	db.Take(&student5, "first_name = ?", "John") // SELECT * FROM `students` WHERE first_name = 'John'
	fmt.Println(student5)

	// Query a single student by map
	var student6 Student
	// only works with struct fields
	student6.ID = 2
	db.Take(&student6)
	fmt.Println(student6)
}

func MultipleQuery(db *gorm.DB) {
	var students []Student
	db.Find(&students) // SELECT * FROM `students`
	fmt.Println(students)

	// 由于 email 是指针类型，所以看不到实际的内容
	// 但是序列化之后，会转换为我们可以看得懂的方式
	var studentList []Student
	db.Find(&studentList)
	for _, student := range studentList {
		data, _ := json.Marshal(student)
		fmt.Println(string(data))
	}

	// Query multiple students by primary key
	var studentList1 []Student
	db.Find(&studentList1, []int{1, 3, 5, 7})
	db.Find(&studentList1, 1, 3, 5, 7) // same as above
	fmt.Println(studentList1)

	// Query multiple students by condition
	var studentList2 []Student
	db.Find(&studentList2, "first_name in ?", []string{"John", "Harry"})
	fmt.Println(studentList2)
}
