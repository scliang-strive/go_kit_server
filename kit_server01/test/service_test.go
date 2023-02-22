package test

import (
	"context"
	"kit_server01/service"
	"testing"
	"time"
)

func setup() (srv service.Service, ctx context.Context) {
	return service.NewService(), context.Background()
}

func TestStatus(t *testing.T) {
	srv, ctx := setup()
	s, err := srv.Status(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if s != "ok" {
		t.Errorf("status test error")
	}
}

func TestGet(t *testing.T) {
	srv, ctx := setup()
	d, err := srv.Get(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	timeNow := time.Now()
	today := timeNow.Format("02/01/2006")
	if today != d {
		t.Errorf("test get error")
	}
}

func TestValidate(t *testing.T) {
	srv, ctx := setup()
	d := time.Now().Format("02/01/2006")
	ok, err := srv.Validate(ctx, d)
	if !ok || err != nil {
		t.Errorf("Error: %s", err)
	}
	d = "12/32/1092"
	ok, err = srv.Validate(ctx, d)
	if err != nil || !ok {
		t.Errorf("test validate error")
	}

}
