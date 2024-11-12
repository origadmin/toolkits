package security

import (
	"github.com/origadmin/toolkits/context"
	"github.com/origadmin/toolkits/runtime/config"
	"github.com/origadmin/toolkits/runtime/middleware"
)

func Middleware(cfg *config.Security) (middleware.Middleware, error) {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			return handler(ctx, req)
		}
	}, nil
}
