package loge

import (
	"github.com/goexts/ggb/settings"
)

type Setting = settings.Setting[Option]

// Option custom setup config
type Option struct {
	OutputPath       string
	FileName         string
	Format           Format
	TimeLayout       string
	DisableConsole   bool
	Level            Leveler
	ReplaceAttr      func(groups []string, attr Attr) Attr
	AddSource        bool
	LumberjackConfig *LumberjackConfig
	DevConfig        *DevConfig
	NoColor          bool
	UseDefault       bool
}

var (
	defaultOption = &Option{
		OutputPath: "",
		FileName:   "output.log",
		Format:     FormatText,
		TimeLayout: DefaultTimeLayout,
		Level:      LevelDebug,
	}
)

// WithPath custom path to write log
func WithPath(path string) Setting {
	return func(opt *Option) {
		opt.OutputPath = path
	}
}

// WithFile write log to some File
func WithFile(file string) Setting {
	return func(opt *Option) {
		opt.FileName = file
	}
}

// WithLumberjack write log to some File with rotation
func WithLumberjack(filename string, config *LumberjackConfig) Setting {
	return func(opt *Option) {
		opt.FileName = filename
		opt.LumberjackConfig = config
	}
}

// WithTimeLayout custom time format
func WithTimeLayout(timeLayout string) Setting {
	return func(opt *Option) {
		opt.TimeLayout = timeLayout
	}
}

// WithDisableConsole WithEnableConsole write log to os.Stdout or os.Stderr
func WithDisableConsole() Setting {
	return func(opt *Option) {
		opt.DisableConsole = true
	}
}

// WithLevel custom log level
func WithLevel(level Leveler) Setting {
	return func(opt *Option) {
		opt.Level = level
	}
}

// WithReplaceAttr custom replaceAttr
func WithReplaceAttr(replaceAttr func(groups []string, attr Attr) Attr) Setting {
	return func(opt *Option) {
		opt.ReplaceAttr = replaceAttr
	}
}

// WithFormat custom format
func WithFormat(format Format) Setting {
	return func(opt *Option) {
		opt.Format = format
	}
}

// WithAddSource add source info to log
func WithAddSource() Setting {
	return func(opt *Option) {
		opt.AddSource = true
	}
}

// WithNoColor disable color
func WithNoColor() Setting {
	return func(opt *Option) {
		opt.NoColor = true
	}
}

// WithUseDefault use output as slog.Default()
func WithUseDefault() Setting {
	return func(opt *Option) {
		opt.UseDefault = true
	}
}

// WithDevConfig set dev config
func WithDevConfig(config *DevConfig) Setting {
	return func(opt *Option) {
		opt.DevConfig = config
	}
}
