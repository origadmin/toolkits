package service

import (
	"github.com/origadmin/toolkits/context"
)

// Service represents the metrics service.
type Service interface {
	// Start starts the service.
	Start(ctx context.Context) error
	// Stop stops the service.
	Stop(ctx context.Context) error
}
