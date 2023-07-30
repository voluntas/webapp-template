package webapp

import (
	"crypto/rand"
	"encoding/base64"
)

func generateRandomBytes(n uint) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomString(n uint) (string, error) {
	b, err := generateRandomBytes(n)
	var RawURLEncoding = base64.URLEncoding.WithPadding(base64.NoPadding)
	return RawURLEncoding.EncodeToString(b), err
}
