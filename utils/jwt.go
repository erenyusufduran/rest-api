package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const fileName = "private.pem"

func GenerateToken(email string, userId int64) (string, error) {
	privKey, err := readPrivateKeyFromFile(fileName)
	if err != nil {
		privKey, err = generateKey()
		if err != nil {
			return "", err
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString(privKey)
}
