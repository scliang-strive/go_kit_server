package main

import (
	"context"
	"fmt"
	kitLog "github.com/go-kit/log"
	uuid "github.com/satori/go.uuid"
	myEndpoint "github.com/scliang-strive/go_kit_server/kit_server03_consul/pkg/endpoint"
	"github.com/scliang-strive/go_kit_server/kit_server03_consul/pkg/service"
	myConsul "github.com/scliang-strive/go_kit_server/kit_server03_consul/pkg/service/consul"
	"github.com/scliang-strive/go_kit_server/kit_server03_consul/pkg/transport"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var logger *log.Logger

func main() {
	logger = log.New(os.Stdout, "", 0)
	serviceHost := "10.62.78.120"
	servicePort := 9999
	serviceName := "SayHello"

	consulHost := "8.141.175.100"
	consulPort := 8501

	ctx := context.Background()
	errChan := make(chan error)
	var discoverClient myConsul.DiscoverClient
	discoverClient, err := myConsul.NewKitDiscoverClient(consulHost, consulPort)
	if err != nil {
		log.Println("Get Consul Client failed")
		return
	}

	svc := service.NewDiscoveryService(discoverClient)

	sayHelloEndpoint := myEndpoint.MakeSayHelloEndpoint(svc)
	healthCheckEndpoint := myEndpoint.MakeHealthCheckEndpoint(svc)

	endpoints := myEndpoint.DiscoveryEndpoint{
		SayHelloEndpoint:    sayHelloEndpoint,
		HealthCheckEndpoint: healthCheckEndpoint,
	}

	router := transport.MakeHttpHandler(ctx, endpoints, kitLog.NewNopLogger())

	instanceId := serviceName + "-" + uuid.NewV4().String()

	go func() {
		logger.Println("Http Server start at port:" + strconv.Itoa(servicePort))
		if !discoverClient.Register(serviceName, instanceId, "/health",
			serviceHost, servicePort, nil, logger) {
			logger.Printf("string-service for service %s failed.", serviceName)
			// 注册失败，服务启动失败
			os.Exit(-1)
		}

		handler := router
		errChan <- http.ListenAndServe(":"+strconv.Itoa(servicePort), handler)

	}()

	go func() {
		// 监控系统信号，等待 ctrl + c 系统信号通知服务关闭
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	err = <-errChan

	//服务退出取消注册
	deRegistryOk := discoverClient.DeRegister(instanceId, logger)
	if deRegistryOk {
		logger.Println("DeRegister Service Success!")
	}
	logger.Println(err)

}
