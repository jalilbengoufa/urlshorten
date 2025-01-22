package service

import (
	"crypto/rand"
	"encoding/base64"
	"urlshorten/internal/store"
)

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s - 2)
	return base64.URLEncoding.EncodeToString(b), err
}

func GenerateShortCode() (code string, err error) {

	for {

		code, err := generateRandomString(6)
		if err != nil {
			return "", err
		}

		_, ok := store.DataStore.Load(code)
		if !ok {
			return code, nil
		}

	}

}
