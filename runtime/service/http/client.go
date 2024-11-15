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
	transhttp "github.com/go-kratos/kratos/v2/transport/http"

	"github.com/origadmin/toolkits/context"
	"github.com/origadmin/toolkits/errors"
	configv1 "github.com/origadmin/toolkits/runtime/gen/go/config/v1"
	"github.com/origadmin/toolkits/runtime/middleware"
	"github.com/origadmin/toolkits/runtime/registry"
	"github.com/origadmin/toolkits/utils"
)

const defaultTimeout = 5 * time.Second

// NewClient Creating a GRPC client
func NewClient(ctx context.Context, r registry.Discovery, service *configv1.Service, m ...middleware.Middleware) (*transhttp.Client, error) {
	endpoint := utils.NameDiscovery(service.GetName())

	var ms []middleware.Middleware

	ms = middleware.NewClient(service.GetMiddleware())
	ms = append(ms, m...)

	timeout := defaultTimeout

	serviceHttp := service.GetHttp()
	if serviceHttp != nil {
		if serviceHttp.Timeout != nil {
			timeout = serviceHttp.Timeout.AsDuration()
		}
	}

	options := []transhttp.ClientOption{
		transhttp.WithEndpoint(endpoint),
		transhttp.WithDiscovery(r),
		transhttp.WithTimeout(timeout),
		transhttp.WithMiddleware(ms...),
	}
	options = CreateSelectorOption(options, service.GetSelector())
	conn, err := transhttp.NewClient(ctx, options...)

	if err != nil {
		return nil, errors.Errorf("dial grpc client [%s] failed: %s", service.GetName(), err.Error())
	}

	return conn, nil
}

func CreateSelectorOption(options []transhttp.ClientOption, cfg *configv1.Service_Selector) []transhttp.ClientOption {
	if cfg == nil {
		return options
	}
	if cfg.Version != "" {
		v := filter.Version(cfg.Version)
		options = append(options, transhttp.WithNodeFilter(v))
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
