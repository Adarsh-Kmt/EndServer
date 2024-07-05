package controller

import (
	"net/http"

	"github.com/Adarsh-Kmt/EndServer/service"
	util "github.com/Adarsh-Kmt/EndServer/util"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type UserController struct {
	messageService service.MessageService
	// distributorNodeGRPCClient generatedCode.DistributionServerMessageServiceClient
	// es                        grpc_server.EndServer
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

func NewUserControllerInstance(MessageService service.MessageService) *UserController {

	return &UserController{messageService: MessageService}
}

func (uc *UserController) InitializeRouterEndpoints(router *mux.Router) *mux.Router {

	router.HandleFunc("/sendMessage/{userId}", util.MakeHttpHandlerFunc(uc.SendMessage))

	return router
}

func (uc *UserController) SendMessage(w http.ResponseWriter, r *http.Request) *util.HttpError {

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return &util.HttpError{Error: "error while switching protocols", Status: 500}

	}

	vars := mux.Vars(r)

	userId := vars["userId"]

	err = uc.messageService.UserConnected(userId, conn)

	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte("error while updating connection status of user in the distribution server."))
	}
	uc.messageService.SendMessage(conn)

	return nil
}
