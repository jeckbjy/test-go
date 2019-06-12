package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/web"
	"gomicro/pb"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// https://github.com/micro/go-micro
func main() {
	go runHTTP()
	go runGRPC()
	wait()
}

type Greeter struct{}

func (g *Greeter) Greeting(ctx context.Context, req *pb.GreetingRequest, rsp *pb.GreetingResponse) error {
	rsp.Greeting = "GO-MICRO Hello " + req.Name
	return nil
}

func runGRPC() {
	log.Printf("run grpc\n")
	service := micro.NewService(micro.Name("gomicro-srv-greeter"), micro.Version("latest"))
	service.Init()
	pb.RegisterGreeterHandler(service.Server(), &Greeter{})
	service.Run()
}

func runHTTP() {
	log.Printf("run http\n")
	service := web.NewService(
		web.Name("gomicro-web-greeter"),
		web.Address(":8110"),
	)

	service.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var name string
			vars := r.URL.Query()
			names, exists := vars["name"]
			if !exists || len(names) != 1 {
				name = ""
			} else {
				name = names[0]
			}

			cl := pb.NewGreeterService("gomicro-srv-greeter", client.DefaultClient)
			rsp, err := cl.Greeting(context.Background(), &pb.GreetingRequest{Name: name})
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			js, err := json.Marshal(rsp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
			return
		}
	})

	if service.Init() != nil {
		return
	}

	service.Run()
}

func wait() error {
	log.Printf("server start, wait signal\n")
	log.Printf("test url=> http://127.0.0.1:8110/greeting?name=test\n")
	//cancelInterrupt := make(chan struct{})
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	select {
	case sig := <-c:
		log.Printf("server stop")
		return fmt.Errorf("recevied signal %s", sig)
	}
}
