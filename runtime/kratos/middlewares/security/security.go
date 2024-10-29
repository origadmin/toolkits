package security

import (
	"github.com/go-kratos/kratos/v2/middleware"

	"github.com/origadmin/toolkits/context"
	"github.com/origadmin/toolkits/middlewares"
)

func Middleware(config *middlewares.SecurityConfig) (middleware.Middleware, error) {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			return handler(ctx, req)
		}
	}, nil
}
