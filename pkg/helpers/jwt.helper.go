package helpers

import (
	"strings"
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

func GenerateJWT(id uint, email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	if strings.HasPrefix(tokenString, "Bearer") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &CustomValidationError{
				Msg: "invalid token",
			}
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, &CustomValidationError{
			Msg: "invalid token",
		}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, &CustomValidationError{
			Msg: "invalid token",
		}
	}

	return claims, nil
}