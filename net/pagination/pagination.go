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

type UnimplementedRequester struct{}

func (u UnimplementedRequester) GetCurrent() int32 {
	return 0
}

func (u UnimplementedRequester) GetPageSize() int32 {
	return 0
}

func (u UnimplementedRequester) GetPageToken() string {
	return ""
}

func (u UnimplementedRequester) GetOnlyCount() bool {
	return false
}

func (u UnimplementedRequester) GetNoPaging() bool {
	return false
}

func (u UnimplementedRequester) GetOrderBy() []string {
	return nil
}

func (u UnimplementedRequester) GetFieldMask() []string {
	return nil
}

type UnimplementedResponder struct{}

func (u UnimplementedResponder) GetSuccess() bool {
	return false
}

func (u UnimplementedResponder) GetTotal() int32 {
	return 0
}

func (u UnimplementedResponder) GetData() any {
	return nil
}

func (u UnimplementedResponder) GetCurrent() int32 {
	return 0
}

func (u UnimplementedResponder) GetPageSize() int32 {
	return 0
}

func (u UnimplementedResponder) GetNextPageToken() string {
	return ""
}

func (u UnimplementedResponder) GetExtra() any {
	return nil
}

var _ Requester = (*UnimplementedRequester)(nil)
var _ Responder = (*UnimplementedResponder)(nil)
