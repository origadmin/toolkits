/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package filter implements the functions, types, and interfaces for the module.
package filter

import (
	"net/http"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/goexts/generic/configure"
)

type bloomFilter struct {
	delimiter string
	AllowOff  bool
	allows    [MethodTypeMax]*bloom.BloomFilter
	DenyOff   bool
	denies    [MethodTypeMax]*bloom.BloomFilter
	skipAny   bool
}

func (bf *bloomFilter) Skip(req *http.Request) bool {
	path := req.URL.Path
	method := req.Method
	return bf.Denied(method, path) || !bf.Allowed(method, path)
}

func (bf *bloomFilter) Denied(method, path string) bool {
	if bf.DenyOff {
		return false
	}
	idx := MethodIndex(method)
	if idx >= MethodTypeMax {
		return false
	}
	filter := bf.denies[idx]
	if filter.Test([]byte(path)) {
		return true
	}
	if bf.skipAny && (idx == MethodAny) {
		return false
	}
	return bf.denies[idx].Test([]byte(path))
}

func (bf *bloomFilter) Allowed(method, path string) bool {
	if bf.AllowOff {
		return true
	}
	idx := MethodIndex(method)
	if idx >= MethodTypeMax {
		return false
	}
	filter := bf.allows[idx]
	if filter.Test([]byte(path)) {
		return true
	}
	if bf.skipAny || idx == MethodAny {
		return false
	}
	return bf.allows[MethodAny].Test([]byte(path))
}

func (bf *bloomFilter) AllowStrings(method string) []string {
	//TODO implement me
	panic("implement me")
}

func (bf *bloomFilter) DenyStrings(method string) []string {
	//TODO implement me
	panic("implement me")
}

func NewBloomFilter(ss ...Setting) Filter {
	o := settings.Apply(&Option{
		Delimiter:         DefaultDelimiter,
		SkipAny:           false,
		DenyOff:           false,
		AllowOff:          false,
		ExpectedElements:  10000,
		FalsePositiveRate: 0.01,
	}, ss)
	bf := &bloomFilter{
		delimiter: o.Delimiter,
	}
	for i := MethodType(0); i < MethodTypeMax; i++ {
		bf.allows[i] = bloom.NewWithEstimates(o.ExpectedElements, o.FalsePositiveRate)
		bf.denies[i] = bloom.NewWithEstimates(o.ExpectedElements, o.FalsePositiveRate)
	}

	for _, s := range o.Allows {
		if s == "" {
			continue
		}
		methods, path := splitURI(s, o.Delimiter)
		for i := range methods {
			idx := MethodIndex(methods[i])
			if bf.skipAny && idx == MethodAny {
				continue
			}
			bf.allows[idx].Add([]byte(path))
		}
		bf.AllowOff = false
	}

	for _, s := range o.Denies {
		if s == "" {
			continue
		}
		methods, path := splitURI(s, o.Delimiter)
		for i := range methods {
			idx := MethodIndex(methods[i])
			if bf.skipAny && idx == MethodAny {
				continue
			}
			bf.denies[idx].Add([]byte(path))
		}
		bf.DenyOff = false
	}

	return bf
}

func (bf *bloomFilter) AddAllow(method MethodType, link string) {
	if !bf.AllowOff {
		bf.allows[method].Add([]byte(link))
	}
}

func (bf *bloomFilter) AddDeny(method MethodType, link string) {
	if !bf.DenyOff {
		bf.denies[method].Add([]byte(link))
	}
}

func (bf *bloomFilter) CheckAllow(method MethodType, link string) bool {
	if bf.AllowOff {
		return true
	}
	return bf.allows[method].Test([]byte(link))
}

func (bf *bloomFilter) CheckDeny(method MethodType, link string) bool {
	if bf.DenyOff {
		return false
	}
	return bf.denies[method].Test([]byte(link))
}
