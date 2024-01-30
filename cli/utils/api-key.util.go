package utils

import (
	"crypto/rand"
	"encoding/base64"
)

const apiKeyLength = 16

func GenerateAPIKey() (string, error) {
	key := make([]byte, apiKeyLength)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	encodedKey := base64.URLEncoding.EncodeToString(key)

	return encodedKey, nil
}
