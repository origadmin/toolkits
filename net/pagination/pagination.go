/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package pagination implements the functions, types, and interfaces for the module.
package pagination

// Requester is the request interface for the module.
type Requester interface {
	GetCurrent() int32
	GetPageSize() int32
	GetPageToken() string
	GetOnlyCount() bool
	GetNoPaging() bool
	GetOrderBy() []string
	GetFieldMask() []string
}

// Responder is the response interface for the module.
type Responder interface {
	GetSuccess() bool
	GetTotal() int32
	GetData() any
	GetCurrent() int32
	GetPageSize() int32
	GetNextPageToken() string
	GetExtra() any
}
