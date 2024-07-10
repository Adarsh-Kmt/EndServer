package main

import (
	"log"
	"net"
	"net/http"

	controller "github.com/Adarsh-Kmt/EndServer/controller"

	grpc_server "github.com/Adarsh-Kmt/EndServer/grpc_server"
	service "github.com/Adarsh-Kmt/EndServer/service"
	"github.com/gorilla/mux"
)

// docker run -it --name es1 --network chat-network -p 8081:8080 -e CONTAINER_NAME=es1 end_server

func main() {

	endServerInstance := grpc_server.NewEndServerInstance()
	ENGRPCServer := grpc_server.NewGRPCEndServerInstance(endServerInstance)

	if ENGRPCServer == nil {
		log.Fatal("grpc end server not initialized")
	} else {
		log.Println("grpc end server initialized")
	}

	go func() {

		ENLis, err := net.Listen("tcp", ":3000")

		if err != nil {
			log.Fatal("error")
		}

		if err := ENGRPCServer.Serve(ENLis); err != nil {

			log.Fatal("error man.")
		}
	}()

	DNGRPCClient := grpc_server.NewDistributionServerClientInstance()

	muxRouter := mux.NewRouter()
	ms := service.NewMessageServiceImplInstance(DNGRPCClient, *endServerInstance)

	if ms == nil {
		log.Fatal("ms not initialized")
	} else {
		log.Println("ms initialized")
	}

	uc := controller.NewUserControllerInstance(ms)

	if uc == nil {
		log.Fatal("uc not initialized")
	} else {
		log.Println("uc initialized")
	}

	muxRouter = uc.InitializeRouterEndpoints(muxRouter)

	http.ListenAndServe(":8080", muxRouter)

}
