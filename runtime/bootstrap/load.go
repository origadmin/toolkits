package bootstrap

import (
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/origadmin/toolkits/codec"

	"github.com/origadmin/toolkits/runtime/config"
)

func LoadConfig(path string) config.Config {
	conf := config.De
	codec.DecodeFromFile(path, &Config)

	return config.New(
		config.WithSource(
			file.NewSource(path),
		),
	)
}
