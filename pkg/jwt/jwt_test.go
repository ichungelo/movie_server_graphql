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

var dummySecret = []byte("secret test")
func TestGenerateTokenSuccess(t *testing.T) {
	token,err := GenerateToken(expectedPayload, dummySecret)
	assert.NotNil(t, token)
	assert.Nil(t, err)
}

// func TestGenerateTokenError(t *testing.T) {
// 	token, err := GenerateToken(payload, nil)
// 	assert.Equal(t, "", token)
// 	assert.NotNil(t, err)
// }

func TestParseTokenSuccess (t *testing.T) {
	token, _ := GenerateToken(expectedPayload, dummySecret)

	payload, err := ParseToken("Bearer " + token)
	assert.Nil(t, err)
	assert.Equal(t, expectedPayload.ID, payload.ID)
	assert.Equal(t, expectedPayload.Username, payload.Username)
	assert.Equal(t, expectedPayload.Email, payload.Email)
	assert.Equal(t, expectedPayload.FirstName, payload.FirstName)
	assert.Equal(t, expectedPayload.LastName, payload.LastName)
}

func TestParseTokenError (t *testing.T) {
	token, _ := GenerateToken(expectedPayload, dummySecret)

	expectedError := []struct{
		token string
		secretKey []byte
	}{
		{
			token: token,
			secretKey: dummySecret,
		},
		{
			token: fmt.Sprint("falseBearer " + token),
			secretKey: dummySecret,
		},
		{
			token: fmt.Sprint("Bearer " + token),
			secretKey: []byte("false secret"),
		},
	}

	for _, error := range expectedError {
		payload, err := ParseToken(error.token)
		assert.Equal(t, Payload{}, payload)
		assert.NotNil(t, err)
	}
}