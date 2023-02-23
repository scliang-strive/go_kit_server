package http

import (
	service "kit_server02/service/article"
	transporthttp "kit_server02/transport/http"
	"net/http"
)

func RegisterRouter(mux *http.ServeMux) {
	mux.Handle("/article/create", transporthttp.MakeCreateHandler(service.NewArticleService()))
	mux.Handle("/article/detail", transporthttp.MakeDetailedHandler(service.NewArticleService()))
}
