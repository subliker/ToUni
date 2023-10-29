package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/edgedb/edgedb-go"
	"github.com/gin-gonic/gin"
	"github.com/subliker/ToUni/go-server/db"
	"github.com/subliker/ToUni/go-server/model"
	"github.com/subliker/ToUni/go-server/pkg/auth"
	dv "github.com/subliker/ToUni/go-server/pkg/datavalidator"
)

type SignUpRequest struct {
	user     model.User
	username string
	password string
}

func (c *SignUpRequest) SetUsername(username string, db *db.DataBase) error {
	exist, err := db.UserController.ExistUsername(username)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("username already exists")
	}
	c.username = username
	return nil
}

func (c *SignUpRequest) GetUsername() string { return c.username }

func (c *SignUpRequest) SetPassword(password string) error {
	password, errH := dv.HashPassword(password)
	if errH != nil {
		return errH

	}
	c.password = password
	return nil
}

func (c *SignUpRequest) GetPassword() string { return c.password }

func (c *SignUpRequest) GetUser() model.User {
	c.user.Username = c.GetUsername()
	c.user.Password = c.GetPassword()
	c.user.CreatedAt = time.Now()
	c.user.UpdatedAt = time.Now()
	c.user.Role = "User"
	return c.user
}

type SignUpResponse struct {
	UserId   string `json:"user_id"`
	JwtToken string `json:"jwt_token"`
}

func (c *Client) SignUp(ctx *gin.Context) {
	var req SignUpRequest
	err := req.SetUsername(ctx.PostForm("username"), c.db)
	if err != nil {
		dv.SendMessage(ctx, http.StatusBadRequest, fmt.Sprintf("%s", err))
		return
	}
	err = req.SetPassword(ctx.PostForm("password"))
	if err != nil {
		dv.SendMessage(ctx, http.StatusBadRequest, fmt.Sprintf("%s", err))
		return
	}
	idData, httpCode, errA := c.db.UserController.Add(req.GetUser())
	if errA != nil {
		dv.SendMessage(ctx, http.StatusBadRequest, fmt.Sprintf("%s", errA))
		return
	}
	var res SignUpResponse
	res.UserId = idData.InsertUser[0].Id
	res.JwtToken, err = auth.GenerateJWT(res.UserId)
	if err != nil {
		dv.SendMessage(ctx, http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}
	data, errM := json.Marshal(res)
	if errM != nil {
		dv.SendMessage(ctx, http.StatusInternalServerError, fmt.Sprintf("%s", errM))
		return
	}
	ctx.Data(httpCode, "application/json", data)
}

type SignInRequest struct {
	username string
	password string
}

func (c *SignInRequest) SetUsername(username string, db *db.DataBase) (int, error) {
	exist, err := db.UserController.ExistUsername(username)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if !exist {
		return http.StatusBadRequest, errors.New("doesn't exist")
	}
	c.username = username
	return 0, nil
}

func (c *SignInRequest) GetUsername() string { return c.username }

func (c *SignInRequest) SetPassword(password string, username string, db *db.DataBase) (int, edgedb.UUID, error) {
	hash, id, err := db.UserController.GetHashedPasswordByUsername(username)
	if err != nil {
		return http.StatusInternalServerError, id, err
	}
	err = dv.ComparePassword(hash, password)
	if err != nil {
		return http.StatusBadRequest, id, errors.New("doesn't exist")
	}
	c.password = password
	return 0, id, nil
}

func (c *SignInRequest) GetPassword() string { return c.password }

type SignInResponse struct {
	UserId   string `json:"user_id"`
	JwtToken string `json:"jwt_token"`
}

func (c *Client) SignIn(ctx *gin.Context) {
	var req SignInRequest
	username := ctx.PostForm("username")
	httpCode, err := req.SetUsername(username, c.db)
	if err != nil {
		dv.SendMessage(ctx, httpCode, fmt.Sprint(err))
		return
	}
	var id edgedb.UUID
	httpCode, id, err = req.SetPassword(ctx.PostForm("password"), username, c.db)
	if err != nil {
		dv.SendMessage(ctx, httpCode, fmt.Sprint(err))
		return
	}
	var res SignInResponse
	res.UserId = id.String()
	res.JwtToken, err = auth.GenerateJWT(id.String())
	if err != nil {
		dv.SendMessage(ctx, http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}
	data, errM := json.Marshal(res)
	if errM != nil {
		dv.SendMessage(ctx, http.StatusInternalServerError, fmt.Sprintf("%s", errM))
		return
	}
	ctx.Data(httpCode, "application/json", data)
}
