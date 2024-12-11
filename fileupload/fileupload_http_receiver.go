/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package fileupload implements the functions, types, and interfaces for the module.
package fileupload

import (
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"sync"
	"time"

	"github.com/goexts/generic/settings"
	"github.com/origadmin/toolkits/errors"
)

const (
	ModTimeFormat = "Mon, 02 Jan 2006 15:04:05 MST"
	bufSize       = 1024 * 64
)

type httpReceiver struct {
	builder  *uploadBuilder
	header   FileHeader
	file     multipart.File
	response http.ResponseWriter
	buf      []byte
	err      error
	//pr       *io.PipeReader
	pw *io.PipeWriter
}

type httpFileResponse struct {
	Success    bool   `json:"success"`
	Hash       string `json:"hash"`
	Path       string `json:"path"`
	Size       uint32 `json:"size"`
	FailReason string `json:"fail_reason"`
}

func (r *httpFileResponse) String() string {
	bytes, _ := json.MarshalIndent(r, "", "  ")
	return string(bytes)
}

func (r *httpFileResponse) GetSuccess() bool {
	if r == nil {
		return false
	}
	return r.Success
}

func (r *httpFileResponse) GetHash() string {
	return r.Hash
}

func (r *httpFileResponse) GetPath() string {
	return r.Path
}

func (r *httpFileResponse) GetSize() uint32 {
	return r.Size
}

func (r *httpFileResponse) GetFailReason() string {
	return r.FailReason
}

// GetFileHeader read the file header from the request.
func (d *httpReceiver) GetFileHeader(ctx context.Context) (FileHeader, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.header, nil
}

// ReceiveFile read the file data to the server with path.
func (d *httpReceiver) ReceiveFile(ctx context.Context) (io.ReadCloser, error) {
	if d.err != nil {
		return nil, d.err
	}
	if d.file == nil {
		d.err = ErrNoFile
		return nil, d.err
	}
	if d.buf == nil {
		d.buf = d.builder.NewBuffer()
	}
	var pr *io.PipeReader
	// Create a pipe to stream the data
	pr, d.pw = io.Pipe()

	// Start copying in a goroutine
	go func() {

		defer d.file.Close()

		for {
			select {
			case <-ctx.Done():
				d.pw.CloseWithError(ctx.Err())
				return
			default:
				n, err := d.file.Read(d.buf)
				if n > 0 {
					if _, werr := d.pw.Write(d.buf[:n]); werr != nil {
						d.pw.CloseWithError(werr)
						return
					}
				}
				if err == io.EOF {
					d.pw.Close()
					return
				}
				if err != nil {
					d.pw.CloseWithError(err)
					return
				}
			}
		}
	}()

	return pr, nil
}

// Finalize write the finalize status to the client and close the upload process.
func (d *httpReceiver) Finalize(ctx context.Context, resp UploadResponse) error {
	d.builder.Free(d.buf)
	if d.err != nil {
		return d.err
	}
	if resp.GetSuccess() {
		d.response.WriteHeader(http.StatusOK)
		// Write response headers
		d.response.Header().Set("Content-Type", "application/json")
		return json.NewEncoder(d.response).Encode(resp)
	}
	d.response.WriteHeader(http.StatusInternalServerError)
	return json.NewEncoder(d.response).Encode(map[string]string{
		"error": resp.GetFailReason(),
	})
}

func NewHTTPReceiver(req *http.Request, resp http.ResponseWriter, ss ...BuildSetting) Receiver {
	b := settings.Apply(&uploadBuilder{
		bufSize: bufSize,
	}, ss)
	b.bufPool = &sync.Pool{
		New: func() interface{} {
			return make([]byte, b.bufSize)
		},
	}
	// Read the file header from the request
	var receiver httpReceiver
	file, header, err := req.FormFile("file")
	if err != nil {
		receiver.err = errors.Wrap(err, "failed to get file from request")
		return &receiver
	}

	fileheader := make(map[string]string, len(header.Header))
	for k, v := range header.Header {
		fileheader[k] = v[0]
	}
	modTime := uint32(time.Now().Unix())
	if mod, err := time.Parse(ModTimeFormat, header.Header.Get("Last-Modified")); err == nil {
		modTime = uint32(mod.Unix())
	}

	// Create a FileHeader struct and populate it with the file information
	fileHeader := &httpFileHeader{
		Filename:    header.Filename,
		Size:        uint32(header.Size),
		ContentType: header.Header.Get("Content-Type"),
		Header:      fileheader,
		ModTime:     modTime,
	}

	return &httpReceiver{
		builder:  b,
		file:     file,
		response: resp,
		header:   fileHeader,
	}
}

var _ Receiver = &httpReceiver{}
