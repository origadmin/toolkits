package logger

import (
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	logger "github.com/origadmin/slog-kratos"

	"github.com/origadmin/toolkits/loge"
	"github.com/origadmin/toolkits/runtime/config"
)

func Middleware(cfg *config.LoggerConfig, l log.Logger) (middleware.Middleware, error) {
	if cfg == nil && l == nil {
		// todo: init l from config
		l = log.NewStdLogger(os.Stdout)
		return logging.Server(l), nil
	}

	nlog := loge.New(FromLoggerConfig(cfg))
	l = logger.NewLogger(logger.WithLogger(nlog))
	return logging.Server(l), nil
}

func NewLogger(cfg *config.LoggerConfig) log.Logger {
	return log.NewStdLogger(os.Stdout)
}

func FromLoggerConfig(cfg *config.LoggerConfig) loge.Setting {
	return func(option *loge.Option) {
		//option.OutputPath = cfg.Path
		//option.Format = cfg.Format
		//option.TimeLayout = cfg.TimeLayout
		//option.DisableConsole = cfg.DisableConsole
		//option.Level = cfg.Level
		//option.ReplaceAttr = cfg.ReplaceAttr
		//option.AddSource = cfg.AddSource
		//option.LumberjackConfig = cfg.Lumberjack
		//option.DevConfig = cfg.Dev
		//option.NoColor = cfg.NoColor
		//option.UseDefault = cfg.UseDefault
	}
}
