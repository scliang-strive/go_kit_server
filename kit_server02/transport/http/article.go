package http

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	endpoint "kit_server02/endpoint/article"
	model "kit_server02/modle/article"
	service "kit_server02/service/article"
	"net/http"
)

func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &model.CreateReq{}
	err := json.NewDecoder(r.Body).Decode(req)
	return req, err
}

func encodeCreateResponse(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}

func decodeDetailedRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &model.DetailReq{}
	err := json.NewDecoder(r.Body).Decode(req)
	return req, err
}

func encodeDetailedResponse(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}

func MakeCreateHandler(svc service.ArticleService) http.Handler {
	return httptransport.NewServer(
		endpoint.MakeCreateEndpoint(svc),
		decodeCreateRequest,
		encodeCreateResponse,
	)
}

func MakeDetailedHandler(svc service.ArticleService) http.Handler {
	return httptransport.NewServer(
		endpoint.MakeDetailedEndpoint(svc),
		decodeDetailedRequest,
		encodeDetailedResponse,
	)
}
