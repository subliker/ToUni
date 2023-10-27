package lessonController

import (
	"context"
	"time"

	"github.com/edgedb/edgedb-go"
	"github.com/machinebox/graphql"
)

type Lesson struct {
	Owner struct {
		Id edgedb.UUID
	} `edgedb:"owner"`
	// OwnerId     *edgedb.UUID `edgedb:"owner"`
	Name        string    `edgedb:"name"`
	CreatedAt   time.Time `edgedb:"created_at"`
	UpdatedAt   time.Time `edgedb:"updated_at"`
	Data        string    `edgedb:"data"`
	Description string    `edgedb:"description"`
	Files       []string  `edgedb:"files"`
	client      *graphql.Client
	ctx         context.Context
}

func (c *Lesson) SetClient(client *graphql.Client) {
	c.client = client
}

func (c *Lesson) SetCtx(ctx context.Context) {
	c.ctx = ctx
}

// func (c *Lesson) Add(lesson Lesson) error {
// 	var inserted struct{ id edgedb.UUID }
// 	return c.client.QuerySingle(c.ctx, `
// 	INSERT Lesson {
// 		owner := <User>$0,
// 		name := <str>$1,
// 		created_at := <datetime>$2,
// 		updated_at := <datetime>$3,
// 		data := <str>$4,
// 		description := <str>$5
// 	}
// 	`, &inserted, lesson.Owner.Id, lesson.Name, lesson.CreatedAt, lesson.UpdatedAt, lesson.Data, lesson.Description)
// }
