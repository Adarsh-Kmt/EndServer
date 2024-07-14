package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/Adarsh-Kmt/EndServer/db/mysql_code_gen"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLDatabase struct {
	Client *mysql_code_gen.Queries
}

func NewMySQLDatabaseInstance() (*MySQLDatabase, error) {

	conn, err := sql.Open("mysql", "root:52abx-32qj@(mysqlDB:3306)/chatDB")

	if err != nil {
		log.Println("error while establishing connection with mysql database.")
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		log.Println("pong error.")
		return nil, err
	}

	client := mysql_code_gen.New(conn)

	client.CreateChatDBDatabase(context.Background())

	return &MySQLDatabase{Client: client}, nil
}
