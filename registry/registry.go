package registry

import (
	"sync"

	"github.com/go-kratos/kratos/v2/registry"
)

type Registry interface {
	NewClient(conf Config) (registry.Discovery, error)
	NewServer(conf Config) (registry.Registrar, error)
}

var (
	registryMap   = make(map[string]Registry)
	registryMutex sync.Mutex
)

func Register(s string, r Registry) {
	registryMutex.Lock()
	defer registryMutex.Unlock()
	if _, ok := registryMap[s]; ok {
		panic("registry already registered")
	}
	registryMap[s] = r
}

func BuildDiscovery(config Config) (registry.Discovery, error) {
	registryMutex.Lock()
	defer registryMutex.Unlock()
	r, ok := registryMap[config.Type]
	if !ok {
		return nil, ErrNotFound
	}
	return r.NewClient(config)
}

func BuildRegistrar(config Config) (registry.Registrar, error) {
	registryMutex.Lock()
	defer registryMutex.Unlock()
	r, ok := registryMap[config.Type]
	if !ok {
		return nil, ErrNotFound
	}
	return r.NewServer(config)
}
