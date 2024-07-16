package sloge

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/origadmin/toolkits/setting"
)

const (
	// DefaultTimeLayout the default time layout;
	DefaultTimeLayout = time.RFC3339
)

// Setting custom setup config

type Option struct {
	File             io.Writer
	TimeLayout       string
	DisableConsole   bool
	Level            slog.Leveler
	LumberjackConfig *lumberjack.Logger
}

// WithFile write log to some File
func WithFile(file string) setting.Setting[Option] {
	dir := filepath.Dir(file)
	if err := os.MkdirAll(dir, 0766); err != nil {
		panic(err)
	}

	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0766)
	if err != nil {
		panic(err)
	}

	return func(opt *Option) {
		opt.File = f
	}
}

// WithLumberjack write log to some File with rotation
func WithLumberjack(log *lumberjack.Logger) setting.Setting[Option] {
	if log.Filename == "" {
		dir := filepath.Dir(log.Filename)
		if err := os.MkdirAll(dir, 0766); err != nil {
			panic(err)
		}
	}
	return func(opt *Option) {
		opt.File = log
	}
}

// WithTimeLayout custom time format
func WithTimeLayout(timeLayout string) setting.Setting[Option] {
	return func(opt *Option) {
		opt.TimeLayout = timeLayout
	}
}

// WithDisableConsole WithEnableConsole write log to os.Stdout or os.Stderr
func WithDisableConsole() setting.Setting[Option] {
	return func(opt *Option) {
		opt.DisableConsole = true
	}
}

// New create a new slog.Logger
func New(settings ...setting.Setting[Option]) *slog.Logger {
	opt := setting.Apply(Option{
		Level:      slog.LevelDebug,
		TimeLayout: DefaultTimeLayout,
		File:       os.Stdout,
	}, settings...)
	return slog.New(slog.NewJSONHandler(opt.File, &slog.HandlerOptions{
		Level:       opt.Level,
		ReplaceAttr: nil,
		AddSource:   true,
	}))
}
