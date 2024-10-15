package models

import (
	"errors"
	"fmt"

	"api.com/db"
	"api.com/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	// Save the user to the database
	query := `
	INSERT INTO users(email, password) 
	VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := res.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = userId
	return nil
}

func (u *User) ValidateCredentials() error {
	query := `select id, password from users where email = ?`

	var retrievedPassword string
	row := db.DB.QueryRow(query, u.Email)

	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		fmt.Println(err)
		return errors.New("invalid password")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		fmt.Println(err)
		return errors.New("invalid password")
	}

	return nil
}
