package service

import (
	"log"

	"github.com/Adarsh-Kmt/EndServer/repository"
	"github.com/Adarsh-Kmt/EndServer/types"
	"github.com/Adarsh-Kmt/EndServer/util"
)

type UserService interface {
	RegisterUser(*types.UserRegisterRequest) *util.HttpError
	LoginUser(*types.UserLoginRequest) (string, *util.HttpError)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserServiceImplInstance(UserRepository repository.UserRepository) *UserServiceImpl {

	return &UserServiceImpl{userRepository: UserRepository}
}
func (usi *UserServiceImpl) RegisterUser(urr *types.UserRegisterRequest) *util.HttpError {

	userExists, err := usi.userRepository.UserExists(urr.Username)

	if err != nil {
		log.Println("error : " + err.Error() + " occured in user service.")
		return &util.HttpError{Status: 500, Error: "internal server error."}
	}
	if userExists {

		return &util.HttpError{Status: 409, Error: "user with id " + urr.Username + " already exists."}
	}

	err = usi.userRepository.SaveUser(urr)

	if err != nil {
		return &util.HttpError{Status: 500, Error: "internal server error."}
	}

	return nil
}

func (usi *UserServiceImpl) LoginUser(ulr *types.UserLoginRequest) (string, *util.HttpError) {

	password, err := usi.userRepository.GetUserCredentials(ulr.Username)

	if err != nil {
		log.Println("error : " + err.Error() + " occured in user service.")
		return "", &util.HttpError{Status: 500, Error: "internal server error."}
	}
	if password != ulr.Password {

		return "", &util.HttpError{Status: 401, Error: "incorrect username/password"}
	}

	jwtToken, httpError := util.GenerateJwtToken(ulr.Username)

	if httpError != nil {
		return "", httpError
	}

	return jwtToken, nil

}
