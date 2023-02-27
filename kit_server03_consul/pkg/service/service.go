package service

import (
	"github.com/scliang-strive/go_kit_server/kit_server03_consul/pkg/service/consul"
)

type Service interface {
	HealthCheck() bool
	SayHello() string
}

// 接口实现
type DiscoverService struct {
	discoveryClient consul.DiscoverClient
}

func NewDiscoveryService(discoveryClient consul.DiscoverClient) Service {
	return &DiscoverService{
		discoveryClient: discoveryClient,
	}
}

func (service *DiscoverService) HealthCheck() bool {
	return true
}

func (service *DiscoverService) SayHello() string {
	return "Hello World!"
}
