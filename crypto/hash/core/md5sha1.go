/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package core implements the functions, types, and interfaces for the module.
package core

import (
	"crypto/md5"
	"crypto/sha1"
	"hash"
)

type md5sha1 struct {
	md5  hash.Hash
	sha1 hash.Hash
}

func (h md5sha1) Write(p []byte) (n int, err error) {
	h.md5.Write(p)
	return h.sha1.Write(p)
}

func (h md5sha1) Sum(b []byte) []byte {
	md5sum := h.md5.Sum(nil)
	sha1sum := h.sha1.Sum(nil)
	return append(append(b, md5sum...), sha1sum...)
}

func (h md5sha1) Reset() {
	h.md5.Reset()
	h.sha1.Reset()
}

func (h md5sha1) Size() int {
	return h.md5.Size() + h.sha1.Size()
}

func (h md5sha1) BlockSize() int {
	if h.md5.BlockSize() > h.sha1.BlockSize() {
		return h.md5.BlockSize()
	}
	return h.sha1.BlockSize()
}

func NewMD5SHA1() hash.Hash {
	return &md5sha1{
		md5:  md5.New(),
		sha1: sha1.New(),
	}
}
