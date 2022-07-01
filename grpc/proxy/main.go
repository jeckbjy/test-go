package main

import (
	"context"
	"log"
	"net/http"

	"grpc-demo/bgw"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var grpcClient grpc.ClientConnInterface

func main() {
	// Set up a connection to the server.
	addr := "localhost:50051"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	grpcClient = conn

	http.HandleFunc("/hello", proxy)
	http.ListenAndServe(":8080", nil)
}

func proxy(w http.ResponseWriter, r *http.Request) {
	c := bgw.NewServiceClient(grpcClient)
	req := &bgw.ProxyRequest{
		Path: r.URL.Path,
	}
	rsp, err := c.Proxy(context.Background(), req)
	if err != nil {
		log.Println("invoke fail", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(rsp.Message))
}
