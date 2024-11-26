/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package security

import (
	"github.com/origadmin/toolkits/context"
	configv1 "github.com/origadmin/toolkits/runtime/gen/go/config/v1"
	"github.com/origadmin/toolkits/runtime/middleware"
)

func Middleware(cfg *configv1.Security) (middleware.Middleware, error) {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			return handler(ctx, req)
		}
	}, nil
}
