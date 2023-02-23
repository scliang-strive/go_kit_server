package grpc

import (
	"google.golang.org/grpc"
	grpcrouter "kit_server02/router/grpc"
	"net"
)

var opts = []grpc.ServerOption{}

var grpcServer = grpc.NewServer(opts...)

func Run(addr string, errc chan error) {
	grpcrouter.RegisterRouter(grpcServer)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		errc <- err
		return
	}
	errc <- grpcServer.Serve(lis)
}
