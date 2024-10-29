package kratos

import (
	"sync"

	"github.com/origadmin/toolkits/runtime/kratos/discovery"
)

var (
	registryMap   = make(map[string]discovery.Registry)
	registryMutex sync.Mutex
)
