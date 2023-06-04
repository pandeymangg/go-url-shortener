package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomString(length int) string {
	var alphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var randomString []byte

	for i := 0; i < length; i++ {
		randomNumber, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphaNumeric))))

		if err != nil {
			panic(err)
		}

		randomString = append(randomString, alphaNumeric[randomNumber.Int64()])
	}

	return string(randomString)
}
