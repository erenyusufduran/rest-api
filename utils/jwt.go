package utils

import (
	"errors"
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

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// _, ok := t.Method.(*jwt.SigningMethodHMAC)
		// if !ok {
		// 	return nil, errors.New("unexpected signing method")
		// }
		privateKey, err := readPrivateKeyFromFile(fileName)
		if err != nil {
			return nil, err
		}
		return extractECDSAPublicKey(privateKey), nil
	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	return userId, nil
}
