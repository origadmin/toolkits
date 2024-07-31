package sloge

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/lmittmann/tint"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/goexts/ggb/settings"
)

const (
	// DefaultTimeLayout the default time layout;
	DefaultTimeLayout = time.RFC3339
)

type Format int

const (
	// JSONFormat json format
	JSONFormat Format = iota
	// TextFormat text format
	TextFormat
	// TintFormat tint format
	TintFormat
)

type tintOption struct {
	Tint *tint.Options
}

// Option custom setup config
type Option struct {
	Files            []io.Writer
	LogFormat        Format
	TimeLayout       string
	DisableConsole   bool
	Level            slog.Leveler
	ReplaceAttr      func(groups []string, a slog.Attr) slog.Attr
	AddSource        bool
	LumberjackConfig *lumberjack.Logger
	NoColor          bool
	*tintOption
}

// WithFile write log to some File
func WithFile(file string) settings.Setting[Option] {
	dir := filepath.Dir(file)
	if err := os.MkdirAll(dir, 0766); err != nil {
		panic(err)
	}

	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0766)
	if err != nil {
		panic(err)
	}

	return func(opt *Option) {
		opt.Files = append(opt.Files, f)
	}
}

// WithLumberjack write log to some File with rotation
func WithLumberjack(log *lumberjack.Logger) settings.Setting[Option] {
	if log.Filename == "" {
		dir := filepath.Dir(log.Filename)
		if err := os.MkdirAll(dir, 0766); err != nil {
			panic(err)
		}
	}
	return func(opt *Option) {
		opt.Files = append(opt.Files, log)
	}
}

// WithTimeLayout custom time format
func WithTimeLayout(timeLayout string) settings.Setting[Option] {
	return func(opt *Option) {
		opt.TimeLayout = timeLayout
	}
}

// WithDisableConsole WithEnableConsole write log to os.Stdout or os.Stderr
func WithDisableConsole() settings.Setting[Option] {
	return func(opt *Option) {
		opt.DisableConsole = true
	}
}

// New create a new slog.Logger
func New(ss ...settings.Setting[Option]) *slog.Logger {
	opt := settings.Apply(&Option{
		Files:            []io.Writer{os.Stdout},
		LogFormat:        TextFormat,
		TimeLayout:       DefaultTimeLayout,
		DisableConsole:   false,
		Level:            slog.LevelDebug,
		ReplaceAttr:      nil,
		AddSource:        true,
		LumberjackConfig: nil,
		NoColor:          false,
		tintOption:       nil,
	}, ss)

	if opt.DisableConsole {
		opt.Files = nil
	}

	if opt.LumberjackConfig != nil {
		opt.Files = append(opt.Files, opt.LumberjackConfig)
	}

	var output io.Writer
	fileLen := len(opt.Files)
	switch {
	case fileLen == 1:
		output = opt.Files[0]
	case fileLen > 1:
		output = io.MultiWriter(opt.Files...)
	default:
		output = io.Discard
	}

	var handler slog.Handler = slog.NewTextHandler(output, &slog.HandlerOptions{
		Level:       opt.Level,
		ReplaceAttr: opt.ReplaceAttr,
		AddSource:   opt.AddSource,
	})
	switch opt.LogFormat {
	case JSONFormat:
		handler = slog.NewJSONHandler(output, &slog.HandlerOptions{
			Level:       opt.Level,
			ReplaceAttr: opt.ReplaceAttr,
			AddSource:   opt.AddSource,
		})
	// case TextFormat:
	case TintFormat:
		handler = tint.NewHandler(output, &tint.Options{
			AddSource:   opt.AddSource,
			Level:       opt.Level,
			ReplaceAttr: opt.ReplaceAttr,
			TimeFormat:  opt.TimeLayout,
			NoColor:     opt.NoColor,
		})
	default:
		handler = slog.NewTextHandler(output, &slog.HandlerOptions{
			Level:       opt.Level,
			ReplaceAttr: opt.ReplaceAttr,
			AddSource:   opt.AddSource,
		})
	}

	return slog.New(handler)
}
