package grpc_server

import (
	"log"
	"os"

	generatedCode "github.com/Adarsh-Kmt/EndServer/generatedCode"
	"github.com/gorilla/websocket"
	"golang.org/x/net/context"
)

type EndServer struct {
	ContainerName string
	generatedCode.UnimplementedEndServerMessageServiceServer
	ActiveConn map[string]*websocket.Conn
	//DistributionServerClient DN_GeneratedCode.MessageServiceClient
}

func NewEndServerInstance() *EndServer {

	endServerContainerName := os.Getenv("CONTAINER_NAME")
	return &EndServer{ContainerName: endServerContainerName, ActiveConn: make(map[string]*websocket.Conn)}

}
func (ed *EndServer) ReceiveMessage(ctx context.Context, message *generatedCode.EndServerMessage) (*generatedCode.EndServerResponse, error) {

	ReceiverWebsocketConnection := ed.ActiveConn[message.ReceiverId]
	if ReceiverWebsocketConnection == nil {
		log.Fatal("no connection exists")
	}
	ReceiverWebsocketConnection.WriteMessage(websocket.TextMessage, []byte(message.Body))

	return &generatedCode.EndServerResponse{Status: 200}, nil
}
