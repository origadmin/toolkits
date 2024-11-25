// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package grpc implements the functions, types, and interfaces for the module.
package grpc

import (
	"time"

	transhttp "github.com/go-kratos/kratos/v2/transport/http"

	"github.com/origadmin/toolkits/context"
	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/helpers"
	configv1 "github.com/origadmin/toolkits/runtime/gen/go/config/v1"
	"github.com/origadmin/toolkits/runtime/middleware"
	"github.com/origadmin/toolkits/runtime/registry"
	"github.com/origadmin/toolkits/runtime/service/selector"
)

const defaultTimeout = 5 * time.Second

// NewClient Creating an HTTP client instance.
func NewClient(ctx context.Context, r registry.Discovery, service *configv1.Service, m ...middleware.Middleware) (*transhttp.Client, error) {
	var ms []middleware.Middleware
	ms = middleware.NewClient(service.GetMiddleware())
	ms = append(ms, m...)

	timeout := defaultTimeout
	if serviceHttp := service.GetHttp(); serviceHttp != nil {
		if serviceHttp.Timeout != nil {
			timeout = serviceHttp.Timeout.AsDuration()
		}
	}

	options := []transhttp.ClientOption{
		transhttp.WithTimeout(timeout),
		transhttp.WithMiddleware(ms...),
	}

	if r != nil {
		endpoint := helpers.ServiceDiscoveryName(service.GetName())
		options = append(options,
			transhttp.WithEndpoint(endpoint),
			transhttp.WithDiscovery(r),
		)
	}

	if option, err := selector.NewHTTP(service.GetSelector()); err == nil {
		options = append(options, option)
	}
	conn, err := transhttp.NewClient(ctx, options...)

	if err != nil {
		return nil, errors.Errorf("dial grpc client [%s] failed: %s", service.GetName(), err.Error())
	}

	return conn, nil
}
