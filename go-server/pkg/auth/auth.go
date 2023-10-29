package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type UserClaim struct {
	jwt.RegisteredClaims
	Id string
}

func GenerateJWT(id string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	exp := jwt.NumericDate{Time: time.Now().Add(time.Hour * 24 * 14)}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: &exp},
		Id:               id,
	})
	tokenString, errT := token.SignedString(secretKey)
	if errT != nil {
		return "", errT
	}
	return tokenString, nil
}
