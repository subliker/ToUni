package datavalidator

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func JsonString(msg string) string {
	replacer := strings.NewReplacer(`"`, `\"`, `\`, `\\`, `/`, `\/`)
	return replacer.Replace(fmt.Sprint(msg))
}

func SendMessage(c *gin.Context, httpCode int, msg string) {
	text := fmt.Sprintf(`
	{
		"message":"%s"
	}
	`, JsonString(msg))
	c.Data(httpCode, "application/json", []byte(text))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
