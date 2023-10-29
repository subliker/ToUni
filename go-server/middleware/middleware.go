package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/subliker/ToUni/go-server/db"
	"github.com/subliker/ToUni/go-server/pkg/auth"
	dv "github.com/subliker/ToUni/go-server/pkg/datavalidator"
)

type Middleware struct {
	db *db.DataBase
}

func (c *Middleware) SetDataBase(db *db.DataBase) {
	c.db = db
}

func (c *Middleware) CheckRole(ctx *gin.Context) {
	tokenText := strings.Replace(ctx.Request.Header["Authorization"][0], "Bearer ", "", 1)
	claims := auth.UserClaim{}
	_, err := jwt.ParseWithClaims(tokenText, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		dv.SendMessage(ctx, http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	role, errR := c.db.UserController.GetRoleById(claims.Id)
	if errR != nil {
		dv.SendMessage(ctx, http.StatusInternalServerError, fmt.Sprint(errR))
		return
	}
	fmt.Println(role)
	if role != "Admin" {
		dv.SendMessage(ctx, http.StatusInternalServerError, "you aren't admin")
		ctx.Abort()
		return
	}
	ctx.Next()
}
