package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/edgedb/edgedb-go"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/subliker/ToUni/go-server/db"
	"github.com/subliker/ToUni/go-server/model"
	dv "github.com/subliker/ToUni/go-server/pkg/datavalidator"
	_ "golang.org/x/crypto/bcrypt"
)

type Router struct {
	router *gin.Engine
	db     *db.DataBase
	user   User
	lesson Lesson
}

type User struct {
	db *db.DataBase
}

type Lesson struct {
	db *db.DataBase
}

func (c *Router) SetupRouter() {
	if c.db == nil {
		log.Fatal("use SetDataBase on router(DataBase isn't set)")
		return
	}
	c.router = gin.Default()
	c.user.db = c.db
	c.lesson.db = c.db

	// docs.SwaggerInfo.BasePath = "/api"

	c.router.GET("/api/user/:id", c.user.GetOneById)
	c.router.GET("/api/user", c.user.GetAll)
	c.router.POST("/api/user", c.user.Add)
	// router.POST("/api/user", route.AddNewUser)
	// router.DELETE("/api/user/:id", route.DeleteUserDataByID)
	// router.PUT("/api/user/:id", route.UpdateUserDataById)

	// router.GET("/api/booking/:id", route.GetBookingDataById)
	// router.GET("/api/booking", route.GetBookings)
	// router.POST("/api/booking", route.AddNewBooking)
	// router.DELETE("/api/booking/:id", route.DeleteBookingByID)
	// router.PUT("/api/booking/:id", route.UpdateBookingDataById)

	// router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (c *Router) SetDataBase(db *db.DataBase) {
	c.db = db
}

func (c *Router) Run(port string) {
	c.router.Run(":" + port)
}

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
	createdAt, errc := time.Parse(time.RFC3339, ctx.PostForm("created_at"))
	if errc != nil {
		dv.SendMessage(ctx, http.StatusBadRequest, fmt.Sprintf("%s", errc))
		return
	}
	updatedAt, erru := time.Parse(time.RFC3339, ctx.PostForm("updated_at"))
	if errH != nil || errc != nil || erru != nil {
		dv.SendMessage(ctx, http.StatusBadRequest, fmt.Sprintf("%s", erru))
		return
	}
	user.Password = password
	user.CreatedAt = createdAt
	user.UpdatedAt = updatedAt
	idData, httpCode, errA := c.db.UserController.Add(user)
	if errA != nil {
		dv.SendMessage(ctx, http.StatusBadRequest, fmt.Sprintf("%s", errA))
		return
	}
	var res AddUserResponse
	fmt.Print(idData)
	res.Id = idData.InsertUser[0].Id
	data, errM := json.Marshal(res)
	if errM != nil {
		dv.SendMessage(ctx, http.StatusInternalServerError, fmt.Sprintf("%s", errM))
		return
	}
	ctx.Data(httpCode, "application/json", data)
}
