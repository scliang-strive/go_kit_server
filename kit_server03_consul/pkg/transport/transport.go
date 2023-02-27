package transport

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	myEndpoint "github.com/scliang-strive/go_kit_server/kit_server03_consul/pkg/endpoint"
	"net/http"
)

var (
	ErrorRequest = errors.New("invalid request parameter")
)

func MakeHttpHandler(ctx context.Context, endpoints myEndpoint.DiscoveryEndpoint, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}
	r.Methods("GET").Path("/SayHello").Handler(kithttp.NewServer(
		endpoints.SayHelloEndpoint,
		decodeSayHelloRequest,
		encodeJsonResponse,
		options...,
	))

	r.Methods("GET").Path("/health").Handler(kithttp.NewServer(
		endpoints.HealthCheckEndpoint,
		decodeHealthRequest,
		encodeJsonResponse,
		options...,
	))
	return r
}

func decodeSayHelloRequest(ctx context.Context, req *http.Request) (request interface{}, err error) {
	return myEndpoint.SayHelloRequest{}, nil
}

func decodeHealthRequest(ctx context.Context, req *http.Request) (request interface{}, err error) {
	return myEndpoint.HealthCheckRequest{}, nil
}

func encodeJsonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
