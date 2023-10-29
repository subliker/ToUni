package router

import (
	"github.com/gin-contrib/cors"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/subliker/ToUni/go-server/db"
	"github.com/subliker/ToUni/go-server/middleware"
	_ "golang.org/x/crypto/bcrypt"
)

type Router struct {
	router      *gin.Engine
	adminRouter *gin.RouterGroup
	db          *db.DataBase
	user        User
	lesson      Lesson
	client      Client
}

type User struct {
	db *db.DataBase
}

type Lesson struct {
	db *db.DataBase
}

type Client struct {
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
	c.client.db = c.db

	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://google.com"}
	config.AllowAllOrigins = true

	c.router.Use(cors.New(config))

	// docs.SwaggerInfo.BasePath = "/api"
	c.adminRouter = c.router.Group("/")

	c.adminRouter.Use(middleware.CheckRole)
	c.adminRouter.GET("/api/user/:id", c.user.GetOneById)
	c.adminRouter.GET("/api/user", c.user.GetAll)
	c.adminRouter.POST("/api/user", c.user.Add)

	c.router.POST("/api/signup", c.client.SignUp)
	c.router.POST("/api/signin", c.client.SignIn)
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
