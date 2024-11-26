/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package helloword

import (
	"context"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/transport/http/binding"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"examples/helloword/services/helloworld/v1"
	"github.com/origadmin/toolkits/runtime/transport/gins"
)

var _ helloworld.GreeterServiceHTTPServer = &helloServer{}

type helloServer struct {
	helloworld.UnimplementedGreeterServiceServer
	cli helloworld.GreeterServiceClient
}

func (h helloServer) SayHello(ctx context.Context, request *helloworld.SayHelloRequest) (*helloworld.SayHelloResponse, error) {
	var out helloworld.SayHelloResponse
	out.Message = strconv.FormatInt(int64(rand.Intn(100)), 10)
	c, ok := gins.FromContext(ctx)
	if ok {
		log.Info("Request:", c.FullPath())
		return h.cli.SayHello(ctx, request)
	}
	log.Info("Request RPC:", "hello ", request.Name, " give ", out.Message)
	return &out, nil
}

func TestServer(t *testing.T) {
	hsrv := gins.NewServer(
		gins.Address(":8000"),
	)
	gsrv := transgrpc.NewServer(
		transgrpc.Address(":9000"),
	)
	srv := kratos.New(
		kratos.Server(gsrv, hsrv),
	)
	con, _ := transgrpc.DialInsecure(
		context.Background(),
		transgrpc.WithEndpoint("dns:///127.0.0.1:9000"),
		transgrpc.WithMiddleware(
			func(handler middleware.Handler) middleware.Handler {
				log.Info("Middleware Call")
				return func(ctx context.Context, req interface{}) (interface{}, error) {
					c, ok := gins.FromContext(ctx)
					if ok {
						log.Info("Gin trigger middleware", c.FullPath(), " args ", c.Params)
					}

					return handler(ctx, req)
				}
			},
		),
	)
	s := &helloServer{
		cli: helloworld.NewGreeterServiceClient(con),
	}

	helloworld.RegisterGreeterServiceServer(gsrv, s)
	helloworld.RegisterGreeterServiceGINSServer(hsrv, s)
	hsrv.Use(gins.Recovery(log.DefaultLogger, true))
	hsrv.Use(gins.Logger(log.DefaultLogger))

	hsrv.GET("/login/*param", func(c *gin.Context) {
		if len(c.Params.ByName("param")) > 1 {
			c.AbortWithStatus(404)
			return
		}
		c.String(200, "Hello World!")
	})

	hsrv.GET("/helloworld", func(c *gin.Context) {
		var out helloworld.SayHelloResponse
		out.Message = strconv.FormatInt(int64(rand.Intn(100)), 10)
		c.JSON(200, &out)
	})
	go func() {
		time.Sleep(15 * time.Second)
		srv.Stop()
	}()
	if err := srv.Run(); err != nil {
		panic(err)
	}
	//testClient(t)
	//testGinClient(t)
	//testGRPCClient(t)
}

func testClient(t *testing.T) {
	ctx := context.Background()

	cli, err := transhttp.NewClient(ctx,
		transhttp.WithEndpoint("127.0.0.1:8000"),
	)
	assert.Nil(t, err)
	assert.NotNil(t, cli)

	resp, err := GetHelloReply(ctx, cli, nil, transhttp.EmptyCallOption{})
	assert.Nil(t, err)
	t.Log(resp)
}

func GetHelloReply(ctx context.Context, cli *transhttp.Client, in *helloworld.SayHelloRequest, opts ...transhttp.CallOption) (*helloworld.SayHelloResponse, error) {
	var out helloworld.SayHelloResponse

	pattern := "/helloworld"
	path := binding.EncodeURL(pattern, in, true)

	opts = append(opts, transhttp.Operation(helloworld.GreeterService_SayHello_OperationName))
	opts = append(opts, transhttp.PathTemplate(pattern))

	err := cli.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func testGinClient(t *testing.T) {
	ctx := context.Background()

	cli, err := transhttp.NewClient(ctx,
		transhttp.WithEndpoint("127.0.0.1:8000"),
	)
	assert.Nil(t, err)
	c := helloworld.NewGreeterServiceGINSClient(cli)
	resp, err := c.SayHello(context.Background(), nil, transhttp.EmptyCallOption{})
	assert.Nil(t, err)
	t.Log(resp)
}

func testGRPCClient(t *testing.T) {
	ctx := context.Background()

	cli, err := transgrpc.DialInsecure(ctx,
		transgrpc.WithEndpoint("127.0.0.1:8000"),
	)
	assert.Nil(t, err)
	c := helloworld.NewGreeterServiceClient(cli)
	resp, err := c.SayHello(context.Background(), &helloworld.SayHelloRequest{
		Name: "Mynameisworld",
	}, grpc.EmptyCallOption{})
	assert.Nil(t, err)
	t.Log(resp)
}
