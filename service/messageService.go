package service

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Adarsh-Kmt/EndServer/generatedCode"
	"github.com/Adarsh-Kmt/EndServer/grpc_server"

	types "github.com/Adarsh-Kmt/EndServer/types"
	"github.com/gorilla/websocket"
)

type MessageService interface {
	SendMessage(conn *websocket.Conn)
	UserConnected(userId string, conn *websocket.Conn) error
}

type MessageServiceImpl struct {
	DistributionServerClient generatedCode.DistributionServerMessageServiceClient
	EndServer                grpc_server.EndServer
}

func NewMessageServiceImplInstance(distributionServerClient generatedCode.DistributionServerMessageServiceClient, endServer grpc_server.EndServer) *MessageServiceImpl {

	return &MessageServiceImpl{DistributionServerClient: distributionServerClient, EndServer: endServer}
}
func (ms *MessageServiceImpl) SendMessage(conn *websocket.Conn) {

	for {

		_, message, err := conn.ReadMessage()

		if err != nil {
			log.Println(err.Error())
			log.Fatal("error while reading message")
			//return &HttpError{Error: "error while reading message", status: 500}
		}

		var localMessage types.MessageRequest
		var grpcMessage generatedCode.DistributionServerMessage

		if err := json.Unmarshal(message, &localMessage); err != nil {
			log.Fatal(err.Error())

			//return &HttpError{Error: "error unmarshalling message", status: 500}

		}

		log.Println(localMessage)
		grpcMessage = generatedCode.DistributionServerMessage{ReceiverId: localMessage.ReceiverUserId, SenderId: localMessage.SenderUserId, Body: localMessage.Body}
		log.Println(grpcMessage.Body + " message received")

		response, err := ms.DistributionServerClient.SendMessage(context.Background(), &grpcMessage)

		if err != nil {
			log.Fatal("grpcError")

			//return &HttpError{Error: "grpc error", status: 500}
		}

		if response.ResponseStatus != 200 {
			log.Fatal("grpcError")

			//return &HttpError{Error: "grpc error", status: 500}
		}

	}

}

func (ms *MessageServiceImpl) UserConnected(userId string, conn *websocket.Conn) error {

	// if _, exists := uc.es.activeConn[grpcMessage.SenderId]; !exists {

	// 	connectionMessage := DistributionServerConnectionRequest{UserId: grpcMessage.SenderId, EndServerAddress: "es:3000"}

	// 	connectionResponse, _ := uc.distributorNodeGRPCClient.UserConnected(context.Background(), &connectionMessage)

	// 	//log.Println(connectionResponse)
	// 	if connectionResponse == nil {
	// 		log.Fatal("connection response is nil.")
	// 	}
	// 	if connectionResponse.ResponseStatus == 500 {
	// 		log.Fatal("error while informing distributorNode.")

	// 		//return &HttpError{Error: "error while informing distributor Node.", status: 500}
	// 	}

	// 	uc.es.activeConn[grpcMessage.SenderId] = conn
	// }

	//log.Println("user connected")

	userConnectionRequest := &generatedCode.DistributionServerConnectionRequest{UserId: userId, EndServerAddress: ms.EndServer.ContainerName + ":3000"}
	userConnectionResponse, err := ms.DistributionServerClient.UserConnected(context.Background(), userConnectionRequest)

	if userConnectionResponse == nil {
		log.Fatal("did not receive user connection response from distribution server for user: " + userId)
	}

	if err != nil || userConnectionResponse.ResponseStatus == 500 {
		return err
	}

	ms.EndServer.ActiveConn[userId] = conn

	log.Println("distribution server has successfully logged user: " + userId + "'s connection status.")
	return nil
}
