package service

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/Adarsh-Kmt/EndServer/generatedCode"
	"github.com/Adarsh-Kmt/EndServer/grpc_server"

	types "github.com/Adarsh-Kmt/EndServer/types"
	"github.com/gorilla/websocket"
)

type MessageService interface {
	SendMessage(senderUsername string, conn *websocket.Conn)
	UserConnected(userId string, conn *websocket.Conn) error
	UserDisconnected(username string, conn *websocket.Conn) error
}

type MessageServiceImpl struct {
	DistributionServerClient generatedCode.DistributionServerMessageServiceClient
	EndServer                grpc_server.EndServer
}

func NewMessageServiceImplInstance(distributionServerClient generatedCode.DistributionServerMessageServiceClient, endServer grpc_server.EndServer) *MessageServiceImpl {

	return &MessageServiceImpl{DistributionServerClient: distributionServerClient, EndServer: endServer}
}
func (ms *MessageServiceImpl) SendMessage(senderUsername string, senderConn *websocket.Conn) {

	for {

		_, message, err := senderConn.ReadMessage()

		if websocket.IsCloseError(err, websocket.CloseNormalClosure) {

			ms.UserDisconnected(senderUsername, senderConn)
			return
		}

		if err != nil {
			log.Println("error while reading message: " + err.Error())
			senderConn.WriteMessage(websocket.TextMessage, []byte("internal server error."))
		}

		var localMessage types.MessageRequest
		var grpcMessage generatedCode.DistributionServerMessage

		if err := json.Unmarshal(message, &localMessage); err != nil {
			log.Println("error occured while unmarshalling message: " + err.Error())
			senderConn.WriteMessage(websocket.TextMessage, []byte("internal server error."))
		}
		if localMessage.Body == "" {
			// blank messages are sent automatically. must fix.
			continue
		}
		log.Println(localMessage)
		grpcMessage = generatedCode.DistributionServerMessage{ReceiverUsername: localMessage.ReceiverUsername, SenderUsername: senderUsername, Body: localMessage.Body}
		log.Println("message: " + grpcMessage.Body + " received. To be sent from user: " + senderUsername + " to user: " + localMessage.ReceiverUsername)

		if _, exists := ms.EndServer.ActiveConn[localMessage.ReceiverUsername]; exists {

			log.Println("user: " + localMessage.ReceiverUsername + " connected to same container as user: " + senderUsername)
			ReceiverWebsocketConnection := ms.EndServer.ActiveConn[localMessage.ReceiverUsername]
			if ReceiverWebsocketConnection == nil {
				log.Println("websocket connection not found for user: " + localMessage.ReceiverUsername)
				senderConn.WriteMessage(websocket.TextMessage, []byte("internal server error."))
			} else {
				ReceiverWebsocketConnection.WriteMessage(websocket.TextMessage, []byte(localMessage.Body))
			}

		} else {

			response, err := ms.DistributionServerClient.SendMessage(context.Background(), &grpcMessage)

			if response == nil {
				log.Println("response from distribution server is nil.")
			} else {
				if response.ResponseStatus == 404 {
					senderConn.WriteMessage(websocket.TextMessage, []byte("user "+localMessage.ReceiverUsername+" is not online right now."))
				} else if response.ResponseStatus == 500 {
					if err != nil {
						log.Println("error while communicating to Distribution Server using gRPC: " + err.Error())
					} else {
						log.Println("error while communicating to Distribution Server using gRPC.")
					}
					senderConn.WriteMessage(websocket.TextMessage, []byte("internal server error."))
				}
			}
		}

	}

}

func (ms *MessageServiceImpl) UserConnected(username string, conn *websocket.Conn) error {

	userConnectionRequest := &generatedCode.DistributionServerConnectionRequest{Username: username, EndServerAddress: ms.EndServer.ContainerName + ":3000"}
	userConnectionResponse, err := ms.DistributionServerClient.UserConnected(context.Background(), userConnectionRequest)

	if userConnectionResponse == nil {
		log.Fatal("did not receive user connection response from distribution server for user: " + username)
	}

	if err != nil || userConnectionResponse.ResponseStatus == 500 {
		return err
	}

	ms.EndServer.ActiveConn[username] = conn
	ms.EndServer.ConnMutexMap[username] = &sync.Mutex{}
	log.Println("distribution server has successfully logged user: " + username + "'s connection status.")
	return nil
}

func (ms *MessageServiceImpl) UserDisconnected(username string, conn *websocket.Conn) error {

	delete(ms.EndServer.ActiveConn, username)
	log.Println("user " + username + " has disconnected.")
	endServerAddress := os.Getenv("CONTAINER_NAME") + ":3000"
	response, err := ms.DistributionServerClient.UserDisconnected(context.Background(), &generatedCode.DistributionServerConnectionRequest{Username: username, EndServerAddress: endServerAddress})

	if response.ResponseStatus != 200 || err != nil {
		log.Println("error while communicating to Distribution Server using gRPC: " + err.Error())
		return err
	}

	return nil
}
