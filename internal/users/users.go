package users

import (
	"fmt"
	"log"
	"movie_graphql_be/pkg/db/mysql"
	"movie_graphql_be/pkg/jwt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

type User struct {
	ID              float64 `json:"id"`
	Username        string  `json:"username"`
	Email           string  `json:"email"`
	FirstName       string  `json:"first_name"`
	LastName        string  `json:"last_name"`
	Password        string  `json:"password"`
	ConfirmPassword string  `json:"confirm_password"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) CreateUser() (string, error) {
	state, err := mysql.Db.Prepare("INSERT INTO users (username, email, first_name, last_name, password) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return "", err
	}

	if err := UsernameValidator(user.Username); err != nil {
		log.Println(err)
		return "", err
	}

	if err := EmailValidator(user.Email); err != nil {
		log.Println(err)
		return "", err
	}

	hashedPass, err := HashPassword(user.Password, user.ConfirmPassword)
	if err != nil {
		log.Println(err)
		return "", err
	}

	affect, err := state.Exec(user.Username, user.Email, user.FirstName, user.LastName, hashedPass)
	if err != nil {
		log.Println(err)
		return "", err
	}

	result, err := affect.RowsAffected()
	if err != nil {
		log.Println(err)
		return "", nil
	}

	if result == 0 {
		log.Println(err)
		return "", fmt.Errorf("error add user")
	}

	return "Register Success", nil
}

func (login *Login) LoginUser() (string, error) {

	state, err := mysql.Db.Prepare("SELECT id, username, email, first_name, last_name, password FROM users WHERE username = ?")
	if err != nil {
		log.Println(err)
		return "", err
	}

	if err := GetUserByUsername(login.Username); err != nil {
		log.Println(err)
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
		log.Println(err)
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(login.Password))
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("password incorrect")
	}

	payload := jwt.Payload{
		ID:        result.ID,
		Username:  result.Username,
		Email:     result.Email,
		FirstName: result.FirstName,
		LastName:  result.LastName,
	}

	token, err := jwt.GenerateToken(payload, secretKey)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return token, nil
}

func GetUserByUsername(username string) error {
	state, err := mysql.Db.Prepare("SELECT username FROM users WHERE username = ?")
	if err != nil {
		log.Println(err)
		return err
	}

	rows, err := state.Query(username)
	if err != nil {
		log.Println(err)
		return err
	}

	defer rows.Close()

	if rows.Next() {
		return nil
	}

	return fmt.Errorf("username not exist")
}
