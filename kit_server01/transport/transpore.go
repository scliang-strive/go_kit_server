package transport

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetRequest struct{}

type GetResponse struct {
	Date string `json:"date"`
	Err  string `json:"err"`
}

type ValidateRequest struct {
	Date string `json:"date"`
}

type ValidateResponse struct {
	Valid bool   `json:"valid"`
	Err   string `json:"err,omitempty"`
}

type StatusRequest struct{}

type StatusResponse struct {
	Status string `json:"status"`
}

// 定义所有的请求编码器
func DecodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetRequest
	return req, nil
}

func DecodeValidateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req ValidateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req StatusRequest
	return req, nil
}

// 定义响应的解码器
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(&response)

}
