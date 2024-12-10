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
)

const (
	ModTimeFormat = "Mon, 02 Jan 2006 15:04:05 MST"
	bufSize       = 1024 * 64
)

type httpDownloader struct {
	builder  Builder
	header   FileHeader
	file     multipart.File
	response http.ResponseWriter
	buf      []byte
}

type httpFileResponse struct {
	Success    bool   `json:"success"`
	Hash       string `json:"hash"`
	Path       string `json:"path"`
	Size       uint32 `json:"size"`
	FailReason string `json:"fail_reason"`
}

func (r *httpFileResponse) GetSuccess() bool {
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
func (d *httpDownloader) GetFileHeader(ctx context.Context) (FileHeader, error) {
	return d.header, nil
}

// DownloadFile read the file data to the server with path.
func (d *httpDownloader) DownloadFile(ctx context.Context) (io.Reader, error) {
	if d.file == nil {
		return nil, ErrNoFile
	}
	if d.buf == nil {
		d.buf = d.builder.NewBuffer()
	}

	// Create a pipe to stream the data
	pr, pw := io.Pipe()

	// Start copying in a goroutine
	go func() {

		defer d.file.Close()

		for {
			select {
			case <-ctx.Done():
				pw.CloseWithError(ctx.Err())
				return
			default:
				n, err := d.file.Read(d.buf)
				if n > 0 {
					if _, werr := pw.Write(d.buf[:n]); werr != nil {
						pw.CloseWithError(werr)
						return
					}
				}
				if err == io.EOF {
					pw.Close()
					return
				}
				if err != nil {
					pw.CloseWithError(err)
					return
				}
			}
		}
	}()

	return pr, nil
}

// Finalize write the finalize status to the client and close the upload process.
func (d *httpDownloader) Finalize(ctx context.Context, resp UploadResponse) error {
	d.builder.Free(d.buf)
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

var _ Downloader = &httpDownloader{}
