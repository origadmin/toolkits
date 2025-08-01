/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package slogx implements enhanced logging functions for slog
package slogx

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
	if cfg.LumberjackConfig != nil {
		outputs = append(outputs, cfg.LumberjackConfig)
	} else if cfg.Output != "" {
		pathname := cfg.Output
		if stat, err := os.Stat(pathname); err == nil && !stat.IsDir() {
			if err := os.Rename(pathname, backupLog(pathname)); err != nil {
				return defaultLogger
			}
		}

		dir := filepath.Dir(cfg.Output)
		if err := os.MkdirAll(dir, 0766); err != nil {
			return defaultLogger
		}
		file, err := os.OpenFile(pathname, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
		if err != nil {
			return defaultLogger
		}
		outputs = append(outputs, file)
	}

	multiOutput := io.Discard
	if len(outputs) > 0 {
		multiOutput = io.MultiWriter(outputs...)
	}
	handler := createHandler(cfg, multiOutput)
	defaultLogger = slog.New(handler)
	if cfg.Default {
		slog.SetDefault(defaultLogger)
	}
	return defaultLogger
}

func createHandler(opt *Options, output io.Writer) slog.Handler {
	switch opt.Format {
	case FormatDev:
		timeFormat := DefaultTimeLayout
		if opt.TimeLayout != "" {
			timeFormat = opt.TimeLayout
		}
		if opt.DevConfig == nil {
			opt.DevConfig = &DevConfig{
				HandlerOptions: &HandlerOptions{
					Level:       opt.Level,
					ReplaceAttr: opt.ReplaceAttr,
					AddSource:   opt.AddSource,
				},
				TimeFormat: timeFormat,
				NoColor:    opt.NoColor,
			}
		} else {
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
	case FormatTint:
		handler := &TintOptions{
			Level:       opt.Level,
			ReplaceAttr: opt.ReplaceAttr,
			AddSource:   opt.AddSource,
			NoColor:     opt.NoColor,
		}
		return NewTintHandler(output, handler)
	case FormatJSON:
		handler := &slog.HandlerOptions{
			Level:       opt.Level,
			ReplaceAttr: opt.ReplaceAttr,
			AddSource:   opt.AddSource,
		}
		return NewJSONHandler(output, handler)
	default:
		handler := &slog.HandlerOptions{
			Level: opt.Level,
		}
		return NewTextHandler(output, handler)
	}
}

func backupLog(filename string) string {
	ext := filepath.Ext(filename)
	prefix := filename[:len(filename)-len(ext)]
	t := time.Now()
	timestamp := t.Format("20060102150405")
	return fmt.Sprintf("%s-%s%s", prefix, timestamp, ext)
}
