package main

import (
	"context"
	"fmt"
	"kit_server01/endpoint"
	"kit_server01/server"
	"kit_server01/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	httpAddr := ":8080"
	ctx := context.Background()
	srv := service.NewService()
	errChan := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// 映射端点
	endpoints := endpoint.Endpoints{
		GetEndpoint:      endpoint.MakeGetEndpoint(srv),
		StatusEndpoint:   endpoint.MakeStatusEndpoint(srv),
		ValidateEndpoint: endpoint.MakeValidateEndpoint(srv),
	}

	// HTTP 传输
	go func() {
		log.Println("server listening on port: ", httpAddr)
		handler := server.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
