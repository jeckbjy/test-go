package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"golang.org/x/net/trace"
	"io/ioutil"
	"log"
	"os"

	"srv/auth/proto"
)

const ServiceName = "go.micro.srv.auth"

func main() {
	log.Printf("args:%+v\n", os.Args)
	service := micro.NewService(micro.Name(ServiceName))
	service.Init()
	auth.RegisterAuthHandler(service.Server(), newAuthService())
	service.Run()
}

func newAuthService() *AuthService {
	//file := data.MustAsset(path)
	file, err := ioutil.ReadFile("data/customers.json")
	if err != nil {
		log.Fatalf("load data fail:%+v\n", err)
		return nil
	}

	var customers []*auth.Customer

	// unmarshal JSON
	if err := json.Unmarshal(file, &customers); err != nil {
		log.Fatalf("Failed to unmarshal json: %v", err)
	}

	// create customer lookup map
	cache := make(map[string]*auth.Customer)
	for _, c := range customers {
		cache[c.AuthToken] = c
	}

	return &AuthService{customers:cache}
}

type AuthService struct {
	customers map[string]*auth.Customer
}

func (s *AuthService) VerifyToken(ctx context.Context, req *auth.Request, rsp *auth.Result) error {
	md, _ := metadata.FromContext(ctx)
	traceID := md["traceID"]

	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("traceID %s", traceID)
	}

	customer := s.customers[req.AuthToken]
	if customer == nil {
		return errors.New("Invalid Token")
	}

	rsp.Customer = customer
	return nil
}