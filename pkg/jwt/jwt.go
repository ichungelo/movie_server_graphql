package jwt

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)


var secretKey = []byte(os.Getenv("SECRET_KEY"))
type Payload struct {
	ID int `json:"id"`
	Username string `json:"username"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}
func GenerateToken(payload Payload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": payload.ID,
		"username": payload.Username,
		"first_name": payload.FirstName,
		"last_name": payload.LastName,
		"exp": time.Now().Add(time.Minute * 1).Unix(),
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
			ID: int(claim["id"].(float64)),
			Username: claim["username"].(string),
			FirstName: claim["first_name"].(string),
			LastName: claim["last_name"].(string),
		}
		return claims, nil
	} else {
		return Payload{}, err
	}
}