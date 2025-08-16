package grpcclient

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"github.com/harry14-glitch/hsmicroutils/genproto/mail"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type MailClient struct {
	conn   *grpc.ClientConn
	client mail.MailServiceClient
}

var mailClient *MailClient

func GetMailClient() *MailClient {
	return mailClient
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	certFileName := os.Getenv("CA_CERT_FILE")
	if certFileName == "" {
		certFileName = "ca-cert.pem" // Default path if
	}
	caCert, err := os.ReadFile("/app/cert/" + certFileName)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}

func StartMailClient() {

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	// create client connection
	conn, err := grpc.Dial(
		"mailer-srv:50051",
		grpc.WithTransportCredentials(tlsCredentials),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := mail.NewMailServiceClient(conn)

	mailClient = &MailClient{
		conn:   conn,
		client: client,
	}
}

func StopMailClient() {
	orgClient.conn.Close()
}
