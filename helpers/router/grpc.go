package router

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"passkit-golang-sdk/helpers/logging"
)

// NewTLSGRPCClient returns a secure gRPC client which will authenticate with the the provided certificate
func NewTLSGRPCClient(address string, certFile string) (*grpc.ClientConn, error) {

	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		return nil, fmt.Errorf("Could create gRPC TLS client credentials %v", err)
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, fmt.Errorf("Could not connect to gRPC Server: %v", err)
	}
	return conn, nil
}

// NewCertAuthTLSGRPCClient returns a secure gRPC client which will authenticate with the a certificate and key
func NewCertAuthTLSGRPCClient(address, certFile, keyFile, caFile string) (*grpc.ClientConn, error) {

	// Load the client certificates from disk
	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("could not load client key pair: %s", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("could not read ca certificate: %s", err)
	}

	// Append the certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, errors.New("failed to append ca certs")
	}

	// Get the host from the address
	host, _, err := net.SplitHostPort(address)
	if err != nil {
		return nil, err
	}

	creds := credentials.NewTLS(&tls.Config{
		ServerName:   host,
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, fmt.Errorf("Could not connect to gRPC Server: %v", err)
	}
	return conn, nil
}

// New GRPCClient returns and insecure gRPC client to the provided address
func NewGRPCClient(address string) (*grpc.ClientConn, error) {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("Could not connect to gRPC Server: %v", err)
	}
	return conn, nil
}

// GetContextWithJWT provides a helper function to provide a gRPC context with a JWT token
func GetContextWithJWT(jwt string) context.Context {
	ctx := context.Background()
	md := metadata.Pairs("authorization", jwt)
	return metadata.NewOutgoingContext(ctx, md)
}
