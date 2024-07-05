package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	generatedCode "github.com/Adarsh-Kmt/EndServer/generatedCode"
	grpc_server "github.com/Adarsh-Kmt/EndServer/grpc_server"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type UserController struct {
	distributorNodeGRPCClient generatedCode.DistributionServerMessageServiceClient
	es                        grpc_server.EndServer
}

type HttpError struct {
	Error  string
	status int
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

type HttpFunc func(w http.ResponseWriter, r *http.Request) *HttpError

func MakeHttpHandlerFunc(f HttpFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if HttpError := f(w, r); HttpError != nil {

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(HttpError.status)
			json.NewEncoder(w).Encode(HttpError.Error)
		}
	}
}

func NewUserControllerInstance(DNGRPCClient generatedCode.DistributionServerMessageServiceClient, es grpc_server.EndServer) *UserController {

	return &UserController{distributorNodeGRPCClient: DNGRPCClient, es: es}
}

func (uc *UserController) InitializeRouterEndpoints(router *mux.Router) *mux.Router {

	router.HandleFunc("/sendMessage/{userId}", MakeHttpHandlerFunc(uc.SendMessage))

	return router
}

func (uc *UserController) SendMessage(w http.ResponseWriter, r *http.Request) *HttpError {

	conn, err := upgrader.Upgrade(w, r, nil)

	vars := mux.Vars(r)

	userId := vars["userId"]

	connectionMessage := generatedCode.DistributionServerConnectionRequest{UserId: userId, EndServerAddress: uc.es.ContainerName + ":3000"}

	connectionResponse, _ := uc.distributorNodeGRPCClient.UserConnected(context.Background(), &connectionMessage)

	//log.Println(connectionResponse)
	if connectionResponse == nil {
		log.Fatal("connection response is nil.")
	}
	if connectionResponse.ResponseStatus == 500 {
		log.Fatal("error while informing distributorNode.")

		//return &HttpError{Error: "error while informing distributor Node.", status: 500}
	}

	uc.es.ActiveConn[userId] = conn

	log.Println("user connected")

	if err != nil {
		return &HttpError{Error: "error while switching protocols", status: 500}

	}

	for {

		_, message, err := conn.ReadMessage()

		if err != nil {
			log.Println(err.Error())
			log.Fatal("error while reading message")
			//return &HttpError{Error: "error while reading message", status: 500}
		}

		var localMessage Message
		var grpcMessage generatedCode.DistributionServerMessage

		if err := json.Unmarshal(message, &localMessage); err != nil {
			log.Fatal(err.Error())

			//return &HttpError{Error: "error unmarshalling message", status: 500}

		}

		log.Println(localMessage)
		grpcMessage = generatedCode.DistributionServerMessage{ReceiverId: localMessage.ReceiverId, SenderId: localMessage.SenderId, Body: localMessage.Body}
		log.Println(grpcMessage.Body + " message received")
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
		response, err := uc.distributorNodeGRPCClient.SendMessage(context.Background(), &grpcMessage)

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
