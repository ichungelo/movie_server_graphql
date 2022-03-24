package users

import (
	"fmt"
	mysql "movie_graphql_be/pkg/db"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

func UsernameValidator(username string) (bool, error) {
	if len(username) < 6 {
		return false, fmt.Errorf("username must be at least 6 characters")
	}

	state, err := mysql.Db.Prepare("SELECT username FROM users WHERE username = ?")
	if err != nil {
		return false, err
	}

	rows, err := state.Query(username)
	if err != nil {
		return false, err
	}

	defer rows.Close()

	if rows.Next() {
		return true, fmt.Errorf("username is already in use")
	}

	return false, nil
}

func EmailValidator(email string) (bool, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false, fmt.Errorf("email is not valid")
	}

	state, err := mysql.Db.Prepare("SELECT email FROM users WHERE email = ?")
	if err != nil {
		return false, err
	}

	rows, err := state.Query(email)
	if err != nil {
		return false, err
	}

	defer rows.Close()

	if rows.Next() {
		return true, fmt.Errorf("email is already in use")
	}

	return false, nil
}

func HashPassword(password string, confirm string) (string, error) {
	if password != confirm {
		return "", fmt.Errorf("password and confirm password must be same")
	}

	//password validation
	if len(password) < 8 {
		return "", fmt.Errorf("password must be at least 8 characters")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (login *Login)AuthPassword() (bool, error) {
	state, err := mysql.Db.Prepare("SELECT password FROM users WHERE username = ?")
	if err != nil {
		return false, err
	}

	row := state.QueryRow(login.Username)
	var password string

	err = row.Scan(&password)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(login.Password))
	if err != nil {
		return false, err
	}

	return true, nil
}