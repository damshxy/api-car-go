package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secret")

type CustomValidationError struct {
	Msg string
}

func (e *CustomValidationError) Error() string {
	return e.Msg
}

func GenerateJWT(id int, name, phone string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["name"] = name
	claims["phone"] = phone
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &CustomValidationError{Msg: "unexpected signing method"}
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, &CustomValidationError{Msg: "invalid token"}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, &CustomValidationError{Msg: "invalid claims"}
	}

	return claims, nil
}