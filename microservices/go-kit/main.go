package main

import (
	"fmt"
	"gokit/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// https://gokit.io/faq/#transports-mdash-which-transports-are-supported
// https://gokit.io/examples/stringsvc.html
/**
go-kit抽象出三个概念:service,endpoint,transport
service:定义抽象接口,实现相关服务逻辑
endpoint:定义内部协议，调用接口服务，每个函数调用都需要定义一个对应的endpoint
transport:将不同的协议的编码(比如grpc中使用protobuf,http使用json)转换成endpoint中定义的消息结构
总结:go-kit使用起来还是很麻烦的，需要手动写很多代码来定义service，endpoint，transport转换
*/
func main() {
	// step1:create service
	service := &GreeterSrv{}
	// step2:create endpoint
	endpoint := MakeGreetingEndpoint(service)
	// step3:create transport
	transGRPC := NewGRPCServer(endpoint)
	transHTTP := NewHTTPHandler(endpoint)
	go runGRPC(transGRPC)
	go runHTTP(transHTTP)
	// wait for
	wait()
}

func runGRPC(greeter pb.GreeterServer) {
	log.Printf("run grpc,listen 9120\n")
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, greeter)
	listener, err := net.Listen("tcp", ":9120")
	if err != nil {
		os.Exit(1)
	}

	err = server.Serve(listener)
	if err != nil {
		fmt.Errorf("%+v", err)
	}
}

func runHTTP(handler http.Handler) {
	log.Printf("run http,listen 9110\n")
	http.ListenAndServe(":9110", handler)
}

func wait() error {
	log.Printf("server start, waiting signal\n")
	log.Printf("test url=> http://127.0.0.1:9110/greeting?name=test\n")
	//cancelInterrupt := make(chan struct{})
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	select {
	case sig := <-c:
		log.Printf("server stop")
		return fmt.Errorf("recevied signal %s", sig)
	}
}
