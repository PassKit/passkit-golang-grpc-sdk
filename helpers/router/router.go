package router

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	"stash.passkit.com/p/common-tools/config"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ServerConfig struct {
	GRPCHost               string
	RestHost               string
	GRPCPort               string
	RestPort               string
	GRPCCertFile           string
	GRPCKeyFile            string
	InternalClientCertFile string
	InternalClientKeyFile  string
	CAChainFile            string
}

// Endpoint describes a gRPC endpoint
type Endpoint struct {
	Network string
	Address string
}

var cors []string

// Options is a set of options to be passed to Run
type ServerOptions struct {
	// Addr is the address for the REST server to listen
	RESTAddress string
	// GRPCServer defines an endpoint of a gRPC service
	GRPCServer Endpoint
	// SwaggerDir is a path to a directory from which the server serves swagger specs.
	SwaggerDir string
	// CORS authorized domains
	CORS []string
	// gRPC protoc generated gateway register handler function
	Function grpcFunction
	// Mux is a list of options to be passed to the grpc-gateway multiplexer
	Mux []runtime.ServeMuxOption
	// Server Certificate File
	GRPCServerCertFile string
	// Server Key FIle
	GRPCServerKeyFile string
	// Certificate File - If absent will connect insecurely; should be a chained file if not using mutual tls
	GRPCClientCertFile string
	// Private Key File - optional, but if present will connect with mutual TLS
	GRPCClientKeyFile string
	// Certificate Chain file - required only if GRPC key file is present
	GRPCChainFile string
	// If service has a root route, set to true
	HasRoot bool
}

func GetDefaultOptions() ServerOptions {

	var path string
	if config.IsLocal() {
		path = "../../assets"
	}
	return ServerOptions{
		RESTAddress: "0.0.0.0:8080",
		GRPCServer: Endpoint{
			Network: "tcp",
			Address: "0.0.0.0:9999",
		},
		GRPCServerCertFile: path + "/crypto/server.pem",
		GRPCServerKeyFile:  path + "/crypto/server-key.pem",
		GRPCClientCertFile: path + "/crypto/client.pem",
		GRPCClientKeyFile:  path + "/crypto/client-key.pem",
		GRPCChainFile:      path + "/crypto/ca-chain.pem",
	}
}

func ServeRESTWithCORSDefaults(f grpcFunction, mw []MiddlewareFunc, cors []string) error {

	o := GetDefaultOptions()
	o.Function = f
	o.CORS = cors

	// Start REST Server
	ctx := context.Background()
	return Run(ctx, o, mw)
}

func ServeRESTWithCORSOptions(f grpcFunction, mw []MiddlewareFunc, cors []string, opts ServerOptions) error {

	o := opts
	o.Function = f
	o.CORS = cors

	// Start REST Server
	ctx := context.Background()
	return Run(ctx, o, mw)
}

func ServeRESTDefaults(f grpcFunction, mw []MiddlewareFunc) error {

	o := GetDefaultOptions()
	o.Function = f

	// Start REST Server
	ctx := context.Background()
	return Run(ctx, o, mw)
}

// Run starts a HTTP server and blocks while running if successful.
// The server will be shutdown when "ctx" is canceled.
func Run(ctx context.Context, opts ServerOptions, middleware []MiddlewareFunc) error {

	// If gRPC server is set to listen globally (default), assume running on same server and replace with localhost
	opts.GRPCServer.Address = strings.Replace(opts.GRPCServer.Address, "0.0.0.0", "127.0.0.1", 1)
	GRPCHost, _, err := net.SplitHostPort(opts.GRPCServer.Address)
	if err != nil {
		return err
	}

	var handler http.Handler

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var dialOptions []grpc.DialOption
	switch {
	case opts.GRPCClientCertFile == "":
		dialOptions = []grpc.DialOption{grpc.WithInsecure()}
	case opts.GRPCClientKeyFile == "":
		creds, err := credentials.NewClientTLSFromFile(opts.GRPCClientCertFile, "")
		if err != nil {
			return err
		}
		dialOptions = []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	default:
		// Load the client certificates from disk
		certificate, err := tls.LoadX509KeyPair(opts.GRPCClientCertFile, opts.GRPCClientKeyFile)
		if err != nil {
			return fmt.Errorf("could not load client key pair: %s", err)
		}

		// Create a certificate pool from the certificate authority
		certPool := x509.NewCertPool()
		ca, err := ioutil.ReadFile(opts.GRPCChainFile)
		if err != nil {
			return fmt.Errorf("could not read ca certificate: %s", err)
		}

		// Append the certificates from the CA
		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			return errors.New("failed to append ca certs")
		}

		creds := credentials.NewTLS(&tls.Config{
			ServerName:   GRPCHost,
			Certificates: []tls.Certificate{certificate},
			RootCAs:      certPool,
		})
		dialOptions = []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	}

	// attempt to connect to the GPRC server
	conn, err := dial(ctx, opts.GRPCServer.Network, opts.GRPCServer.Address, dialOptions[0])
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		if err := conn.Close(); err != nil {
			log.Printf("Failed to close a client connection to the gRPC server: %v", err)
		}
	}()

	mux := http.NewServeMux()
	if opts.SwaggerDir != "" {
		mux.HandleFunc("/swagger/", serveSwagger(opts.SwaggerDir))
	}

	mux.HandleFunc("/healthz", healthCheck(conn))

	gw, err := newGateway(opts.Function, ctx, opts.GRPCServer.Address, dialOptions, opts.Mux)
	if err != nil {
		return err
	}
	mux.Handle("/", gw)

	if len(opts.CORS) != 0 {
		cors = opts.CORS
		middleware = append(middleware, allowCORS)
	}

	middleware = append(middleware, passKitIoDefault)
	handler = middleware[0](mux)

	if !opts.HasRoot {
		middleware = append(middleware, mockRoot)
	}

	for i := 1; i < len(middleware); i++ {
		handler = middleware[i](handler)
	}

	s := &http.Server{
		Addr:    opts.RESTAddress,
		Handler: handler,
	}

	go func() {
		<-ctx.Done()
		log.Printf("Shutting down the http server")
		if err := s.Shutdown(context.Background()); err != nil {
			glog.Errorf("Failed to shutdown http server: %v", err)
		}
	}()

	log.Printf("Starting listening at %s", opts.RESTAddress)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Failed to listen and serve: %v", err)
		return err
	}
	return nil
}
