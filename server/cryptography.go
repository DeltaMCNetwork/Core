package server

import (
	"crypto/rand"
)

func GenerateVerificationToken() []byte {
	token := make([]byte, 4)
	rand.Read(token)

	return token
}
