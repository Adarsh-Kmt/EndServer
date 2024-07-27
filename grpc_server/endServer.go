package grpc_server

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
	"os/exec"

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
	ActiveConn map[string]*websocket.Conn
	//DistributionServerClient DN_GeneratedCode.MessageServiceClient
}

func NewEndServerInstance() *EndServer {

	endServerContainerName := os.Getenv("CONTAINER_NAME")
	return &EndServer{ContainerName: endServerContainerName, ActiveConn: make(map[string]*websocket.Conn)}

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

	cmd := exec.Command("/bin/sh", "/app/csrGenerationScript.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Error running shell script: %v", err)
	}

	log.Println("Shell script executed successfully.")

}
func GenerateTLSConfigObjectForEndServer() *tls.Config {

	EndServerCert, err := tls.LoadX509KeyPair("/app/EndServer.pem", "/app/EndServer-key.pem")

	if err != nil {

		log.Fatal("error while loading key pair of End Server: " + err.Error())
	}

	RootCA := x509.NewCertPool()

	caBytes, err := os.ReadFile("/app/root.pem")

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

func (ed *EndServer) ReceiveMessage(ctx context.Context, message *generatedCode.EndServerMessage) (*generatedCode.EndServerResponse, error) {

	ReceiverWebsocketConnection := ed.ActiveConn[message.ReceiverUsername]
	if ReceiverWebsocketConnection == nil {
		log.Fatal("no connection exists")
	}
	ReceiverWebsocketConnection.WriteMessage(websocket.TextMessage, []byte(message.Body))

	return &generatedCode.EndServerResponse{Status: 200}, nil
}
