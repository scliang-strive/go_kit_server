package consul

import (
	"fmt"
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
	"log"
	"sync"
)

type DiscoverClient interface {
	Register(serverName, instanceId, healthCheckUtl, instanceHost string, instancePort int, meta map[string]string, logger *log.Logger) bool
	DeRegister(instanceId string, logger *log.Logger) bool
	DiscoverServices(serviceName string, logger *log.Logger) []interface{}
}

type KitDiscoverClient struct {
	consulHost  string
	consulPort  int
	client      consul.Client
	config      *api.Config
	mutex       sync.Mutex
	instanceMap sync.Map
}

func NewKitDiscoverClient(consulHost string, consulPort int) (DiscoverClient, error) {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = fmt.Sprintf("%s:%d", consulHost, consulPort)
	log.Printf("consul host: %s", consulConfig.Address)
	apiClient, err := api.NewClient(consulConfig)
	if err != nil {
		return nil, err
	}
	client := consul.NewClient(apiClient)
	return &KitDiscoverClient{
		consulHost: consulHost,
		consulPort: consulPort,
		client:     client,
		config:     consulConfig,
	}, nil
}

func (cli *KitDiscoverClient) Register(serviceName, instanceId, healthCheckUrl string,
	instanceHost string, instancePort int, meta map[string]string, logger *log.Logger) bool {
	url := fmt.Sprintf("http://%s:%d%s", instanceHost, instancePort, healthCheckUrl)
	log.Printf("url: %s", url)
	serviceRegistration := &api.AgentServiceRegistration{
		ID:      instanceId,
		Name:    serviceName,
		Address: instanceHost,
		Port:    instancePort,
		Meta:    meta,
		Check: &api.AgentServiceCheck{
			DeregisterCriticalServiceAfter: "30s",
			HTTP:                           url,
			Interval:                       "300s",
		},
	}
	err := cli.client.Register(serviceRegistration)
	if err != nil {
		log.Printf("Register Service Error!")
		return false
	}
	log.Println("Register Service Success!")
	return true
}

func (cli *KitDiscoverClient) DeRegister(instanceId string, logger *log.Logger) bool {
	serviceRegistration := &api.AgentServiceRegistration{
		ID: instanceId,
	}
	err := cli.client.Deregister(serviceRegistration)
	if err != nil {
		logger.Println("Deregister Service Error!")
		return false
	}
	logger.Println("Deregister Service Success!")
	return true
}

func (cli *KitDiscoverClient) DiscoverServices(serviceName string, logger *log.Logger) []interface{} {
	// 判断服务是否已缓存
	instanceList, ok := cli.instanceMap.Load(serviceName)
	if ok {
		return instanceList.([]interface{})
	}
	cli.mutex.Lock()
	defer cli.mutex.Unlock()
	// 枷锁后在判断一次
	instanceList, ok = cli.instanceMap.Load(serviceName)
	if ok {
		return instanceList.([]interface{})
	}
	// 响应服务变更通知，更新服务map
	go func() {
		params := make(map[string]interface{})
		params["type"] = "service"
		params["service"] = serviceName
		plan, _ := watch.Parse(params)
		plan.Handler = func(u uint64, i interface{}) {
			if i == nil {
				return
			}
			v, ok := i.([]*api.ServiceEntry)
			if !ok {
				return
			}
			if len(v) == 0 {
				cli.instanceMap.Store(serviceName, []interface{}{})
			}
			var healthServices []interface{}
			for _, service := range v {
				if service.Checks.AggregatedStatus() == api.HealthPassing {
					healthServices = append(healthServices, service)
				}
			}
			cli.instanceMap.Store(serviceName, healthServices)
		}
		defer plan.Stop()
		plan.Run(cli.config.Address)
	}()
	// 调用go-kit 库想consul获取服务
	entries, _, err := cli.client.Service(serviceName, "", false, nil)
	if err != nil {
		cli.instanceMap.Store(serviceName, []interface{}{})
		logger.Println("Discover Service Error!")
		return nil
	}
	instances := make([]interface{}, 0, len(entries))
	for _, instance := range entries {
		instances = append(instances, instance)
	}
	cli.instanceMap.Store(serviceName, instances)
	return instances
}
