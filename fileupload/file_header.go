/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package fileupload implements the functions, types, and interfaces for the module.
package fileupload

import (
	"io/fs"
	"mime/multipart"
	"time"
)

type FileInfo interface {
	fs.FileInfo
	ContentType() string
}

type httpFileHeader struct {
	Filename    string            `json:"filename"`
	Size        uint32            `json:"size"`
	Header      map[string]string `json:"header"`
	ContentType string            `json:"content_type"`
	ModTime     uint32            `json:"mod_time"`
	IsDir       bool              `json:"is_dir"`
}

func (h httpFileHeader) GetIsDir() bool {
	return h.IsDir
}

func (h httpFileHeader) GetFilename() string {
	return h.Filename
}

func (h httpFileHeader) GetSize() uint32 {
	return h.Size
}

func (h httpFileHeader) GetModTimeString() string {
	return time.Unix(int64(h.ModTime), 0).Format(ModTimeFormat)
}

func (h httpFileHeader) GetModTime() uint32 {
	return h.ModTime
}

func (h httpFileHeader) GetContentType() string {
	return h.ContentType
}

func (h httpFileHeader) GetHeader() map[string]string {
	return h.Header
}

type multipartFileInfo struct {
	header FileHeader
}

func (f *multipartFileInfo) ContentType() string {
	return f.header.GetContentType()
}

func (f *multipartFileInfo) Name() string {
	return f.header.GetFilename()
}

func (f *multipartFileInfo) Size() int64 {
	return int64(f.header.GetSize())
}

func (f *multipartFileInfo) Sys() any {
	return nil
}

func (f *multipartFileInfo) Mode() fs.FileMode {
	return 0o644
}

func (f *multipartFileInfo) ModTime() time.Time {
	return time.Unix(int64(f.header.GetModTime()), 0)
}

func (f *multipartFileInfo) IsDir() bool {
	return f.header.GetIsDir()
}

func ParseMultipart(header *multipart.FileHeader) FileInfo {
	headerStr := header.Header.Get("Last-Modified")
	mod, _ := time.Parse("Mon, 02 Jan 2006 15:04:05 MST", headerStr)
	requestHeader := make(map[string]string, len(header.Header))
	for k, v := range header.Header {
		requestHeader[k] = v[0]
	}
	fileHeader := &httpFileHeader{
		ContentType: header.Header.Get("Content-Type"),
		Filename:    header.Filename,
		Header:      requestHeader,
		ModTime:     uint32(mod.Unix()),
		Size:        uint32(header.Size),
		IsDir:       false,
	}
	return &multipartFileInfo{
		header: fileHeader,
	}
}

func ParseHeader(header FileHeader) FileInfo {
	return &multipartFileInfo{
		header: header,
	}
}

var _ fs.FileInfo = (*multipartFileInfo)(nil)
var _ FileInfo = (*multipartFileInfo)(nil)