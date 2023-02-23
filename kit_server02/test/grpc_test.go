package test

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "kit_server02/pb/article"
	"testing"
	"time"
)

func TestGrpcServer(t *testing.T) {
	conn, err := grpc.Dial("0.0.0.0:9999", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Errorf("test grpc server err: %s \n", err.Error())
	}
	defer func() {
		conn.Close()
	}()
	svc := pb.NewArticleServiceClient(conn)
	ctx, cancle := context.WithTimeout(context.Background(), time.Second)
	defer cancle()
	r, err := svc.Create(ctx, &pb.CreateReq{
		Title:   "grpc test",
		Content: "grpc test content",
		CateId:  int64(64),
	})
	//r, err := svc.Detail(ctx, &pb.DetailReq{
	//	Id: 9999,
	//})
	if err != nil {
		t.Errorf("request grpc err: %s \n", err.Error())
	}
	fmt.Println(r.GetId())
}
