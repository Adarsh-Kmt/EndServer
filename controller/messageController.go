package controller

import (
	"net/http"

	"github.com/Adarsh-Kmt/EndServer/service"
	util "github.com/Adarsh-Kmt/EndServer/util"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type MessageController struct {
	messageService service.MessageService
}

func NewMessageControllerInstance(MessageService service.MessageService) *MessageController {

	return &MessageController{messageService: MessageService}

}
func (mc *MessageController) InitializeRouterEndpoints(router *mux.Router) *mux.Router {

	router.HandleFunc("/sendMessage", util.MakeJWTAuthHttpHandlerFunc(util.MakeHttpHandlerFunc(mc.SendMessage)))

	return router
}
func (mc *MessageController) SendMessage(w http.ResponseWriter, r *http.Request) *util.HttpError {

	jwtToken := r.Header.Get("Auth")

	senderUsername, err := util.GetUsernameFromJwtToken(jwtToken)

	if err != nil {

		return &util.HttpError{Error: "Internal Server Error.", Status: 500}
	}
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return &util.HttpError{Error: "error while switching protocols.", Status: 500}

	}

	err = mc.messageService.UserConnected(senderUsername, conn)

	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte("error while updating connection status of user in the distribution server."))
	}
	mc.messageService.SendMessage(senderUsername, conn)

	return nil
}
