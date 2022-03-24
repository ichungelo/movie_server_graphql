package users

import (
	"fmt"
	"log"
	mysql "movie_graphql_be/pkg/db"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (user *User) CreateUser() (string, error) {
	state, err := mysql.Db.Prepare("INSERT INTO users (username, email, first_name, last_name, password) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return "", err
	}

	hashedPass, err := HashPassword(user.Password, user.ConfirmPassword)
	if err != nil {
		log.Println(err)
		return "", err
	}

	_, err = state.Exec(user.Username, user.Email, user.FirstName, user.LastName, hashedPass)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return "Register Success", nil
}

func HashPassword(password string, confirm string) (string, error) {
	if password != confirm {
		return "", fmt.Errorf("password and confirm password must be same")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}