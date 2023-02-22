package endpoint

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"kit_server01/service"
	"kit_server01/transport"
)

// Endpoints 公开端点
type Endpoints struct {
	GetEndpoint      endpoint.Endpoint
	StatusEndpoint   endpoint.Endpoint
	ValidateEndpoint endpoint.Endpoint
}

// MakeGetEndpoint 返回get服务的response
func MakeGetEndpoint(srv service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(transport.GetRequest)
		d, err := srv.Get(ctx)
		if err != nil {
			return transport.GetResponse{
				Date: d,
				Err:  err.Error(),
			}, nil
		}
		return transport.GetResponse{
			Date: d,
			Err:  "",
		}, nil
	}
}

// MakeStatusEndpoint 返回status服务的response
func MakeStatusEndpoint(srv service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(transport.StatusRequest)
		status, err := srv.Status(ctx)
		if err != nil {
			return transport.StatusResponse{Status: status}, err
		}
		return transport.StatusResponse{Status: status}, nil
	}
}

// MakeValidateEndpoint 返回validate服务的response
func MakeValidateEndpoint(srv service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.ValidateRequest)
		valid, err := srv.Validate(ctx, req.Date)
		if err != nil {
			return transport.ValidateResponse{Valid: valid, Err: err.Error()}, nil
		}
		return transport.ValidateResponse{Valid: valid}, nil
	}
}

// Get 端点映射
func (e Endpoints) Get(ctx context.Context) (string, error) {
	req := transport.GetRequest{}
	resp, err := e.GetEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	getResp := resp.(transport.GetResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Date, nil
}

// Status 映射
func (e Endpoints) Status(ctx context.Context) (string, error) {
	req := transport.StatusRequest{}
	resp, err := e.StatusEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	getResp := resp.(transport.StatusResponse)
	return getResp.Status, nil
}

// Validate 端点映射
func (e Endpoints) Validate(ctx context.Context, date string) (bool, error) {
	req := transport.ValidateRequest{Date: date}
	resp, err := e.ValidateEndpoint(ctx, req)
	if err != nil {
		return false, err
	}
	validateResp := resp.(transport.ValidateResponse)
	if validateResp.Err != "" {
		return false, errors.New(validateResp.Err)
	}
	return validateResp.Valid, nil
}
