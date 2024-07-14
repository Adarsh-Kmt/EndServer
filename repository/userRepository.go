package repository

import (
	"context"

	"github.com/Adarsh-Kmt/EndServer/db"
	"github.com/Adarsh-Kmt/EndServer/db/mysql_code_gen"
	"github.com/Adarsh-Kmt/EndServer/types"
	//"database/sql"
)

// docker exec -it mysqlDB mysql -u root -p
type UserRepository interface {
	SaveUser(*types.UserRegisterRequest) error
	GetUserCredentials(string) (string, error)
	UserExists(string) (bool, error)
}

type UserRepositoryImpl struct {
	MySQLDatabase *db.MySQLDatabase
}

func NewUserRepositoryImplInstance(db *db.MySQLDatabase) *UserRepositoryImpl {

	return &UserRepositoryImpl{MySQLDatabase: db}
}

func (uri *UserRepositoryImpl) SaveUser(urr *types.UserRegisterRequest) error {

	err := uri.MySQLDatabase.Client.AddUser(context.Background(), mysql_code_gen.AddUserParams{
		Username: urr.UserId,
		Password: urr.Password,
	})
	return err
}

func (uri *UserRepositoryImpl) GetUserCredentials(username string) (string, error) {

	password, err := uri.MySQLDatabase.Client.GetUserCredentials(context.Background(), username)

	if err != nil {
		return "", err
	}
	return password, nil
}

func (uri *UserRepositoryImpl) UserExists(username string) (bool, error) {

	userExists, err := uri.MySQLDatabase.Client.UserExists(context.Background(), username)

	return userExists, err
}
