package article

import (
	"context"
	"fmt"
	"kit_server02/modle/article"
)

type ArticleService interface {
	Create(ctx context.Context, req *article.CreateReq) (*article.CreateResp, error)
	Detailed(ctx context.Context, req *article.DetailReq) (*article.DetailResp, error)
}

type articleService struct {
}

func NewArticleService() ArticleService {
	svc := articleService{}
	{
		// middleware
	}
	return svc
}

func (articleService) Create(ctx context.Context, req *article.CreateReq) (*article.CreateResp, error) {
	fmt.Printf("req: %+v\n", req)
	return &article.CreateResp{
		Id: int64(999),
	}, nil
}

func (articleService) Detailed(ctx context.Context, req *article.DetailReq) (*article.DetailResp, error) {
	fmt.Printf("req: %+v \n", req)
	return &article.DetailResp{
		Id:      int64(99),
		Title:   "detailed...",
		Content: "xxxxx",
		CateId:  int64(99),
		UserId:  int64(99),
	}, nil
}
