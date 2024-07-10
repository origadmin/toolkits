package redisid

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/origadmin/toolkits/ident"
)

// Generator represents a Redis-based random ID Generator.
type Generator struct {
	Pool *redis.Pool // Pool is the connection Pool to the Redis server.
	size int         // size indicates the length of the generated ID string.
	mu   sync.Mutex  // mu ensures thread safety when generating IDs.
}

// redisIdentify is the global instance of the Redis random ID Generator.
var redisIdentify *Generator

// init initializes the global Redis random ID Generator when the package is first loaded.
func init() {
	// Initializes the random number Generator with a default seed.
	redisIdentify = New()
}

// Gen generates and returns a new random ID as a string using Redis.
func (i Generator) Gen() string {
	i.mu.Lock()
	defer i.mu.Unlock()

	conn := i.Pool.Get()
	defer conn.Close()

	// Generate a random ID and ensure its uniqueness by checking in Redis.
	for {
		id := fmt.Sprintf("%08x", rand.Uint32())
		_, err := conn.Do("SETNX", id, id)
		if err == nil {
			return id
		}
	}
}

// Validate checks if the provided ID string is unique within the Redis database.
func (i Generator) Validate(id string) bool {
	conn := i.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("GET", id)
	return errors.Is(err, redis.ErrNil)
}

// Size returns the size of the generated random ID string.
func (i Generator) Size() int {
	return i.size
}

// Name returns the name of the Generator, which is "redis".
func (i Generator) Name() string {
	return "redis"
}

// Default returns the global Redis random ID Generator instance.
func Default() ident.Identifier {
	return redisIdentify
}

// New creates a new Redis random ID Generator.
// It returns a pointer to the Generator.
func New() *Generator {
	return &Generator{
		Pool: &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", "localhost:6379")
				if err != nil {
					return nil, err
				}
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		},
		size: 8, // Default ID size.
	}
}
