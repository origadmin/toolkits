package metrics

import (
	"net/http"
	"time"
)

const (
	defaultApplication = "origadmin"
	defaultNamespace   = "backend"
	defaultSubSystem   = "http"
	defaultListenPort  = 9100
)

// Config represents the configuration for the Service server.
type Config struct {
	Enable          bool                    // Enable specifies if the server is enabled.
	ListenPort      int                     // ListenPort is the port on which the server listens.
	ReadTimeout     time.Duration           // ReadTimeout is the duration for reading the entire request.
	WriteTimeout    time.Duration           // WriteTimeout is the duration for writing the entire response.
	MaxHeaderBytes  int                     // MaxHeaderBytes is the maximum number of bytes the server will read parsing the request header.
	DurationBuckets []float64               // DurationBuckets are the buckets for histogram observations based on time.
	SizeBuckets     []float64               // SizeBuckets are the buckets for histogram observations based on size.
	Namespace       string                  // Namespace is the identifier for metrics in the namespace.
	SubSystem       string                  // SubSystem is the subsystem identifier for metrics.
	Buckets         []float64               // Buckets are the predefined histogram buckets.
	Objectives      map[float64]float64     // Objectives contains quantile ranks and their error allowances.
	DefaultCollect  bool                    // DefaultCollect specifies if metrics should be collected by default.
	BasicUserName   string                  // BasicUserName is the username for basic authentication.
	BasicPassword   string                  // BasicPassword is the password for basic authentication.
	MetricLabels    map[MetricType][]string // MetricLabels is the set of labels for each metric type.
	LogMethod       map[string]struct{}     // LogMethod is the set of methods to log.
	LogAPI          map[string]struct{}     // LogAPI is the set of APIs to log.
	Application     string                  // Application is the name of the application.
	Handler         http.Handler            // Handler is the HTTP handler
}

// DefaultConfig returns the default configuration for the Service server.
func DefaultConfig() *Config {
	return &Config{
		ListenPort:     defaultListenPort,
		LogMethod:      map[string]struct{}{},
		LogAPI:         map[string]struct{}{},
		Application:    defaultApplication,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Namespace:      defaultNamespace,
		SubSystem:      defaultSubSystem,
		Buckets:        []float64{},
		Objectives:     map[float64]float64{0.9: 0.01, 0.95: 0.005, 0.99: 0.001},
	}
}
