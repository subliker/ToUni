package model

import (
	"time"

	"github.com/edgedb/edgedb-go"
)

type User struct {
	Id        edgedb.UUID `edgedb:"id" json:"id"`
	Username  string      `edgedb:"username" json:"username"`
	Password  string      `edgedb:"password" json:"password"`
	CreatedAt time.Time   `edgedb:"created_at" json:"created_at"`
	UpdatedAt time.Time   `edgedb:"updated_at" json:"updated_at"`
	Role      string      `edgedb:"role" json:"role"`
}
