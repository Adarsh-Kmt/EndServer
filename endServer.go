package main

import (
	"log"
	"os"

	"github.com/gorilla/websocket"
	"golang.org/x/net/context"
)

type EndServer struct {
	containerName string
	UnimplementedEndServerMessageServiceServer
	activeConn map[string]*websocket.Conn
	//DistributionServerClient DN_GeneratedCode.MessageServiceClient
}

func NewEndServerInstance() *EndServer {

	endServerContainerName := os.Getenv("CONTAINER_NAME")
	return &EndServer{containerName: endServerContainerName, activeConn: make(map[string]*websocket.Conn)}

}
func (ed *EndServer) ReceiveMessage(ctx context.Context, message *EndServerMessage) (*EndServerResponse, error) {

	ReceiverWebsocketConnection := ed.activeConn[message.ReceiverId]
	if ReceiverWebsocketConnection == nil {
		log.Fatal("no connection exists")
	}
	ReceiverWebsocketConnection.WriteMessage(websocket.TextMessage, []byte(message.Body))

	return &EndServerResponse{Status: 200}, nil
}
