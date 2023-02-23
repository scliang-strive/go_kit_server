package main

import (
	"fmt"
	grpcserver "kit_server02/server/grpc"
	httpserver "kit_server02/server/http"
)

func main() {
	errc := make(chan error)
	go httpserver.Run("0.0.0.0:9998", errc)
	go grpcserver.Run("0.0.0.0:9999", errc)
	fmt.Printf("error: %v Exit", <-errc)
}
