package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/scliang-strive/go_kit_server/kit_server03_consul/pkg/service"
)

type DiscoveryEndpoint struct {
	SayHelloEndpoint    endpoint.Endpoint
	HealthCheckEndpoint endpoint.Endpoint
}

type SayHelloRequest struct {
}
type SayHelloResponse struct {
	Message string `json:"message"`
}

type HealthCheckRequest struct {
}
type HealthCheckResponse struct {
	Status bool `json:"status"`
}

func MakeSayHelloEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		message := svc.SayHello()
		return SayHelloResponse{Message: message}, nil
	}
}

func MakeHealthCheckEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		status := svc.HealthCheck()
		return HealthCheckResponse{
			Status: status,
		}, nil
	}
}
