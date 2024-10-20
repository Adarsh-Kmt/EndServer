package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Adarsh-Kmt/EndServer/service"
	types "github.com/Adarsh-Kmt/EndServer/types"
	util "github.com/Adarsh-Kmt/EndServer/util"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type UserController struct {
	userService service.UserService
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewUserControllerInstance(UserService service.UserService) *UserController {

	return &UserController{userService: UserService}
}

func (uc *UserController) InitializeRouterEndpoints(router *mux.Router) *mux.Router {

	router.HandleFunc("/register", util.MakeHttpHandlerFunc(uc.RegisterUser))
	router.HandleFunc("/login", util.MakeHttpHandlerFunc(uc.LoginUser))
	router.HandleFunc("/healthCheck", util.MakeHttpHandlerFunc(uc.HealthCheck))
	return router
}

func (uc *UserController) HealthCheck(w http.ResponseWriter, r *http.Request) *util.HttpError {

	log.Println("performing health check....")
	log.Println("sent response 200 OK")
	response := types.HealthCheckResponse{Status: 200}
	util.WriteJSON(w, 200, response)
	return nil
}
func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) *util.HttpError {

	rq := new(types.UserRegisterRequest)

	json.NewDecoder(r.Body).Decode(rq)
	log.Printf("received register request for user %s", rq.Username)
	if errorMap := util.ValidateRegisterRequest(*rq); errorMap != nil {

		return &util.HttpError{Status: 422, Error: errorMap}
	}
	httpError := uc.userService.RegisterUser(rq)

	if httpError != nil {
		return httpError
	}

	util.WriteJSON(w, 200, map[string]string{"successMessage": "you have registered successfully."})

	return nil

}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) *util.HttpError {

	rq := new(types.UserLoginRequest)

	err := json.NewDecoder(r.Body).Decode(rq)

	if err != nil {
		log.Printf("error while decoding : %s", err.Error())
	}
	log.Printf("received login request for user %s", rq.Username)
	if errorMap := util.ValidateLoginRequest(*rq); errorMap != nil {

		return &util.HttpError{Status: 422, Error: errorMap}
	}

	jwtToken, httpError := uc.userService.LoginUser(rq)

	if httpError != nil {
		return httpError
	}

	util.WriteJSON(w, 200, map[string]string{"jwtToken": jwtToken})

	return nil

}
