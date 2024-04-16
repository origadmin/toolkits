package service

import "context"

// Discovery provides an interface for service discovery
// and an abstraction over varying implementations
// {consul, etcd, zookeeper, ...}
type Discovery interface {
	Init(...Config) error
	ConfigSetting() Setting
	RegisterService(context.Context, Service, ...DiscoveryConfig) error
	Deregister(context.Context, Service, ...DiscoveryConfig) error
	GetService(context.Context, string, ...DiscoveryConfig) ([]Service, error)
	ListServices(context.Context, ...DiscoveryConfig) ([]*Service, error)
	Watch(context.Context, ...DiscoveryConfig) (Watcher, error)
	String() string
}
