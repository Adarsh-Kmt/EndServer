// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: queries.sql

package mysql_code_gen

import (
	"context"
)

const addUser = `-- name: AddUser :exec
INSERT INTO user_table(username, password) VALUES(?, ?)
`

type AddUserParams struct {
	Username string
	Password string
}

func (q *Queries) AddUser(ctx context.Context, arg AddUserParams) error {
	_, err := q.db.ExecContext(ctx, addUser, arg.Username, arg.Password)
	return err
}

const createChatDBDatabase = `-- name: CreateChatDBDatabase :exec
CREATE DATABASE chatDB
`

func (q *Queries) CreateChatDBDatabase(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createChatDBDatabase)
	return err
}

const getUserCredentials = `-- name: GetUserCredentials :one
SELECT password FROM user_table WHERE username = ?
`

func (q *Queries) GetUserCredentials(ctx context.Context, username string) (string, error) {
	row := q.db.QueryRowContext(ctx, getUserCredentials, username)
	var password string
	err := row.Scan(&password)
	return password, err
}

const userExists = `-- name: UserExists :one
SELECT EXISTS(SELECT 1 FROM user_table WHERE username = ?)
`

func (q *Queries) UserExists(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRowContext(ctx, userExists, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}