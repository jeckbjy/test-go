package main

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"golang.org/x/net/trace"
	"io/ioutil"
	"log"
	rate "srv/rate/proto"
)

const ServiceName = "go.micro.srv.rate"

func main() {
	service := micro.NewService(micro.Name(ServiceName))
	service.Init()
	rate.RegisterRateHandler(service.Server(), newService())
	service.Run()
}

func newService() *RateService {
	file, err := ioutil.ReadFile("data/rates.json")
	if err != nil {
		log.Fatalf("load data fail:%+v\n", err)
		return nil
	}

	rates := []*rate.RatePlan{}
	if err := json.Unmarshal(file, &rates); err != nil {
		log.Fatalf("Failed to load json: %v", err)
	}

	rateTable := make(map[stay]*rate.RatePlan)
	for _, ratePlan := range rates {
		stay := stay{
			HotelID: ratePlan.HotelId,
			InDate:  ratePlan.InDate,
			OutDate: ratePlan.OutDate,
		}
		rateTable[stay] = ratePlan
	}

	return &RateService{rateTable}
}

type stay struct {
	HotelID string
	InDate  string
	OutDate string
}

type RateService struct {
	rates map[stay]*rate.RatePlan
}

// GetRates gets rates for hotels for specific date range.
func (s *RateService) GetRates(ctx context.Context, req *rate.Request, rsp *rate.Result) error {
	md, _ := metadata.FromContext(ctx)
	traceID := md["traceID"]

	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("traceID %s", traceID)
	}

	for _, hotelID := range req.HotelIds {
		stay := stay{
			HotelID: hotelID,
			InDate:  req.InDate,
			OutDate: req.OutDate,
		}
		if s.rates[stay] != nil {
			rsp.RatePlans = append(rsp.RatePlans, s.rates[stay])
		}
	}

	return nil
}