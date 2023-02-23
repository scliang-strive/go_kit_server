package http

import (
	httprouter "kit_server02/router/http"
	"net"
	"net/http"
)

var mux = http.NewServeMux()

var httpServer = http.Server{Handler: mux}

// http run
func Run(addr string, errc chan error) {

	// 注册路由
	httprouter.RegisterRouter(mux)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		errc <- err
		return
	}
	errc <- httpServer.Serve(lis)
}
