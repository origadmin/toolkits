package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/origadmin/toolkits/metrics"
)

const (
	defaultApplication = "origadmin"
	defaultNamespace   = "backend"
	defaultSubSystem   = "http"
	defaultListenPort  = 9100
)

type Config struct {
	LogHandler     map[string]struct{}
	LogMethod      map[string]struct{}
	Application    string
	Namespace      string
	SubSystem      string
	Buckets        []float64
	Objectives     map[float64]float64
	DefaultCollect bool
	MetricLabels   map[metrics.MetricType][]string
	UseSecure      bool
	BasicUserName  string
	BasicPassword  string
	HandlerFunc    http.HandlerFunc
}

func (c *Config) setup() {
	if c.Application == "" {
		c.Application = defaultApplication
	}
	if c.Namespace == "" {
		c.Namespace = defaultNamespace
	}

	if c.SubSystem == "" {
		c.SubSystem = defaultSubSystem
	}

	// Set default listen port if not provided and enable Prometheus.
	if c.Enable && c.ListenPort == 0 {
		c.ListenPort = defaultListenPort
	}

	if len(c.Buckets) == 0 {
		c.Buckets = prometheus.DefBuckets
	}

	if len(c.Objectives) == 0 {
		c.Objectives = map[float64]float64{0.9: 0.01, 0.95: 0.005, 0.99: 0.001}
	}

	labels := metrics.MetricLabels()
	for i, l := range c.MetricLabels {
		if len(l) > 0 {
			labels[i] = l
		}
	}
	c.MetricLabels = labels
}
