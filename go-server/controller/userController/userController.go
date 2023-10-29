package userController

import (
	"context"
	"fmt"
	"time"

	"github.com/edgedb/edgedb-go"
	"github.com/machinebox/graphql"
	"github.com/subliker/ToUni/go-server/model"
)

type User struct {
	model.User
	client *graphql.Client
}

func (c *User) SetClient(client *graphql.Client) {
	fmt.Print(c.client)
	c.client = client
}

type QueryUserResponse struct {
	User []struct {
		model.User
	}
}

func (c *User) GetAll() (QueryUserResponse, int, error) {
	ctx := context.Background()
	query := `
	{
		User {
			id
			username
			password
			created_at
			updated_at
			role
		}
	}
	`
	req := graphql.NewRequest(query)
	var res QueryUserResponse
	return res, 200, c.client.Run(ctx, req, &res)
}

func (c *User) GetOneById(id edgedb.UUID) (QueryUserResponse, int, error) {
	ctx := context.Background()
	query := fmt.Sprintf(`
	{
		User (filter:{id: {eq: "%s"}}) {
			id
			username
			password
			created_at
			updated_at
			role
		}
	}
	`, id)
	req := graphql.NewRequest(query)
	var res QueryUserResponse
	return res, 200, c.client.Run(ctx, req, &res)
}

type MutationInsertUserResponse struct {
	InsertUser []struct {
		Id string `json:"id"`
	} `json:"insert_User"`
}

func (c *User) Add(user model.User) (MutationInsertUserResponse, int, error) {
	ctx := context.Background()
	mutation := fmt.Sprintf(`
	mutation{
		insert_User(
			data: {username: "%s", password: "%s", created_at: "%s", updated_at: "%s", role: "%s"}
		) {
			id
		}
	}
	`, user.Username, user.Password, user.CreatedAt.Format(time.RFC3339), user.UpdatedAt.Format(time.RFC3339), user.Role)
	req := graphql.NewRequest(mutation)
	var res MutationInsertUserResponse
	return res, 200, c.client.Run(ctx, req, &res)
}

func (c *User) ExistUsername(username string) (bool, error) {
	ctx := context.Background()
	query := fmt.Sprintf(`
	{
		User (filter:{username: {eq: "%s"}}) {
			username
		}
	}
	`, username)
	req := graphql.NewRequest(query)
	var res QueryUserResponse
	err := c.client.Run(ctx, req, &res)
	if err != nil {
		return false, err
	}
	if len(res.User) != 0 {
		return true, nil
	}
	return false, nil
}

func (c *User) GetHashedPasswordByUsername(username string) (string, edgedb.UUID, error) {
	ctx := context.Background()
	query := fmt.Sprintf(`
	{
		User (filter:{username: {eq: "%s"}}) {
			id
			password
		}
	}
	`, username)
	req := graphql.NewRequest(query)
	var res QueryUserResponse
	err := c.client.Run(ctx, req, &res)
	if err != nil {
		return "", [16]byte{}, err
	}
	return res.User[0].Password, res.User[0].Id, nil
}
