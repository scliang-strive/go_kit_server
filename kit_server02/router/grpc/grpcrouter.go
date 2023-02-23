package grpc

import (
	"google.golang.org/grpc"
	pb "kit_server02/pb/article"
	service "kit_server02/service/article"
	transport "kit_server02/transport/grpc"
)

func RegisterRouter(grpcServer *grpc.Server) {
	pb.RegisterArticleServiceServer(grpcServer, transport.NewArticleGrpcServer(service.NewArticleService()))
}
