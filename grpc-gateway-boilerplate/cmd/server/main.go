package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	todov1 "github.com/levinhne/grpc-gateway-boilerplate/proto/todo/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding/gzip"
)

func main() {
	grpcPort := ":8080"
	go func() {
		server := grpc.NewServer()
		_, err := Initialize(server)
		if err != nil {
			panic(err)
		}
		listen, err := net.Listen("tcp", grpcPort)
		if err != nil {
			panic(err)
		}
		err = server.Serve(listen)
		if err != nil {
			panic(err)
		}
		log.Println("Grpc")
	}()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)),
	}
	ctx := context.Background()
	httpPort := ":8090"
	gwmux := runtime.NewServeMux()
	handlers := []func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error{
		todov1.RegisterTodoServiceHandlerFromEndpoint,
	}

	for _, h := range handlers {
		err := h(ctx, gwmux, grpcPort, opts)
		if err != nil {
			panic(err)
		}
	}
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	httpServer := &http.Server{
		Addr:    httpPort,
		Handler: allowCORS(mux),
	}
	err := httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
}
