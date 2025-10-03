package models

import (
	"errors"

	"github.com/PetarGeorgiev-hash/learning-go/db"
	"github.com/PetarGeorgiev-hash/learning-go/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := `
		INSERT INTO users(email,password) VALUES (?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	hashPass, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashPass)

	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	return err

}

func (user *User) ValidateCredentials() error {
	query := `
		SELECT id,password FROM users WHERE email = ?
	`
	row := db.DB.QueryRow(query, user.Email)
	var pass string
	err := row.Scan(&user.Id, &pass)
	if err != nil {
		return err
	}
	passwordIsValid := utils.CompareHashedPassword(user.Password, pass)
	if !passwordIsValid {
		return errors.New("Invalid credentails")
	}
	return nil
}
