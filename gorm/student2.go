package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Student2 struct {
	ID     int     `gorm:"primaryKey;size:11"`
	Name   string  `gorm:"size:50"`
	Age    int     `gorm:"type:smallint"`
	Email  *string // Pointer to allow null values
	Gender bool    `gorm:"type:tinyint(1)"`
}

func CreateStudentTable2(db *gorm.DB) {
	// Create the table
	err := db.AutoMigrate(&Student2{})
	if err != nil {
		return
	}
}

func InsertStudent2(db *gorm.DB) {
	var studentList []Student2
	db.Find(&studentList).Delete(&studentList)
	studentList = []Student2{
		{ID: 1, Name: "李元芳", Age: 32, Email: parseString("lyf@yf.com"), Gender: true},
		{ID: 2, Name: "张武", Age: 18, Email: parseString("zhangwu@lly.cn"), Gender: true},
		{ID: 3, Name: "枫枫", Age: 23, Email: parseString("ff@yahoo.com"), Gender: true},
		{ID: 4, Name: "刘大", Age: 54, Email: parseString("liuda@qq.com"), Gender: true},
		{ID: 5, Name: "李武", Age: 23, Email: parseString("liwu@lly.cn"), Gender: true},
		{ID: 6, Name: "李琦", Age: 14, Email: parseString("liqi@lly.cn"), Gender: false},
		{ID: 7, Name: "晓梅", Age: 25, Email: parseString("xiaomeo@sl.com"), Gender: false},
		{ID: 8, Name: "如燕", Age: 26, Email: parseString("ruyan@yf.com"), Gender: false},
		{ID: 9, Name: "魔灵", Age: 21, Email: parseString("moling@sl.com"), Gender: true},
	}
	db.Create(&studentList)
}

func parseString(email string) *string {
	return &email
}

func Query(db *gorm.DB) {
	var users []Student2

	// 查询用户名是枫枫的
	db.Where("name = ?", "枫枫").Find(&users)
	fmt.Println(users)

	// 查询用户名不是枫枫的
	db.Where("name != ?", "枫枫").Find(&users)
	fmt.Println(users)

	// 查询用户名包含 如燕，李元芳的
	db.Where("name in ?", []string{"如燕", "李元芳"}).Find(&users)
	fmt.Println(users)

	// 查询姓李的
	db.Where("name like ?", "李%").Find(&users)
	fmt.Println(users)

	// 查询年龄大于23，是qq邮箱的
	db.Where("age > ? and email like ?", "23", "%@qq.com").Find(&users)
	fmt.Println(users)

	// 查询是qq邮箱的，或者是女的
	db.Where("gender = ? or email like ?", false, "%@qq.com").Find(&users)
	fmt.Println(users)

	// 查询是qq邮箱的，或者是男的
	db.Where("email like ?", "%@qq.com").Or("gender = ?", true).Find(&users)
	fmt.Println(users)
}

func Query2(db *gorm.DB) {
	var users []Student2

	// 这里使用的是 结构体实例作为条件。
	// 如果 Age 的值是零值（比如 0 或空字符串），GORM 默认会忽略零值字段。
	// 所以生成的 SQL 查询可能是：
	// SELECT * FROM `students` WHERE `name` = '李元芳';
	db.Where(&Student2{Name: "李元芳", Age: 0}).Find(&users)
	fmt.Println(users)

	// 这里使用的是 map[string]any 作为条件。
	// 在这种写法中，map 的所有键值对都会被视为查询条件，无论值是否为零值。
	// 所以生成的 SQL 查询是：
	// SELECT * FROM `students` WHERE `age` = 0 AND `name` = '李元芳'
	db.Where(map[string]any{"name": "李元芳", "age": 0}).Find(&users)
	fmt.Println(users)
}

type User struct {
	Name string
	Age  int
}

func Query3(db *gorm.DB) {
	var users []Student2

	// select name, age from students;
	// 但是这里的 Find 方法只会返回 name 和 age 两个字段的值，其他字段的值会赋为零值
	db.Select("name", "age").Find(&users)
	fmt.Println(users)

	var users2 []User
	// select name, age from students;
	// 两种写法都可以
	db.Model(&Student2{}).Select("name", "age").Scan(&users2)
	db.Table("student2").Select("name", "age").Scan(&users2)
	fmt.Println(users2)
}
