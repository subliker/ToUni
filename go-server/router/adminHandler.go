package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/edgedb/edgedb-go"
	"github.com/gin-gonic/gin"
	"github.com/subliker/ToUni/go-server/model"
	dv "github.com/subliker/ToUni/go-server/pkg/datavalidator"
)

func (c *User) GetAll(ctx *gin.Context) {
	res, httpCode, err := c.db.UserController.GetAll()
	if err != nil {
		log.Panic(err)
	}

	usersData, errm := json.MarshalIndent(res, "", "   ")
	if errm != nil {
		log.Panic(errm)
	}
	ctx.Data(httpCode, "application/json", usersData)
}

func (c *User) GetOneById(ctx *gin.Context) {
	var id edgedb.UUID
	id.UnmarshalText([]byte(ctx.Param("id")))
	res, httpCode, err := c.db.UserController.GetOneById(id)
	if err != nil {
		log.Panic(err)
	}

	userData, errm := json.MarshalIndent(res.User[0], "", "   ")
	if errm != nil {
		log.Panic(errm)
	}
	ctx.Data(httpCode, "application/json", userData)
}

type AddUserResponse struct {
	Id string `json:"id"`
}

func (c *User) Add(ctx *gin.Context) {
	var user model.User
	user.Username = ctx.PostForm("username")
	password, errH := dv.HashPassword(ctx.PostForm("password"))
	if errH != nil {
		dv.SendMessage(ctx, http.StatusBadRequest, fmt.Sprintf("%s", errH))
	}
	createdAt := time.Now()
	user.Password = password
	user.CreatedAt = createdAt
	user.UpdatedAt = createdAt
	idData, httpCode, errA := c.db.UserController.Add(user)
	if errA != nil {
		dv.SendMessage(ctx, http.StatusBadRequest, fmt.Sprintf("%s", errA))
		return
	}
	var res AddUserResponse
	res.Id = idData.InsertUser[0].Id
	data, errM := json.Marshal(res)
	if errM != nil {
		dv.SendMessage(ctx, http.StatusInternalServerError, fmt.Sprintf("%s", errM))
		return
	}
	ctx.Data(httpCode, "application/json", data)
}
