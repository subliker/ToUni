package middleware

import "github.com/gin-gonic/gin"

var jwtSecretKey = []byte("abc123")

func CheckRole(c *gin.Context) {
	c.Next()
}
