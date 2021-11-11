package authentication

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/martian/log"
	"github.com/wasimkhan042/ustore-auth/config"
	"strings"
	"time"
)

// ValidateHeader for token validation.
func ValidateHeader(bearerHeader string) (interface{}, error) {
	bearerToken := strings.Split(bearerHeader, " ")[1]
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(bearerToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error decoding token")
		}
		return config.JWTSecretKey, nil
	})
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	if token.Valid {
		return claims["user"].(string), nil
	}
	return nil, errors.New("invalid token ...")
}

// GenerateJWT for token generation.
func GenerateJWT(userEmail string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = userEmail
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(config.JWTSecretKey)
	if err != nil {
		fmt.Println("reached")
		return "", err
	}
	return tokenString, nil
}

func ServiceGenerateJWT(itemName, pkgCode string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["item"] = itemName
	claims["code"] = pkgCode
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(config.JWTSecretKey)
	if err != nil {
		fmt.Println("reached")
		return "", err
	}
	return tokenString, nil
}
