package model

import (
	"time"

	"github.com/edgedb/edgedb-go"
)

type User struct {
	Id        edgedb.UUID `edgedb:"id"`
	Username  string      `edgedb:"username"`
	Password  string      `edgedb:"password"`
	CreatedAt time.Time   `edgedb:"created_at"`
	UpdatedAt time.Time   `edgedb:"updated_at"`
}
