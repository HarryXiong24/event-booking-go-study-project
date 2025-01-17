// Code generated by Kitex v0.12.0. DO NOT EDIT.
package bizservice

import (
	http "github.com/cloudwego/kitex-examples/generic/kitex_gen/http"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler http.BizService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	options = append(options, server.WithCompatibleMiddlewareForUnary())

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler http.BizService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}