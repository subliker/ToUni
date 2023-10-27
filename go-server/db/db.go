package db

import (
	"fmt"
	"os"

	"github.com/machinebox/graphql"
	"github.com/subliker/ToUni/go-server/controller/lessonController"
	"github.com/subliker/ToUni/go-server/controller/userController"
)

type DataBase struct {
	Client           *graphql.Client
	UserController   *userController.User
	LessonController *lessonController.Lesson
}

func (c *DataBase) Init() {
	c.Client = graphql.NewClient(fmt.Sprintf("http://%s:%s/db/edgedb/graphql", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")))

	c.UserController = new(userController.User)
	c.UserController.SetClient(c.Client)
	c.LessonController = new(lessonController.Lesson)
	c.LessonController.SetClient(c.Client)
}

// func (c *DataBase) Init() {
// 	ctx := context.Background()
// 	var err error
// 	c.Client, err = edgedb.CreateClient(ctx, edgedb.Options{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	c.UserController = &userController.User{}
// 	c.UserController.SetClient(c.Client)
// 	c.UserController.SetCtx(ctx)

// 	c.LessonController = &lessonController.Lesson{}
// 	c.LessonController.SetClient(c.Client)
// 	c.LessonController.SetCtx(ctx)
// }
