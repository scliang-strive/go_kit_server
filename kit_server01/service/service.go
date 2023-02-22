package service

import (
	"context"
	"time"
)

type Service interface {
	// Status 返回一个状态确认服务开启
	Status(ctx context.Context) (string, error)
	// Get 返回当前的日期
	Get(ctx context.Context) (string, error)
	// Validate 接收dd/mm/yyyy的时间串，根据正则对其校验
	Validate(ctx context.Context, date string) (bool, error)
}

type dateService struct{}

func NewService() Service {
	return dateService{}
}

func (dateService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

func (dateService) Get(ctx context.Context) (string, error) {
	now := time.Now()
	return now.Format("02/01/2006"), nil
}

func (dateService) Validate(ctx context.Context, date string) (bool, error) {
	_, err := time.Parse("02/01/2006", date)
	if err != nil {
		return false, err
	}
	return true, nil
}
