package bootstrap

import (
	"github.com/origadmin/toolkits/codec"

	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/runtime/config"
	"github.com/origadmin/toolkits/runtime/kratos"
)

func LoadConfig(boot *Bootstrap) (config.Config, error) {
	var cfg config.SourceConfig
	err := codec.DecodeFromFile(boot.ConfigPath, &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "LoadConfig")
	}

	return kratos.NewConfig(&cfg)
}
