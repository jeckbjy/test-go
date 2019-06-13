package main

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"golang.org/x/net/trace"
	"io/ioutil"
	"log"
	profile "srv/profile/proto"
)

const ServiceName = "go.micro.srv.profile"

func main() {
	service := micro.NewService(micro.Name(ServiceName))
	service.Init()
	profile.RegisterProfileHandler(service.Server(), newProfileService())
	service.Run()
}

func newProfileService()*ProfileService {
	file, err := ioutil.ReadFile("data/profiles.json")
	if err != nil {
		log.Fatalf("load data fail:%+v\n", err)
		return nil
	}

	// unmarshal json profiles
	hotels := []*profile.Hotel{}
	if err := json.Unmarshal(file, &hotels); err != nil {
		log.Fatalf("Failed to load json: %v", err)
	}

	profiles := make(map[string]*profile.Hotel)
	for _, hotel := range hotels {
		profiles[hotel.Id] = hotel
	}

	return &ProfileService{profiles}
}

type ProfileService struct {
	hotels map[string]*profile.Hotel
}

// GetProfiles returns hotel profiles for requested IDs
func (s *ProfileService) GetProfiles(ctx context.Context, req *profile.Request, rsp *profile.Result) error {
	md, _ := metadata.FromContext(ctx)
	traceID := md["traceID"]
	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("traceID %s", traceID)
	}

	for _, i := range req.HotelIds {
		rsp.Hotels = append(rsp.Hotels, s.hotels[i])
	}

	return nil
}