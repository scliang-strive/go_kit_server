package grpc

import (
	"context"
	"fmt"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	endpoint "kit_server02/endpoint/article"
	"kit_server02/modle/article"
	pb "kit_server02/pb/article"
	service "kit_server02/service/article"
)

func decodeCreateRequest(c context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pb.CreateReq)
	if !ok {
		return nil, fmt.Errorf("grpc server decode request出错！")
	}
	// 过滤数据
	r := &article.CreateReq{
		Title:   req.Title,
		Content: req.Content,
		CateId:  req.CateId,
	}
	return r, nil
}

func encodeCreateResponse(c context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(*article.CreateResp)

	if !ok {
		return nil, fmt.Errorf("grpc server encode response error (%T)", response)
	}
	r := &pb.CreateResp{
		Id: resp.Id,
	}
	return r, nil
}

func decodeDetailReq(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pb.DetailReq)
	if !ok {
		return nil, fmt.Errorf("grpc server decode detail request出错！")
	}
	r := article.DetailReq{
		Id: req.Id,
	}
	return r, nil
}

func encodeDetailResp(_ context.Context, grpcResp interface{}) (interface{}, error) {
	resp, ok := grpcResp.(*pb.DetailResp)
	if !ok {
		return nil, fmt.Errorf("grpc server encode detail response error (%T)", resp)
	}
	r := &pb.DetailResp{
		Title:   resp.Title,
		CateId:  resp.CateId,
		Content: resp.Content,
		UserId:  resp.UserId,
	}
	return r, nil
}

type ArticleGrpcServer struct {
	createHandler grpctransport.Handler
	detailHandler grpctransport.Handler
}

func (s *ArticleGrpcServer) Create(ctx context.Context, req *pb.CreateReq) (*pb.CreateResp, error) {
	_, rsp, err := s.createHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*pb.CreateResp), nil
}

func (s *ArticleGrpcServer) Detail(ctx context.Context, req *pb.DetailReq) (*pb.DetailResp, error) {
	_, rsp, err := s.detailHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*pb.DetailResp), nil
}

func NewArticleGrpcServer(svc service.ArticleService, opts ...grpctransport.ServerOption) pb.ArticleServiceServer {
	createHandler := grpctransport.NewServer(
		endpoint.MakeCreateEndpoint(svc),
		decodeCreateRequest,
		encodeCreateResponse,
		opts...,
	)

	detailHandler := grpctransport.NewServer(
		endpoint.MakeDetailedEndpoint(svc),
		decodeDetailReq,
		encodeDetailResp,
		opts...,
	)

	articleGrpServer := new(ArticleGrpcServer)
	articleGrpServer.createHandler = createHandler
	articleGrpServer.detailHandler = detailHandler

	return articleGrpServer
}
