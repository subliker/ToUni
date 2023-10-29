package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateJWT(username string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	fmt.Print(os.Getenv("SECRET_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(14 * 24 * time.Hour)
	claims["authorized"] = true
	claims["user"] = username
	tokenString, errT := token.SignedString(secretKey)
	if errT != nil {
		return "", errT
	}
	return tokenString, nil
}
