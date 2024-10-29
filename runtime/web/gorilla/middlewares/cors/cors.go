package cors

import (
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/origadmin/toolkits/middlewares"
)

type FilterFunc = func(http.Handler) http.Handler

func WithCors(cfg *middlewares.CorsConfig) FilterFunc {
	opts := []handlers.CORSOption{
		handlers.AllowedOrigins(cfg.AllowOrigins),
		handlers.AllowedHeaders(cfg.AllowHeaders),
		handlers.AllowedMethods(cfg.AllowMethods),
		handlers.ExposedHeaders(cfg.ExposeHeaders),
		handlers.MaxAge(int(cfg.MaxAge.GetSeconds())),
	}
	if cfg.AllowCredentials {
		opts = append(opts, handlers.AllowCredentials())
	}

	return handlers.CORS(opts...)
}
