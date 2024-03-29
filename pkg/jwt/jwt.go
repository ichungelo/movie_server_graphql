package jwt

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Payload struct {
	ID        float64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateToken(payload Payload, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": payload.ID,
		"username": payload.Username,
		"email": payload.Email,
		"first_name": payload.FirstName,
		"last_name": payload.LastName,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})
	
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(bearerToken string) (Payload, error) {
	token := strings.Split(bearerToken, " ")
	if len(token) <= 1 {
		return Payload{}, fmt.Errorf("token is invalid")
	}

	if token[0] != "Bearer" {
		return Payload{}, fmt.Errorf("token is invalid")
	}

	tkn, err := jwt.Parse(token[1], func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		log.Println(err)
		return Payload{}, err
	}

	if claim, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {
		claims := Payload{
			ID: claim["id"].(float64),
			Username: claim["username"].(string),
			Email: claim["email"].(string),
			FirstName: claim["first_name"].(string),
			LastName: claim["last_name"].(string),
		}
		return claims, nil
	} else {
		return Payload{}, err
	}
}