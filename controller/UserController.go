package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Adarsh-Kmt/EndServer/service"
	types "github.com/Adarsh-Kmt/EndServer/types"
	util "github.com/Adarsh-Kmt/EndServer/util"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type UserController struct {
	messageService service.MessageService
	userService    service.UserService
}

type Message struct {
	Body       string `json:"body"`
	SenderId   string `json:"senderId"`
	ReceiverId string `json:"receiverId"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewUserControllerInstance(MessageService service.MessageService, UserService service.UserService) *UserController {

	return &UserController{messageService: MessageService, userService: UserService}
}

func (uc *UserController) InitializeRouterEndpoints(router *mux.Router) *mux.Router {

	router.HandleFunc("/sendMessage", util.MakeJWTAuthHttpHandlerFunc(util.MakeHttpHandlerFunc(uc.SendMessage)))
	router.HandleFunc("/register", util.MakeHttpHandlerFunc(uc.RegisterUser))
	router.HandleFunc("/login", util.MakeHttpHandlerFunc(uc.LoginUser))
	return router
}

func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) *util.HttpError {

	rq := new(types.UserRegisterRequest)

	json.NewDecoder(r.Body).Decode(rq)

	if errorMap := util.ValidateRegisterRequest(*rq); errorMap != nil {

		return &util.HttpError{Status: 422, Error: errorMap}
	}
	httpError := uc.userService.RegisterUser(rq)

	return httpError

}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) *util.HttpError {

	rq := new(types.UserLoginRequest)

	json.NewDecoder(r.Body).Decode(rq)

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
func (uc *UserController) SendMessage(w http.ResponseWriter, r *http.Request) *util.HttpError {

	jwtToken := r.Header.Get("Auth")

	senderUsername, err := util.GetUsernameFromJwtToken(jwtToken)

	if err != nil {

		return &util.HttpError{Error: "Internal Server Error.", Status: 500}
	}
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return &util.HttpError{Error: "error while switching protocols.", Status: 500}

	}

	err = uc.messageService.UserConnected(senderUsername, conn)

	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte("error while updating connection status of user in the distribution server."))
	}
	uc.messageService.SendMessage(senderUsername, conn)

	return nil
}
