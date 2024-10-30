package kratos

import (
	"sync"

	"github.com/go-kratos/kratos/v2/registry"

	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/runtime/config"
)

var (
	once  = &sync.Once{}
	build = &builder{}
)

func init() {
	build.init()
}

type ConfigBuilder func(*config.SourceConfig, ...config.Option) (config.Config, error)

type RegistryBuilder interface {
	NewDiscovery(cfg *config.RegistryConfig) (registry.Discovery, error)
	NewRegistrar(cfg *config.RegistryConfig) (registry.Registrar, error)
}

type builder struct {
	mutex      sync.RWMutex
	configs    map[string]ConfigBuilder
	registries map[string]RegistryBuilder
}

func (b *builder) init() {
	once.Do(func() {
		b.configs = make(map[string]ConfigBuilder)
		b.registries = make(map[string]RegistryBuilder)
	})
}

func RegistryConfig(name string, configBuilder ConfigBuilder) {
	build.mutex.Lock()
	defer build.mutex.Unlock()
	build.configs[name] = configBuilder
}

var ErrNotFound = errors.String("not found")

func NewConfig(cfg *config.SourceConfig, opts ...config.Option) (config.Config, error) {
	build.mutex.RLock()
	defer build.mutex.RUnlock()
	configBuilder, ok := build.configs[cfg.Type]
	if !ok {
		return nil, ErrNotFound
	}
	return configBuilder(cfg, opts...)
}

func RegistryRegistry(name string, registryBuilder RegistryBuilder) {
	build.mutex.Lock()
	defer build.mutex.Unlock()
	build.registries[name] = registryBuilder
}

func NewDiscovery(cfg *config.RegistryConfig) (registry.Discovery, error) {
	build.mutex.RLock()
	defer build.mutex.RUnlock()
	discoveryBuilder, ok := build.registries[cfg.Type]
	if !ok {
		return nil, ErrNotFound
	}
	return discoveryBuilder.NewDiscovery(cfg)
}

func NewRegistrar(cfg *config.RegistryConfig) (registry.Registrar, error) {
	build.mutex.RLock()
	defer build.mutex.RUnlock()
	discoveryBuilder, ok := build.registries[cfg.Type]
	if !ok {
		return nil, ErrNotFound
	}
	return discoveryBuilder.NewRegistrar(cfg)
}
