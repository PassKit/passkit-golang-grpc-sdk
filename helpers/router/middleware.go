package router

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type MiddlewareFunc func(http.Handler) http.Handler

// MiddlewareFunc also implements the middleware interface.
func (mw MiddlewareFunc) Middleware(handler http.Handler) http.Handler {
	return mw(handler)
}

func passKitIoDefault(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pb := "PassKit.io"
		w.Header().Set("X-Powered-By", pb)
		h.ServeHTTP(w, r)
	})
}

func serveSwagger(dir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "certificateSigningRequest") {
			glog.Errorf("Not Found: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}

		glog.Infof("Serving %s", r.URL.Path)
		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
		p = path.Join(dir, p)
		http.ServeFile(w, r, p)
	}
}

// allowCORS allows Cross Origin Resoruce Sharing from a specified origin.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" && contains(cors, origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				headers := []string{"Content-Type", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"}
				w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
				methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
				w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func mockRoot(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set our conditions to hijack the session
		if strings.Replace(r.URL.Path, "/", "", 1) == "" && r.Method == http.MethodGet {

			// TODO: Add something more useful like version and stage
			response := "PassKit IO API"

			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Length", strconv.Itoa(len(response)))
			r.Header.Set("Content-Length", strconv.Itoa(len(response)))
			w.WriteHeader(http.StatusOK)

			_, _ = w.Write([]byte(response))
			// Don't forget to return to make sure you don't send twice!
			return
		}
		// Or if the URL doesn't match serve the original HTML
		h.ServeHTTP(w, r)
	})
}

func contains(arr []string, check string) bool {
	for _, s := range arr {
		if s == check || s == "*" {
			return true
		}
	}
	return false
}

func healthCheck(conn *grpc.ClientConn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		if s := conn.GetState(); s != connectivity.Ready {
			http.Error(w, fmt.Sprintf("grpc server is %s", s), http.StatusBadGateway)
			return
		}
		_, _ = fmt.Fprintln(w, "ok")
	}
}
