/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package sloge implements the functions, types, and interfaces for the module.
package sloge

type Setting = func(*Option)

// Option custom setup config
type Option struct {
	Output           string
	Format           Format
	TimeLayout       string
	Console          bool
	Level            Leveler
	ReplaceAttr      func(groups []string, attr Attr) Attr
	AddSource        bool
	LumberjackConfig *LumberjackLogger
	DevConfig        *DevConfig
	NoColor          bool
	Default          bool
}

var (
	defaultOption = &Option{
		Output:     "output.log",
		Format:     FormatText,
		TimeLayout: DefaultTimeLayout,
		Level:      LevelInfo,
	}
)

// WithFile write log to some File
func WithFile(file string) Setting {
	return func(opt *Option) {
		opt.Output = file
	}
}

// WithLumberjack write log to some File with rotation
func WithLumberjack(config *LumberjackLogger) Setting {
	return func(opt *Option) {
		opt.LumberjackConfig = config
	}
}

// WithTimeLayout custom time format
func WithTimeLayout(timeLayout string) Setting {
	return func(opt *Option) {
		opt.TimeLayout = timeLayout
	}
}

// WithConsole set the log to console or /dev/null
func WithConsole(set bool) Setting {
	return func(opt *Option) {
		opt.Console = set
	}
}

// WithConsoleOnly set the log to console only
func WithConsoleOnly() Setting {
	return func(opt *Option) {
		opt.Console = true
		opt.Output = ""
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

// WithDefault use output as slog.Default()
func WithDefault(set bool) Setting {
	return func(opt *Option) {
		opt.Default = set
	}
}

// WithDevConfig set dev config
func WithDevConfig(config *DevConfig) Setting {
	return func(opt *Option) {
		opt.DevConfig = config
	}
}
