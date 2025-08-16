package grpcclient

import (
	"fmt"
	"log"
	"os"

	"github.com/adhiman2409/gomicroutils/genproto/auth"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	conn   *grpc.ClientConn
	client auth.AuthServiceClient
}

var grpcClient *GrpcClient

func GetAuthClient() *GrpcClient {
	return grpcClient
}

func StopGrpcClient() {
	grpcClient.conn.Close()
}

func StartAuthClient() {

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	// create client connection
	url := os.Getenv("AUTH_SERVICE_URL")
	if url == "" {
		url = "auth-srv:50051" // default URL if not set
	}
	conn, err := grpc.Dial(
		url,
		grpc.WithTransportCredentials(tlsCredentials),
	)
	if err != nil {
		fmt.Println("Failed to connect to auth service:", err, url)
		log.Fatal(err)
	}

	client := auth.NewAuthServiceClient(conn)

	grpcClient = &GrpcClient{
		conn:   conn,
		client: client,
	}
}
