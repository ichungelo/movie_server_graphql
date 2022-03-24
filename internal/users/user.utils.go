package users

import (
	"fmt"
	mysql "movie_graphql_be/pkg/db"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

func UsernameValidator(username string) error {
	if len(username) < 6 {
		return fmt.Errorf("username must be at least 6 characters")
	}

	state, err := mysql.Db.Prepare("SELECT username FROM users WHERE username = ?")
	if err != nil {
		return err
	}

	rows, err := state.Query(username)
	if err != nil {
		return err
	}

	defer rows.Close()

	if rows.Next() {
		return fmt.Errorf("username is already in use")
	}

	return nil
}

func GetUserByUsername(username string) error {
	state, err := mysql.Db.Prepare("SELECT username FROM users WHERE username = ?")
	if err != nil {
		return  err
	}

	rows, err := state.Query(username)
	if err != nil {
		return  err
	}

	defer rows.Close()

	if rows.Next() {
		return nil
	}

	return fmt.Errorf("username not exist")
}

func EmailValidator(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("email is not valid")
	}

	state, err := mysql.Db.Prepare("SELECT email FROM users WHERE email = ?")
	if err != nil {
		return err
	}

	rows, err := state.Query(email)
	if err != nil {
		return err
	}

	defer rows.Close()

	if rows.Next() {
		return fmt.Errorf("email is already in use")
	}

	return nil
}

func HashPassword(password string, confirm string) (string, error) {
	if password != confirm {
		return "", fmt.Errorf("password and confirm password must be same")
	}

	if len(password) < 8 {
		return "", fmt.Errorf("password must be at least 8 characters")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
