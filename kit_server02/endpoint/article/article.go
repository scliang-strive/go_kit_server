package article

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"kit_server02/errors"
	"kit_server02/modle/article"
	service "kit_server02/service/article"
)

func MakeCreateEndpoint(srv service.ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*article.CreateReq)
		if !ok {
			return nil, errors.EndpointTypeError
		}
		resp, err := srv.Create(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

func MakeDetailedEndpoint(srv service.ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*article.DetailReq)
		if !ok {
			return nil, errors.EndpointTypeError
		}
		resp, err := srv.Detailed(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
