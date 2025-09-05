/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package slogx

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-cz/devslog"
	"github.com/lmittmann/tint"
)

const (
	// DefaultTimeLayout the default time layout;
	DefaultTimeLayout = time.DateTime
)

// LevelFatal is a custom log level, demonstrating how to extend slog's levels.
const (
	LevelFatal = slog.Level(12)
)

// NewDebug creates a new slog.Logger with debug level.
// It's a convenient shortcut for New(WithLevel(slog.LevelDebug)).
func NewDebug(options ...Option) *slog.Logger {
	return New(append(options, WithLevel(slog.LevelDebug))...)
}

// New creates a new slog.Logger by applying the provided functional options.
// This function is a refactored, cleaner version of the original NewSlogx,
// preserving all functionality while improving maintainability.
func New(options ...Option) *slog.Logger {
	// Apply functional options to the default configuration.
	// This replaces the dependency on github.com/goexts/generic/configure.
	cfg := defaultOptions
	for _, apply := range options {
		apply(&cfg)
	}

	// Setup writers based on the configuration.
	var writers []io.Writer
	if cfg.Console {
		writers = append(writers, os.Stderr)
	}
	if cfg.LumberjackLogger != nil {
		writers = append(writers, cfg.LumberjackLogger)
	}
	// Ensure we don't write to the same file twice if lumberjack is also configured for it.
	if cfg.Output != "" {
		// Use the original backup logic for plain file output.
		if _, err := os.Stat(cfg.Output); err == nil {
			_ = os.Rename(cfg.Output, backupLog(cfg.Output))
		}
		if err := os.MkdirAll(filepath.Dir(cfg.Output), 0766); err == nil {
			if file, err := os.OpenFile(cfg.Output, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666); err == nil {
				writers = append(writers, file)
			}
		}
	}

	var output io.Writer = io.Discard
	if len(writers) > 0 {
		output = io.MultiWriter(writers...)
	}

	// Create the appropriate slog.Handler using a clean switch statement.
	var handler slog.Handler
	handlerOpts := &slog.HandlerOptions{
		Level:       cfg.Level,
		AddSource:   cfg.AddSource,
		ReplaceAttr: cfg.ReplaceAttr,
	}

	switch cfg.Format {
	case FormatTint:
		// The adapter provides TintOptions as an alias for tint.Options
		handler = tint.NewHandler(output, &TintOptions{
			Level:       handlerOpts.Level,
			AddSource:   handlerOpts.AddSource,
			ReplaceAttr: handlerOpts.ReplaceAttr,
			NoColor:     cfg.NoColor,
		})

	case FormatDev:
		// The adapter provides DevslogOptions as an alias for devslog.Options
		devOpts := cfg.DevslogOptions
		if devOpts == nil {
			devOpts = &DevslogOptions{}
		}
		// Ensure the underlying handler options are populated.
		if devOpts.HandlerOptions == nil {
			devOpts.HandlerOptions = &slog.HandlerOptions{}
		}
		devOpts.HandlerOptions.Level = handlerOpts.Level
		devOpts.HandlerOptions.AddSource = handlerOpts.AddSource
		devOpts.HandlerOptions.ReplaceAttr = handlerOpts.ReplaceAttr
		if devOpts.TimeFormat == "" {
			devOpts.TimeFormat = cfg.TimeLayout
		}
		if !devOpts.NoColor {
			devOpts.NoColor = cfg.NoColor
		}
		handler = devslog.NewHandler(output, devOpts)

	case FormatJSON:
		handler = slog.NewJSONHandler(output, handlerOpts)

	case FormatText:
		fallthrough
	default:
		handler = slog.NewTextHandler(output, handlerOpts)
	}

	// Create the logger and set it as default if requested.
	logger := NewSlog(handler)
	if cfg.Default {
		slog.SetDefault(logger)
	}
	return logger
}

// backupLog creates a backup filename with a timestamp.
func backupLog(filename string) string {
	ext := filepath.Ext(filename)
	prefix := filename[:len(filename)-len(ext)]
	timestamp := time.Now().Format("20060102-150405")
	return fmt.Sprintf("%s-%s%s", prefix, timestamp, ext)
}
