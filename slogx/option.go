/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package slogx implements enhanced logging functions for slog
package slogx

type Option = func(*Options)

// Options custom setup config
type Options struct {
	Output           string
	Format           Format
	TimeLayout       string
	Console          bool
	Level            Leveler
	ReplaceAttr      func(groups []string, attr Attr) Attr
	AddSource        bool
	LumberjackLogger *LumberjackLogger
	DevslogOptions   *DevslogOptions
	NoColor          bool
	Default          bool
}

func DefaultOptions() Options {
	return Options{
		Format:     FormatText,
		TimeLayout: DefaultTimeLayout,
		Level:      LevelInfo,
		DevslogOptions: &DevslogOptions{ // 预设默认配置
			HandlerOptions: &HandlerOptions{},
		},
	}
}

var (
	defaultOptions = DefaultOptions()
)

// WithOutputFile write log to some File
func WithOutputFile(file string) Option {
	return func(opt *Options) {
		opt.Output = file
	}
}

// WithFile write log to some File
// Deprecated: use WithOutputFile instead
func WithFile(file string) Option {
	return WithOutputFile(file)
}

// WithLumberjack write log to some File with rotation
func WithLumberjack(lumberjackLogger *LumberjackLogger) Option {
	return func(opt *Options) {
		opt.LumberjackLogger = lumberjackLogger
	}
}

// WithTimeLayout custom time format
func WithTimeLayout(timeLayout string) Option {
	return func(opt *Options) {
		opt.TimeLayout = timeLayout
	}
}

// WithConsole set the log to console or /dev/null
func WithConsole(set bool) Option {
	return func(opt *Options) {
		opt.Console = set
	}
}

// WithConsoleOnly set the log to console only
func WithConsoleOnly() Option {
	return func(opt *Options) {
		opt.Console = true
		opt.Output = ""
	}
}

// WithLevel custom log level
func WithLevel(level Leveler) Option {
	return func(opt *Options) {
		opt.Level = level
	}
}

// WithReplaceAttr custom replaceAttr
func WithReplaceAttr(replaceAttr func(groups []string, attr Attr) Attr) Option {
	return func(opt *Options) {
		opt.ReplaceAttr = replaceAttr
	}
}

// WithFormat custom format
func WithFormat(format Format) Option {
	return func(opt *Options) {
		opt.Format = format
	}
}

// WithAddSource add source info to log
func WithAddSource() Option {
	return func(opt *Options) {
		opt.AddSource = true
	}
}

// WithNoColor disable color
func WithNoColor() Option {
	return func(opt *Options) {
		opt.NoColor = true
	}
}

// WithDefault use output as slog.Default()
func WithDefault(set bool) Option {
	return func(opt *Options) {
		opt.Default = set
	}
}

// WithDevConfig set dev config
func WithDevConfig(config *DevslogOptions) Option {
	return func(opt *Options) {
		opt.DevslogOptions = config
	}
}
