package jwt

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Payload struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateToken(payload Payload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": payload.ID,
		"username": payload.Username,
		"email": payload.Email,
		"first_name": payload.FirstName,
		"last_name": payload.LastName,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})
	
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(token string) (Payload, error) {
	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return Payload{}, err
	}

	if claim, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {
		claims := Payload{
			ID: claim["id"].(int64),
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