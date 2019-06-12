package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

var ErrBadRouting = errors.New("inconsistent mapping between route and handler")

func NewHTTPHandler(endpoint endpoint.Endpoint) http.Handler {
	m := mux.NewRouter()
	server := httptransport.NewServer(endpoint, decodeHTTP, encodeHTTP)
	m.Methods(http.MethodGet).Path("/greeting").Handler(server)
	return m
}

func encodeHTTP(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	type errorWrapper struct {
		Error string `json:"error"`
	}

	if f, ok := response.(Failer); ok && f.Failed() != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorWrapper{Error: f.Failed().Error()})
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeHTTP(_ context.Context, r *http.Request) (interface{}, error) {
	vars := r.URL.Query()
	names, exists := vars["name"]
	if !exists || len(names) != 1 {
		return nil, ErrBadRouting
	}

	req := GreetingRequest{Name: names[0]}
	return req, nil
}
