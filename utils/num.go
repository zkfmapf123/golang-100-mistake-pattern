package utils

import (
	"crypto/rand"
)

func RandBytes() []byte {
	rb := make([]byte, 16)
	rand.Read(rb)

	return rb
}
