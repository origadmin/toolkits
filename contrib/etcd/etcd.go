/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package etcd

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	registryv2 "github.com/go-kratos/kratos/v2/registry"
	"github.com/origadmin/runtime/registry"
	etcdclient "go.etcd.io/etcd/client/v3"

	"github.com/origadmin/runtime"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
)

type etcdBuilder struct {
}

func init() {
	runtime.RegisterRegistry(Type, &etcdBuilder{})
}

func (c *etcdBuilder) NewDiscovery(cfg *configv1.Registry, _ ...registry.OptionSetting) (registryv2.Discovery, error) {
	config := fromConfig(cfg)
	etcdCli, err := etcdclient.New(config)
	if err != nil {
		return nil, err
	}
	r := etcd.New(etcdCli)
	return r, nil
}

func (c *etcdBuilder) NewRegistrar(cfg *configv1.Registry, _ ...registry.OptionSetting) (registryv2.Registrar, error) {
	config := fromConfig(cfg)
	etcdCli, err := etcdclient.New(config)
	if err != nil {
		return nil, err
	}
	r := etcd.New(etcdCli)
	return r, nil
}

func fromConfig(cfg *configv1.Registry) etcdclient.Config {
	etcdConfig := cfg.GetEtcd()
	if etcdConfig == nil {
		return etcdclient.Config{}
	}
	apiconfig := etcdclient.Config{
		Endpoints: etcdConfig.Endpoints,
	}
	//if etcdConfig.DialTimeout != 0 {
	//	apiconfig.DialTimeout = etcdConfig.DialTimeout
	//}
	//if etcdConfig.DialKeepAliveTime != 0 {
	//	apiconfig.DialKeepAliveTime = etcdConfig.DialKeepAliveTime
	//}
	//if etcdConfig.DialKeepAliveTimeout != 0 {
	//	apiconfig.DialKeepAliveTimeout = etcdConfig.DialKeepAliveTimeout
	//}
	//if etcdConfig.MaxCallRecvMsgSize != 0 {
	//	apiconfig.MaxCallRecvMsgSize = etcdConfig.MaxCallRecvMsgSize
	//}
	//if etcdConfig.MaxCallSendMsgSize != 0 {
	//	apiconfig.MaxCallSendMsgSize = etcdConfig.MaxCallSendMsgSize
	//}
	//apiconfig.TLS = etcdConfig.TLS
	return apiconfig
}
