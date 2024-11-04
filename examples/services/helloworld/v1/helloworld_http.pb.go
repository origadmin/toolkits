// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             (unknown)
// source: helloworld/v1/helloworld.proto

package helloworld

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationGreeterServiceSayHello = "/helloworld.v1.GreeterService/SayHello"

type GreeterServiceHTTPServer interface {
	// SayHello Sends a greeting
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
}

func RegisterGreeterServiceHTTPServer(s *http.Server, srv GreeterServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/say_hello", _GreeterService_SayHello0_HTTP_Handler(srv))
	r.GET("/helloworld/{name}", _GreeterService_SayHello1_HTTP_Handler(srv))
}

func _GreeterService_SayHello0_HTTP_Handler(srv GreeterServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SayHelloRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterServiceSayHello)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SayHello(ctx, req.(*SayHelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SayHelloResponse)
		return ctx.Result(200, reply)
	}
}

func _GreeterService_SayHello1_HTTP_Handler(srv GreeterServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SayHelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterServiceSayHello)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SayHello(ctx, req.(*SayHelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SayHelloResponse)
		return ctx.Result(200, reply)
	}
}

type GreeterServiceHTTPClient interface {
	SayHello(ctx context.Context, req *SayHelloRequest, opts ...http.CallOption) (rsp *SayHelloResponse, err error)
}

type GreeterServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewGreeterServiceHTTPClient(client *http.Client) GreeterServiceHTTPClient {
	return &GreeterServiceHTTPClientImpl{client}
}

func (c *GreeterServiceHTTPClientImpl) SayHello(ctx context.Context, in *SayHelloRequest, opts ...http.CallOption) (*SayHelloResponse, error) {
	var out SayHelloResponse
	pattern := "/helloworld/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGreeterServiceSayHello))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
