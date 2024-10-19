package grpc_server

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
	"os/exec"
	"sync"

	generatedCode "github.com/Adarsh-Kmt/EndServer/generatedCode"
	"github.com/gorilla/websocket"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
)

type EndServer struct {
	ContainerName string
	generatedCode.UnimplementedEndServerMessageServiceServer
	ActiveConn   map[string]*websocket.Conn
	ConnMutexMap map[string]*sync.Mutex // used to synchronize concurrent writes to the same websocket connection.
}

func NewEndServerInstance() *EndServer {

	endServerContainerName := os.Getenv("CONTAINER_NAME")
	return &EndServer{
		ContainerName: endServerContainerName,
		ActiveConn:    make(map[string]*websocket.Conn),
		ConnMutexMap:  make(map[string]*sync.Mutex),
	}

}

func NewGRPCEndServerInstance(endServerInstance *EndServer) *grpc.Server {

	GenerateTLSCertificate()
	tlsConfig := GenerateTLSConfigObjectForEndServer()
	GRPCEndServer := grpc.NewServer(grpc.Creds(credentials.NewTLS(tlsConfig)), grpc.UnaryInterceptor(MiddlewareHandler))

	generatedCode.RegisterEndServerMessageServiceServer(GRPCEndServer, endServerInstance)

	return GRPCEndServer
}

func MiddlewareHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// you can write your own code here to check client tls certificate
	if p, ok := peer.FromContext(ctx); ok {
		if mtls, ok := p.AuthInfo.(credentials.TLSInfo); ok {
			for _, item := range mtls.State.PeerCertificates {
				log.Println("client certificate subject:", item.Subject)
			}
		}
	}
	return handler(ctx, req)
}

func GenerateTLSCertificate() {

	cmd := exec.Command("/bin/sh", "/prod/csrGenerationScript.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Error running shell script: %v", err)
	}

	log.Println("Shell script executed successfully.")

}
func GenerateTLSConfigObjectForEndServer() *tls.Config {

	EndServerCert, err := tls.LoadX509KeyPair("/prod/EndServer.pem", "/prod/EndServer-key.pem")

	if err != nil {

		log.Fatal("error while loading key pair of End Server: " + err.Error())
	}

	RootCA := x509.NewCertPool()

	caBytes, err := os.ReadFile("/prod/root.pem")

	if len(caBytes) == 0 {
		log.Fatal("signed root certificate was not read")
	}
	if err != nil {

		log.Fatal("error while reading root certificate from file in End Server: " + err.Error())
	}

	if ok := RootCA.AppendCertsFromPEM(caBytes); !ok {

		log.Fatal("failed to load certificate of root CA into certificate poll in End Server.")
	} else {

		log.Println("successfully read certificate of root certificate authority.")
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{EndServerCert},
		ClientCAs:    RootCA,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}

	if err != nil {
		log.Fatal("error while loading TLS certificate of End Server.")
	}

	return tlsConfig
}
func NewDistributionServerClientInstance() generatedCode.DistributionServerMessageServiceClient {

	tlsConfig := GenerateTLSConfigObjectForEndServer()

	DNGRPCConn, err := grpc.NewClient("ds:9000", grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	if err != nil {
		log.Fatal("error")
	}
	if DNGRPCConn != nil {
		log.Println("connection initialized")
	}
	DNGRPCClient := generatedCode.NewDistributionServerMessageServiceClient(DNGRPCConn)

	return DNGRPCClient
}

func (es *EndServer) ReceiveMessage(ctx context.Context, message *generatedCode.EndServerMessage) (*generatedCode.EndServerResponse, error) {

	ReceiverWebsocketConnection := es.ActiveConn[message.ReceiverUsername]

	response := &generatedCode.EndServerResponse{}
	if ReceiverWebsocketConnection == nil {
		log.Println("User " + message.ReceiverUsername + " is not online right now.")
		response.Status = 404
		return response, nil
	} else {
		log.Printf("end server received message %s for user %s", message.Body, message.ReceiverUsername)

		mutex := es.ConnMutexMap[message.ReceiverUsername]

		mutex.Lock()
		ReceiverWebsocketConnection.WriteMessage(websocket.TextMessage, []byte(message.Body))
		mutex.Unlock()

		response.Status = 200
		return response, nil
	}
}
