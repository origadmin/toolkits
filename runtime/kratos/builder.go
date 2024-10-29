package kratos

import (
	"sync"

	"github.com/origadmin/toolkits/runtime/config"
)

var (
	build = &builder{}
)

func init() {
	build.init()
}

type ConfigBuilder func(config *config.DiscoveryConfig) (config.Config, error)

type builder struct {
	once    sync.Once
	mutex   sync.Mutex
	configs map[string]ConfigBuilder
}

func (b *builder) init() {
	b.once.Do(func() {
		b.configs = make(map[string]ConfigBuilder)
	})
}

func RegistryConfig(name string, configBuilder ConfigBuilder) {
	build.mutex.Lock()
	defer build.mutex.Unlock()
	build.configs[name] = configBuilder
}
