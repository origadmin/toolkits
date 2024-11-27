/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package pagination implements the functions, types, and interfaces for the module.
package pagination

import (
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Requester is the request interface for the module.
type Requester interface {
	GetCurrent() int32
	GetPageSize() int32
	GetPageToken() string
	GetOnlyCount() bool
	GetNoPaging() bool
	GetOrderBy() []string
	GetFieldMask() *fieldmaskpb.FieldMask
}

// Responder is the response interface for the module.
type Responder interface {
	GetSuccess() bool
	GetTotal() int32
	GetData() *anypb.Any
	GetCurrent() int32
	GetPageSize() int32
	GetNextPageToken() string
	GetExtra() *anypb.Any
}
