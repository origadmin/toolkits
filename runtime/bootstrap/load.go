package bootstrap

import (
	"github.com/origadmin/toolkits/codec"

	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/runtime/config"
)

func LoadSourceConfig(boot *Bootstrap) (*config.SourceConfig, error) {
	var cfg config.SourceConfig
	err := codec.DecodeFromFile(boot.ConfigPath, &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "LoadConfig")
	}

	return &cfg, nil
}
