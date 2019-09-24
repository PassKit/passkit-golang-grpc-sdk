package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type grpcFunction func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error

func pkHeaders(key string) (mdName string, ok bool) {
	if strings.HasPrefix(key, "pk-") {
		return key[3:], true
	}
	return "", false
}

// newGateway returns a new gateway server which translates HTTP into gRPC.
func newGateway(f grpcFunction, ctx context.Context,
	endpoint string, grpcOpts []grpc.DialOption, opts []runtime.ServeMuxOption) (http.Handler, error) {

	runtime.HTTPError = CustomHTTPError

	customMar := jsonpb.Marshaler{
		EmitDefaults: true,
	}
	mar := runtime.JSONPb(customMar)
	opts = append(opts, runtime.WithMarshalerOption("*/*", &mar))
	opts = append(opts, runtime.WithOutgoingHeaderMatcher(pkHeaders))
	opts = append(opts, runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true,
		EmitDefaults: true}))

	mux := runtime.NewServeMux(opts...)
	if err := f(ctx, mux, endpoint, grpcOpts); err != nil {
		return nil, err
	}
	return mux, nil
}

func dial(ctx context.Context, network, addr string, opts grpc.DialOption) (*grpc.ClientConn, error) {
	switch network {
	case "tcp":
		return dialTCP(ctx, addr, opts)
	case "unix":
		return dialUnix(ctx, addr, opts)
	default:
		return nil, fmt.Errorf("unsupported network type %q", network)
	}
}

// dialTCP creates a client connection via TCP.
// "addr" must be a valid TCP address with a port number.
func dialTCP(ctx context.Context, addr string, opt grpc.DialOption) (*grpc.ClientConn, error) {
	return grpc.DialContext(ctx, addr, opt)
}

// dialUnix creates a client connection via a unix domain socket.
// "addr" must be a valid path to the socket.
func dialUnix(ctx context.Context, addr string, opt grpc.DialOption) (*grpc.ClientConn, error) {
	d := func(addr string, timeout time.Duration) (net.Conn, error) {
		return net.DialTimeout("unix", addr, timeout)
	}
	return grpc.DialContext(ctx, addr, opt, grpc.WithDialer(d))
}

type errorBody struct {
	Err string `json:"error,omitempty"`
}

func CustomHTTPError(_ context.Context, _ *runtime.ServeMux, marshaller runtime.Marshaler, w http.ResponseWriter,
	_ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`

	s, ok := status.FromError(err)
	if ok {
		w.Header().Set("Content-type", marshaller.ContentType())
		w.WriteHeader(runtime.HTTPStatusFromCode(s.Code()))
		if jErr := json.NewEncoder(w).Encode(errorBody{Err: s.Message()}); jErr != nil {
			_, _ = w.Write([]byte(fallback))
		}
	}
}
