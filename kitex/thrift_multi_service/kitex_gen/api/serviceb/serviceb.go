// Code generated by Kitex v0.12.0. DO NOT EDIT.

package serviceb

import (
	"context"
	"errors"
	api "github.com/cloudwego/kitex-examples/thrift_multi_service/kitex_gen/api"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"echoB": kitex.NewMethodInfo(
		echoBHandler,
		newServiceBEchoBArgs,
		newServiceBEchoBResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	serviceBServiceInfo                = NewServiceInfo()
	serviceBServiceInfoForClient       = NewServiceInfoForClient()
	serviceBServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return serviceBServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return serviceBServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return serviceBServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "ServiceB"
	handlerType := (*api.ServiceB)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "api",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.12.0",
		Extra:           extra,
	}
	return svcInfo
}

func echoBHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.ServiceBEchoBArgs)
	realResult := result.(*api.ServiceBEchoBResult)
	success, err := handler.(api.ServiceB).EchoB(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newServiceBEchoBArgs() interface{} {
	return api.NewServiceBEchoBArgs()
}

func newServiceBEchoBResult() interface{} {
	return api.NewServiceBEchoBResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) EchoB(ctx context.Context, req *api.Request) (r *api.Response, err error) {
	var _args api.ServiceBEchoBArgs
	_args.Req = req
	var _result api.ServiceBEchoBResult
	if err = p.c.Call(ctx, "echoB", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
