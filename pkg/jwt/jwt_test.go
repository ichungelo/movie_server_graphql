package jwt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectedPayload = Payload{
	ID: 0,
	Username: "testing",
	Email: "testing@testing.com",
	FirstName: "testing",
	LastName: "testing",
}

var secretKey = []byte("secret test")
func TestGenerateTokenSuccess(t *testing.T) {
	token,err := GenerateToken(expectedPayload, secretKey)
	assert.NotNil(t, token)
	assert.Nil(t, err)
}

// func TestGenerateTokenError(t *testing.T) {
// 	token, err := GenerateToken(payload, nil)
// 	assert.Equal(t, "", token)
// 	assert.NotNil(t, err)
// }

func TestParseTokenSuccess (t *testing.T) {
	token, _ := GenerateToken(expectedPayload, secretKey)

	payload, err := ParseToken("Bearer " + token, secretKey)
	assert.Nil(t, err)
	assert.Equal(t, expectedPayload.ID, payload.ID)
	assert.Equal(t, expectedPayload.Username, payload.Username)
	assert.Equal(t, expectedPayload.Email, payload.Email)
	assert.Equal(t, expectedPayload.FirstName, payload.FirstName)
	assert.Equal(t, expectedPayload.LastName, payload.LastName)
}

func TestParseTokenError (t *testing.T) {
	token, _ := GenerateToken(expectedPayload, secretKey)

	expectedError := []struct{
		token string
		secretKey []byte
	}{
		{
			token: token,
			secretKey: secretKey,
		},
		{
			token: fmt.Sprint("falseBearer " + token),
			secretKey: secretKey,
		},
		{
			token: fmt.Sprint("Bearer " + token),
			secretKey: []byte("false secret"),
		},
	}

	for _, error := range expectedError {
		payload, err := ParseToken(error.token, error.secretKey)
		assert.Equal(t, Payload{}, payload)
		assert.NotNil(t, err)
	}
}