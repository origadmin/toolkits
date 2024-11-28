/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package pagination implements the functions, types, and interfaces for the module.
package pagination

import (
	"github.com/goexts/generic/types"
)

const (
	DefaultPage     = 1
	DefaultPageSize = 15
	MaxPage         = 100
	MaxPageSize     = 1000
)

type PageLimiter interface {
	Current(in int32) int32
	PerPage(in int32) int32
}

type Limiter struct {
	MaxPage         int32
	DefaultPage     int32
	MaxPageSize     int32
	DefaultPageSize int32
}

func (obj Limiter) Current(in int32) int32 {
	in = types.ZeroOr(in, obj.DefaultPage)
	if in > obj.MaxPage {
		in = obj.DefaultPage
	}
	return in
}

func (obj Limiter) PerPage(in int32) int32 {
	in = types.ZeroOr(in, obj.DefaultPageSize)
	if in > (obj.MaxPageSize) {
		in = obj.MaxPageSize
	}
	return in
}

func NewLimiter(page, defaultPage, pageSize, defaultPageSize int32) Limiter {
	return Limiter{
		MaxPage:         page,
		DefaultPage:     defaultPage,
		MaxPageSize:     pageSize,
		DefaultPageSize: defaultPageSize,
	}
}

func DefaultLimiter() Limiter {
	return NewLimiter(MaxPage, DefaultPage, MaxPageSize, DefaultPageSize)
}

var _ PageLimiter = (*Limiter)(nil)
