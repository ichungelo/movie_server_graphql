package users

import (
	"fmt"
	mysql "movie_graphql_be/pkg/db"
	"movie_graphql_be/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID              int64  `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) CreateUser() (string, error) {
	state, err := mysql.Db.Prepare("INSERT INTO users (username, email, first_name, last_name, password) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return "", err
	}

	if err := UsernameValidator(user.Username); err != nil {
		return "", err
	}

	if err := EmailValidator(user.Email); err != nil {
		return "", err
	}

	hashedPass, err := HashPassword(user.Password, user.ConfirmPassword)
	if err != nil {
		return "", err
	}

	_, err = state.Exec(user.Username, user.Email, user.FirstName, user.LastName, hashedPass)
	if err != nil {
		return "", err
	}

	return "Register Success", nil
}

func (login *Login) LoginUser() (string, error) {
	state, err := mysql.Db.Prepare("SELECT id, username, email, first_name, last_name, password FROM users WHERE username = ?")
	if err != nil {
		return "", err
	}

	if err := GetUserByUsername(login.Username); err != nil {
		return "", err
	}

	result := User{
		ID:        0,
		Username:  "",
		Email:     "",
		FirstName: "",
		LastName:  "",
		Password:  "",
	}

	err = state.QueryRow(login.Username).Scan(
		&result.ID,
		&result.Username,
		&result.Email,
		&result.FirstName,
		&result.LastName,
		&result.Password,
	)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(login.Password))
	if err != nil {
		return "", fmt.Errorf("password is not correct")
	}

	payload := jwt.Payload{
		ID:        result.ID,
		Username:  result.Username,
		Email:     result.Email,
		FirstName: result.FirstName,
		LastName:  result.LastName,
	}

	token, err := jwt.GenerateToken(payload)
	if err != nil {
		return "", err
	}

	return token, nil
}
