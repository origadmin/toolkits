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
	DefaultTimeLayout = time.DateTime
)

const (
	LevelFatal = 12
)

// NewDebug create a new slog.Logger with debug level
func NewDebug(options ...Option) *Logger {
	return New(append(options, WithLevel(LevelDebug))...)
}

// New create a new slog.Logger
func New(options ...Option) *Logger {
	cfg := settings.ApplyDefault(defaultOptions, options)

	defaultLogger := Default()
	var outputs []io.Writer
	if cfg.Console {
		outputs = append(outputs, os.Stderr)
	}
	if cfg.Output != "" || cfg.LumberjackConfig != nil {
		pathname := cfg.Output
		if stat, err := os.Stat(pathname); err == nil && !stat.IsDir() {
			if err := os.Rename(pathname, backupLog(pathname)); err != nil {
				return defaultLogger
			}
		}

		if cfg.LumberjackConfig != nil {
			if cfg.LumberjackConfig.Filename != "" {
				pathname = cfg.LumberjackConfig.Filename
				if stat, err := os.Stat(pathname); err == nil && !stat.IsDir() {
					if err := os.Rename(pathname, backupLog(pathname)); err != nil {
						return defaultLogger
					}
				}
			} else {
				cfg.LumberjackConfig.Filename = pathname
			}
			outputs = append(outputs, cfg.LumberjackConfig)
		} else {
			if _, err := os.Stat(filepath.Dir(cfg.Output)); os.IsNotExist(err) {
				if err := os.Mkdir(cfg.Output, 0766); err != nil {
					return defaultLogger
				}
			}
			file, err := os.OpenFile(pathname, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
			if err != nil {
				return defaultLogger
			}
			outputs = append(outputs, file)
		}
	}

	multiOutput := io.Discard
	if len(outputs) > 0 {
		multiOutput = io.MultiWriter(outputs...)
	}
	//var handler Handler = NewTextHandler(output, &HandlerOptions{
	//	Level:       cfg.Level,
	//	ReplaceAttr: cfg.ReplaceAttr,
	//	AddSource:   cfg.AddSource,
	//})
	handler := createHandler(cfg, multiOutput)

	defaultLogger = slog.New(handler)
	if cfg.Default {
		slog.SetDefault(defaultLogger)
		slog.SetLogLoggerLevel(cfg.Level.Level())
	}

	return defaultLogger
}

func createHandler(opt *Options, output io.Writer) slog.Handler {
	switch opt.Format {
	case FormatJSON:
		handler := &HandlerOptions{
			Level:       opt.Level,
			ReplaceAttr: opt.ReplaceAttr,
			AddSource:   opt.AddSource,
		}
		return NewJSONHandler(output, handler)
	case FormatTint:
		return NewTintHandler(output, &TintOptions{
			AddSource:   opt.AddSource,
			Level:       opt.Level,
			ReplaceAttr: opt.ReplaceAttr,
			TimeFormat:  opt.TimeLayout,
			NoColor:     opt.NoColor,
		})
	case FormatDev:
		timeFormat := DefaultTimeLayout
		if opt.TimeLayout != "" {
			timeFormat = opt.TimeLayout
		}
		if opt.DevConfig != nil {
			if opt.DevConfig.HandlerOptions == nil {
				opt.DevConfig.HandlerOptions = &HandlerOptions{
					Level:       opt.Level,
					ReplaceAttr: opt.ReplaceAttr,
					AddSource:   opt.AddSource,
				}
			}
			if opt.DevConfig.TimeFormat == "" {
				opt.DevConfig.TimeFormat = timeFormat
			}
			if !opt.DevConfig.NoColor {
				opt.DevConfig.NoColor = opt.NoColor
			}
		}
		return NewDevSlogHandler(output, opt.DevConfig)
	default:
		handler := &HandlerOptions{
			Level:       opt.Level,
			ReplaceAttr: opt.ReplaceAttr,
			AddSource:   opt.AddSource,
		}
		return slog.NewTextHandler(output, handler)
	}
}

func backupLog(filename string) string {
	ext := filepath.Ext(filename)
	prefix := filename[:len(filename)-len(ext)]
	t := time.Now()
	timestamp := t.Format("20060102150405")
	return fmt.Sprintf("%s-%s%s", prefix, timestamp, ext)
}
