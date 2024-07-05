package main

import (
	"log"
	"net"
	"net/http"

	controller "github.com/Adarsh-Kmt/EndServer/controller"
	generatedCode "github.com/Adarsh-Kmt/EndServer/generatedCode"
	grpc_server "github.com/Adarsh-Kmt/EndServer/grpc_server"
	service "github.com/Adarsh-Kmt/EndServer/service"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// docker run -it --name es1 --network chat-network -p 8081:8080 -e CONTAINER_NAME=es1 end_server

func main() {

	endServerInstance := grpc_server.NewEndServerInstance()

	if endServerInstance == nil {
		log.Fatal("end server instance not initialized")
	}

	log.Println("end server initialized")

	ENGRPCServer := grpc.NewServer()

	if ENGRPCServer == nil {
		log.Fatal("grpc end server not initialized")
	} else {
		log.Println("grpc end server initialized")
	}

	generatedCode.RegisterEndServerMessageServiceServer(ENGRPCServer, endServerInstance)

	go func() {

		ENLis, err := net.Listen("tcp", ":3000")

		if err != nil {
			log.Fatal("error")
		}

		if err := ENGRPCServer.Serve(ENLis); err != nil {

			log.Fatal("error man.")
		}
	}()

	DNGRPCConn, err := grpc.NewClient("ds:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error")
	}
	if DNGRPCConn != nil {
		log.Println("connection initialized")
	}
	DNGRPCClient := generatedCode.NewDistributionServerMessageServiceClient(DNGRPCConn)
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
