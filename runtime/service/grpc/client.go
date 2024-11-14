// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package grpc implements the functions, types, and interfaces for the module.
package grpc

import (
	"time"

	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/filter"
	"github.com/go-kratos/kratos/v2/selector/p2c"
	"github.com/go-kratos/kratos/v2/selector/random"
	"github.com/go-kratos/kratos/v2/selector/wrr"
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/grpc"

	"github.com/origadmin/toolkits/context"
	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/runtime/config"
	"github.com/origadmin/toolkits/runtime/middleware"
	"github.com/origadmin/toolkits/runtime/registry"
	"github.com/origadmin/toolkits/utils"
)

const defaultTimeout = 5 * time.Second

// NewClient Creating a GRPC client
func NewClient(ctx context.Context, r registry.Discovery, service *config.Service, m ...middleware.Middleware) (*grpc.ClientConn, error) {
	endpoint := utils.NameDiscovery(service.GetName())

	var ms []middleware.Middleware

	ms = middleware.NewClient(service.GetMiddleware())
	ms = append(ms, m...)

	timeout := defaultTimeout

	serviceGrpc := service.GetGrpc()
	if serviceGrpc != nil {
		if serviceGrpc.Timeout != nil {
			timeout = serviceGrpc.Timeout.AsDuration()
		}
	}

	options := []transgrpc.ClientOption{
		transgrpc.WithEndpoint(endpoint),
		transgrpc.WithDiscovery(r),
		transgrpc.WithTimeout(timeout),
		transgrpc.WithMiddleware(ms...),
	}
	options = CreateSelectorOption(options, service.GetSelector())
	conn, err := transgrpc.DialInsecure(ctx, options...)

	if err != nil {
		return nil, errors.Errorf("dial grpc client [%s] failed: %s", service.GetName(), err.Error())
	}

	return conn, nil
}

func CreateSelectorOption(options []transgrpc.ClientOption, cfg *config.Service_Selector) []transgrpc.ClientOption {
	if cfg == nil {
		return options
	}
	if cfg.Version != "" {
		v := filter.Version(cfg.Version)
		options = append(options, transgrpc.WithNodeFilter(v))
	}

	var builder selector.Builder
	switch cfg.Builder {
	case "random":
		builder = random.NewBuilder()
	case "wrr":
		builder = wrr.NewBuilder()
	case "p2c":
		builder = p2c.NewBuilder()
	default:
	}

	if builder != nil {
		selector.SetGlobalSelector(builder)
	}

	return options
}
