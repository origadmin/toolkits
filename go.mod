// This function creates a new HTTP server using the provided configuration
func NewHTTPServiceServer(cfg *configv1.Service) (*service.HTTPServer, error) {
	// Call the build.NewHTTPServer function with the provided configuration
	return build.NewHTTPServer(cfg)
}

// This function creates a new HTTP client using the provided context and configuration
func NewHTTPServiceClient(ctx context.Context, cfg *configv1.Service) (*service.HTTPClient, error) {
	// Call the build.NewHTTPClient function with the provided context and configuration
	return build.NewHTTPClient(ctx, cfg)
}

// This function creates a new GRPC server using the provided configuration
func NewGRPCServiceServer(cfg *configv1.Service) (*service.GRPCServer, error) {
	// Call the build.NewGRPCServer function with the provided configuration
	return build.NewGRPCServer(cfg)
}

// This function creates a new GRPC client using the provided context and configuration
func NewGRPCServiceClient(ctx context.Context, cfg *configv1.Service) (*service.GRPCClient, error) {
	// Call the build.NewGRPCClient function with the provided context and configuration
	return build.NewGRPCClient(ctx, cfg)
}

// This function registers a service builder with the provided name
func RegisterService(name string, serviceBuilder ServiceBuilder) {
	// Call the build.RegisterServiceBuilder function with the provided name and service builder
	build.RegisterServiceBuilder(name, serviceBuilder)
}
// This function creates a new HTTP server using the provided configuration
func NewHTTPServiceServer(cfg *configv1.Service) (*service.HTTPServer, error) {
	// Call the build.NewHTTPServer function with the provided configuration
	return build.NewHTTPServer(cfg)
}

// This function creates a new HTTP client using the provided context and configuration
func NewHTTPServiceClient(ctx context.Context, cfg *configv1.Service) (*service.HTTPClient, error) {
	// Call the build.NewHTTPClient function with the provided context and configuration
	return build.NewHTTPClient(ctx, cfg)
}

// This function creates a new GRPC server using the provided configuration
func NewGRPCServiceServer(cfg *configv1.Service) (*service.GRPCServer, error) {
	// Call the build.NewGRPCServer function with the provided configuration
	return build.NewGRPCServer(cfg)
}

// This function creates a new GRPC client using the provided context and configuration
func NewGRPCServiceClient(ctx context.Context, cfg *configv1.Service) (*service.GRPCClient, error) {
	// Call the build.NewGRPCClient function with the provided context and configuration
	return build.NewGRPCClient(ctx, cfg)
}

// This function registers a service builder with the provided name
func RegisterService(name string, serviceBuilder ServiceBuilder) {
	// Call the build.RegisterServiceBuilder function with the provided name and service builder
	build.RegisterServiceBuilder(name, serviceBuilder)
}
module github.com/origadmin/toolkits

go 1.23.1

toolchain go1.23.2

require (
	github.com/coocood/freecache v1.2.4
	github.com/goexts/generic v0.1.0
	github.com/golang-cz/devslog v0.0.11
	github.com/lmittmann/tint v1.0.5
	github.com/nsqio/go-nsq v1.1.0
	github.com/origadmin/toolkits/errors v0.0.5
	github.com/stretchr/testify v1.9.0
	golang.org/x/crypto v0.29.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	golang.org/x/exp v0.0.0-20241108190413-2d47ceb2692f // indirect
	golang.org/x/sys v0.27.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.0.90
	v0.0.89
	v0.0.88
	v0.0.87
)
