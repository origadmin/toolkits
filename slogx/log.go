/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package slogx contains the go:generate directives for adptool.
// This file exists to ensure the generated output is named log.adapter.go.
package slogx

//go:generate adptool ./log.go
//go:adapter:package gopkg.in/natefinch/lumberjack.v2 lumberjack
//go:adapter:package:type Logger
//go:adapter:package:type:rename LumberjackLogger
//go:adapter:package github.com/lmittmann/tint tint
//go:adapter:package:type Options
//go:adapter:package:type:rename TintOptions
//go:adapter:package:func NewHandler
//go:adapter:package:func:rename NewTintHandler
//go:adapter:package:func *
//go:adapter:package:func:prefix Tint
//go:adapter:package github.com/golang-cz/devslog devslog
//go:adapter:package:type Options
//go:adapter:package:type:rename DevslogOptions
//go:adapter:package log/slog slog
//go:adapter:package:func New
//go:adapter:package:func:rename NewSlog
