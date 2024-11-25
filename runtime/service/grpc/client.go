// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package grpc implements the functions, types, and interfaces for the module.
package grpc

import (
	"time"

	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/grpc"

	"github.com/origadmin/toolkits/context"
	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/helpers"
	configv1 "github.com/origadmin/toolkits/runtime/gen/go/config/v1"
	"github.com/origadmin/toolkits/runtime/middleware"
	"github.com/origadmin/toolkits/runtime/registry"
	"github.com/origadmin/toolkits/runtime/service/selector"
)

const defaultTimeout = 5 * time.Second

// NewClient Creating a GRPC client instance
func NewClient(ctx context.Context, r registry.Discovery, service *configv1.Service, m ...middleware.Middleware) (*grpc.ClientConn, error) {
	var ms []middleware.Middleware
	ms = middleware.NewClient(service.GetMiddleware())
	ms = append(ms, m...)

	timeout := defaultTimeout
	if serviceGrpc := service.GetGrpc(); serviceGrpc != nil {
		if serviceGrpc.Timeout != nil {
			timeout = serviceGrpc.Timeout.AsDuration()
		}
	}

	options := []transgrpc.ClientOption{
		transgrpc.WithTimeout(timeout),
		transgrpc.WithMiddleware(ms...),
	}

	if r != nil {
		endpoint := helpers.ServiceDiscoveryName(service.GetName())
		options = append(options,
			transgrpc.WithEndpoint(endpoint),
			transgrpc.WithDiscovery(r),
		)
	}

	if option, err := selector.NewGRPC(service.GetSelector()); err == nil {
		options = append(options, option)
	}
	conn, err := transgrpc.DialInsecure(ctx, options...)

	if err != nil {
		return nil, errors.Errorf("dial grpc client [%s] failed: %s", service.GetName(), err.Error())
	}

	return conn, nil
}
