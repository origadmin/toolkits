/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package sloge implements the functions, types, and interfaces for the module.
package sloge

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/goexts/generic/settings"
)

const (
	// DefaultTimeLayout the default time layout;
	DefaultTimeLayout = time.RFC3339
)

const (
	LevelFatal = 12
)

type (
	LumberjackConfig = struct {
		// MaxSize is the maximum size in megabytes of the log file before it gets
		// rotated. It defaults to 100 megabytes.
		MaxSize int `json:"max_size" yaml:"max_size" toml:"max_size"`

		// MaxAge is the maximum number of days to retain old log files based on the
		// timestamp encoded in their filename.  Note that a day is defined as 24
		// hours and may not exactly correspond to calendar days due to daylight
		// savings, leap seconds, etc. The default is not to remove old log files
		// based on age.
		MaxAge int `json:"max_age" yaml:"max_age" toml:"max_age"`

		// MaxBackups is the maximum number of old log files to retain.  The default
		// is to retain all old log files (though MaxAge may still cause them to get
		// deleted.)
		MaxBackups int `json:"max_backups" yaml:"max_backups" toml:"max_backups"`

		// LocalTime determines if the time used for formatting the timestamps in
		// backup files is the computer's local time.  The default is to use UTC
		// time.
		LocalTime bool `json:"localtime" yaml:"localtime" toml:"localtime"`

		// Compress determines if the rotated log files should be compressed
		// using gzip. The default is not to perform compression.
		Compress bool `json:"compress" yaml:"compress" toml:"compress"`
	}

	DevConfig = struct {
		// Max number of printed elements in slice.
		MaxSlice uint `json:"max_slice" yaml:"max_slice" toml:"max_slice"`

		// If the attributes should be sorted by keys
		SortKeys bool `json:"sort_keys" yaml:"sort_keys" toml:"sort_keys"`

		// Add blank line after each log
		NewLine bool `json:"newline" yaml:"newline" toml:"newline"`

		// Indent \n in strings
		Indent bool `json:"indent" yaml:"indent" toml:"indent"`

		// Set color for Debug level, default: devslog.Blue
		DebugColor Color `json:"debug_color" yaml:"debug_color" toml:"debug_color"`

		// Set color for Info level, default: devslog.Green
		InfoColor Color `json:"info_color" yaml:"info_color" toml:"info_color"`

		// Set color for Warn level, default: devslog.Yellow
		WarnColor Color `json:"warn_color" yaml:"warn_color" toml:"warn_color"`

		// Set color for Error level, default: devslog.Red
		ErrorColor Color `json:"error_color" yaml:"error_color" toml:"error_color"`

		// Max stack trace frames when unwrapping errors
		MaxTrace uint `json:"max_trace" yaml:"max_trace" toml:"max_trace"`

		// Use method String() for formatting value
		Formatter bool `json:"formatter" yaml:"formatter" toml:"formatter"`
	}
)

// New create a new slog.Logger
func New(ss ...Setting) *Logger {
	opt := settings.Apply(defaultOption, ss)

	defaultLogger := Default()
	outputs := []io.Writer{os.Stderr}
	if !opt.Console {
		outputs = nil
	}
	if opt.OutputPath != "" {
		err := os.Mkdir(opt.OutputPath, 0766)
		if err != nil {
			return defaultLogger
		}
	}

	if opt.FileName != "" {
		pathname := filepath.Join(opt.OutputPath, opt.FileName)
		stat, err := os.Stat(pathname)
		if err == nil && !stat.IsDir() {
			err := os.Rename(pathname, backupLog(pathname))
			if err != nil {
				return defaultLogger
			}
		}

		if opt.LumberjackConfig != nil {
			outputs = append(outputs, &LumberjackLogger{
				Filename:   filepath.Join(opt.OutputPath, opt.FileName),
				MaxSize:    opt.LumberjackConfig.MaxSize,
				MaxAge:     opt.LumberjackConfig.MaxAge,
				MaxBackups: opt.LumberjackConfig.MaxBackups,
				LocalTime:  opt.LumberjackConfig.LocalTime,
				Compress:   opt.LumberjackConfig.Compress,
			})
		} else {
			file, err := os.OpenFile(pathname, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
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

	//var handler Handler = NewTextHandler(output, &HandlerOptions{
	//	Level:       opt.Level,
	//	ReplaceAttr: opt.ReplaceAttr,
	//	AddSource:   opt.AddSource,
	//})
	var handler Handler
	switch opt.Format {
	case FormatJSON:
		handler = NewJSONHandler(output, &HandlerOptions{
			Level:       opt.Level,
			ReplaceAttr: opt.ReplaceAttr,
			AddSource:   opt.AddSource,
		})
	case FormatTint:
		handler = NewTintHandler(output, &TintOptions{
			AddSource:   opt.AddSource,
			Level:       opt.Level,
			ReplaceAttr: opt.ReplaceAttr,
			TimeFormat:  opt.TimeLayout,
			NoColor:     opt.NoColor,
		})
	case FormatDev:
		handler = NewDevSlogHandler(output, &DevSlogOptions{
			HandlerOptions: &HandlerOptions{
				Level:       opt.Level,
				ReplaceAttr: opt.ReplaceAttr,
				AddSource:   opt.AddSource,
			},
			MaxSlicePrintSize:  opt.DevConfig.MaxSlice,
			SortKeys:           opt.DevConfig.SortKeys,
			TimeFormat:         opt.TimeLayout,
			NewLineAfterLog:    opt.DevConfig.NewLine,
			StringIndentation:  opt.DevConfig.Indent,
			DebugColor:         opt.DevConfig.DebugColor,
			InfoColor:          opt.DevConfig.InfoColor,
			WarnColor:          opt.DevConfig.WarnColor,
			ErrorColor:         opt.DevConfig.ErrorColor,
			MaxErrorStackTrace: opt.DevConfig.MaxTrace,
			StringerFormatter:  opt.DevConfig.Formatter,
			NoColor:            opt.NoColor,
		})
	default:
		handler = slog.NewTextHandler(output, &HandlerOptions{
			Level:       opt.Level,
			ReplaceAttr: opt.ReplaceAttr,
			AddSource:   opt.AddSource,
		})
	}

	defaultLogger = slog.New(handler)
	if opt.Default {
		slog.SetDefault(defaultLogger)
	}

	return defaultLogger
}

func backupLog(filename string) string {
	ext := filepath.Ext(filename)
	prefix := filename[:len(filename)-len(ext)]
	t := time.Now()
	timestamp := t.Format("20060102150405")
	return fmt.Sprintf("%s-%s%s", prefix, timestamp, ext)
}
