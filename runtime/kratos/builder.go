package kratos

import (
	"sync"

	"github.com/go-kratos/kratos/v2/registry"

	"github.com/origadmin/toolkits/runtime/config"
)

var (
	once  = &sync.Once{}
	build = &builder{}
)

func init() {
	build.init()
}

type ConfigBuilder func(*config.DiscoveryConfig, ...config.Option) (config.Config, error)
type DiscoveryBuilder func(any) (registry.Discovery, error)

type builder struct {
	mutex       sync.RWMutex
	configs     map[string]ConfigBuilder
	discoveries map[string]DiscoveryBuilder
}

func (b *builder) init() {
	once.Do(func() {
		b.configs = make(map[string]ConfigBuilder)
	})
}

func RegistryConfig(name string, configBuilder ConfigBuilder) {
	build.mutex.Lock()
	defer build.mutex.Unlock()
	build.configs[name] = configBuilder
}

func RegistryDiscovery(name string, discoveryBuilder DiscoveryBuilder) {
	build.mutex.Lock()
	defer build.mutex.Unlock()
	build.discoveries[name] = discoveryBuilder
}
