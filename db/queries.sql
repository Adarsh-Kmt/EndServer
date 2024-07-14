
-- name: CreateChatDBDatabase :exec
CREATE DATABASE chatDB;

-- name: GetUserCredentials :one
SELECT password FROM user_table WHERE username = ?;

-- name: AddUser :exec
INSERT INTO user_table(username, password) VALUES(?, ?);

-- name: UserExists :one
SELECT EXISTS(SELECT 1 FROM user_table WHERE username = ?);