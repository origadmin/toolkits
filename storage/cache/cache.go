package cache

import (
	"time"

	"github.com/origadmin/toolkits/context"
)

// Cache is the interface that wraps the basic Get, Set, and Delete methods.
// It uses string as the value type, allowing for flexible conversion to []byte
// and zero-copy resource optimization after go1.22.
type Cache interface {
	// Get retrieves the value associated with the given key.
	// It returns the cached value and an error if the key is not found.
	Get(ctx context.Context, key string) (string, error)

	// GetAndDelete retrieves the value associated with the given key and deletes it.
	// It returns the cached value and an error if the key is not found.
	GetAndDelete(ctx context.Context, key string) (string, error)

	// Exists checks if a value exists for the given key.
	// It returns an error if the key is not found.
	Exists(ctx context.Context, key string) error

	// Set sets the value for the given key.
	// It returns an error if the operation fails.
	Set(ctx context.Context, key string, value string, expiration ...time.Duration) error

	// Delete deletes the value associated with the given key.
	// It returns an error if the operation fails.
	Delete(ctx context.Context, key string) error
}