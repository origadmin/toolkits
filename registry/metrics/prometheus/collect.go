package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Collectable is an interface for collecting metrics.
type Collectable interface {
	*prometheus.CounterVec
}

// Collectors collects metrics from the provided tables.
func Collectors[T Collectable](tables []T) []prometheus.Collector {
	var collectors []prometheus.Collector
	for i := range tables {
		collectors = append(collectors, tables[i])
	}
	return collectors
}

// CollectorsFromMap collects metrics from the provided tables.
func CollectorsFromMap[T Collectable, S comparable](tables map[S]T) []prometheus.Collector {
	var collectors []prometheus.Collector
	for i := range tables {
		collectors = append(collectors, tables[i])
	}
	return collectors
}
