package loge

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/goexts/ggb/settings"
	"github.com/lmittmann/tint"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	// DefaultTimeLayout the default time layout;
	DefaultTimeLayout = time.RFC3339
)

type (
	Handler          = slog.Handler
	Level            = slog.Level
	Attr             = slog.Attr
	Leveler          = slog.Leveler
	HandlerOptions   = slog.HandlerOptions
	TintOptions      = tint.Options
	LumberjackConfig = struct {
		// MaxSize is the maximum size in megabytes of the log file before it gets
		// rotated. It defaults to 100 megabytes.
		MaxSize int `json:"maxsize" yaml:"maxsize" toml:"maxsize"`

		// MaxAge is the maximum number of days to retain old log files based on the
		// timestamp encoded in their filename.  Note that a day is defined as 24
		// hours and may not exactly correspond to calendar days due to daylight
		// savings, leap seconds, etc. The default is not to remove old log files
		// based on age.
		MaxAge int `json:"maxage" yaml:"maxage" toml:"maxage"`

		// MaxBackups is the maximum number of old log files to retain.  The default
		// is to retain all old log files (though MaxAge may still cause them to get
		// deleted.)
		MaxBackups int `json:"maxbackups" yaml:"maxbackups" toml:"maxbackups"`

		// LocalTime determines if the time used for formatting the timestamps in
		// backup files is the computer's local time.  The default is to use UTC
		// time.
		LocalTime bool `json:"localtime" yaml:"localtime" toml:"localtime"`

		// Compress determines if the rotated log files should be compressed
		// using gzip. The default is not to perform compression.
		Compress bool `json:"compress" yaml:"compress" toml:"compress"`
	}
)

const (
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
	LevelFatal = 12
)

//go:generate stringer -type=Format -trimprefix=Format
type Format int

const (
	// FormatJSON json format
	FormatJSON Format = iota
	// FormatText text format
	FormatText
	// FormatTint tint format
	FormatTint
)

// New create a new slog.Logger
func New(ss ...Setting) *slog.Logger {
	opt := settings.Apply(defaultOption, ss)
	defaultLogger := slog.Default()
	outputs := []io.Writer{os.Stderr}
	if opt.DisableConsole {
		outputs = nil
	}
	if opt.OutputPath != "" {
		err := os.Mkdir(opt.OutputPath, 0766)
		if err != nil {
			return defaultLogger
		}
	}

	if opt.FileName != "" {
		pathfile := filepath.Join(opt.OutputPath, opt.FileName)
		stat, err := os.Stat(pathfile)
		if err == nil && !stat.IsDir() {
			err := os.Rename(pathfile, backupLog(pathfile))
			if err != nil {
				return defaultLogger
			}
		}

		if opt.LumberjackConfig != nil {
			outputs = append(outputs, &lumberjack.Logger{
				Filename:   filepath.Join(opt.OutputPath, opt.FileName),
				MaxSize:    opt.LumberjackConfig.MaxSize,
				MaxAge:     opt.LumberjackConfig.MaxAge,
				MaxBackups: opt.LumberjackConfig.MaxBackups,
				LocalTime:  opt.LumberjackConfig.LocalTime,
				Compress:   opt.LumberjackConfig.Compress,
			})
		} else {
			file, err := os.OpenFile(pathfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
			if err != nil {
				return defaultLogger
			}
			outputs = append(outputs, file)
		}
	}

	var output io.Writer
	fileLen := len(outputs)
	switch {
	case fileLen == 1:
		output = outputs[0]
	case fileLen > 1:
		output = io.MultiWriter(outputs...)
	default:
		output = io.Discard
	}

	var handler Handler = slog.NewTextHandler(output, &HandlerOptions{
		Level:       opt.Level,
		ReplaceAttr: opt.ReplaceAttr,
		AddSource:   opt.AddSource,
	})
	switch opt.Format {
	case FormatJSON:
		handler = slog.NewJSONHandler(output, &HandlerOptions{
			Level:       opt.Level,
			ReplaceAttr: opt.ReplaceAttr,
			AddSource:   opt.AddSource,
		})
	case FormatTint:
		handler = tint.NewHandler(output, &TintOptions{
			AddSource:   opt.AddSource,
			Level:       opt.Level,
			ReplaceAttr: opt.ReplaceAttr,
			TimeFormat:  opt.TimeLayout,
			NoColor:     opt.NoColor,
		})
	default:
		handler = slog.NewTextHandler(output, &HandlerOptions{
			Level:       opt.Level,
			ReplaceAttr: opt.ReplaceAttr,
			AddSource:   opt.AddSource,
		})
	}

	defaultLogger = slog.New(handler)
	slog.SetDefault(defaultLogger)
	return defaultLogger
}

func backupLog(filename string) string {
	ext := filepath.Ext(filename)
	prefix := filename[:len(filename)-len(ext)]
	t := time.Now()
	timestamp := t.Format("20060102150405")
	return fmt.Sprintf("%s-%s%s", prefix, timestamp, ext)
}
