/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package fileupload implements the functions, types, and interfaces for the module.
package fileupload

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// HTTPUploader implements the Uploader interface using HTTP
type HTTPUploader struct {
	url string
}

// NewHTTPUploader creates a new HTTPUploader instance
func NewHTTPUploader(url string) *HTTPUploader {
	return &HTTPUploader{url: url}
}

// SetFileHeader sets the file header after the file has been uploaded
func (u *HTTPUploader) SetFileHeader(ctx context.Context, header FileHeader) error {
	// Simulate setting file header
	fmt.Printf("Setting file header: %s\n", header.GetFilename())
	return nil
}

// UploadFile uploads the file first and then sets the header
func (u *HTTPUploader) UploadFile(ctx context.Context, rd io.Reader) error {
	// Simulate file upload
	resp, err := http.Post(u.url, "application/octet-stream", rd)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("upload failed with status code: %d", resp.StatusCode)
	}
	return nil
}

// Finalize finalizes the upload process
func (u *HTTPUploader) Finalize(ctx context.Context) (UploadResponse, error) {
	// Simulate finalizing upload
	return &UploadResponseImpl{
		Success:    true,
		Hash:       "dummy-hash",
		Path:       "dummy-path",
		Size:       0,
		FailReason: "",
	}, nil
}

// HTTPDownloader implements the Downloader interface using HTTP
type HTTPDownloader struct {
	url string
}

// NewHTTPDownloader creates a new HTTPDownloader instance
func NewHTTPDownloader(url string) *HTTPDownloader {
	return &HTTPDownloader{url: url}
}

// GetFileHeader retrieves the file header after the file has been downloaded
func (d *HTTPDownloader) GetFileHeader(ctx context.Context) (FileHeader, error) {
	// Simulate retrieving file header
	return &FileHeaderImpl{
		Filename:      "dummy-file",
		Size:          0,
		ModTime:       0,
		ModTimeString: "0",
		ContentType:   "application/octet-stream",
		Header:        make(map[string]string),
		IsDir:         false,
	}, nil
}

// DownloadFile downloads the file first and then retrieves the header
func (d *HTTPDownloader) DownloadFile(ctx context.Context) (io.Reader, error) {
	// Simulate file download
	resp, err := http.Get(d.url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("download failed with status code: %d", resp.StatusCode)
	}
	return resp.Body, nil
}

// Finalize finalizes the download process
func (d *HTTPDownloader) Finalize(ctx context.Context, resp UploadResponse) error {
	// Simulate finalizing download
	fmt.Printf("Finalizing download: %s\n", resp.GetPath())
	return nil
}

// UploadResponseImpl is a concrete implementation of UploadResponse
type UploadResponseImpl struct {
	Success    bool
	Hash       string
	Path       string
	Size       uint32
	FailReason string
}

func (r *UploadResponseImpl) GetSuccess() bool {
	return r.Success
}

func (r *UploadResponseImpl) GetHash() string {
	return r.Hash
}

func (r *UploadResponseImpl) GetPath() string {
	return r.Path
}

func (r *UploadResponseImpl) GetSize() uint32 {
	return r.Size
}

func (r *UploadResponseImpl) GetFailReason() string {
	return r.FailReason
}

// FileHeaderImpl is a concrete implementation of FileHeader
type FileHeaderImpl struct {
	Filename      string
	Size          uint32
	ModTime       uint32
	ModTimeString string
	ContentType   string
	Header        map[string]string
	IsDir         bool
}

func (h *FileHeaderImpl) GetFilename() string {
	return h.Filename
}

func (h *FileHeaderImpl) GetSize() uint32 {
	return h.Size
}

func (h *FileHeaderImpl) GetModTime() uint32 {
	return h.ModTime
}

func (h *FileHeaderImpl) GetModTimeString() string {
	return h.ModTimeString
}

func (h *FileHeaderImpl) GetContentType() string {
	return h.ContentType
}

func (h *FileHeaderImpl) GetHeader() map[string]string {
	return h.Header
}

func (h *FileHeaderImpl) GetIsDir() bool {
	return h.IsDir
}

func TestUploaderDownloaderSequence(t *testing.T) {
	// Create a test server for upload
	uploadTS := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		io.Copy(w, r.Body)
	}))
	defer uploadTS.Close()

	// Create a test server for download
	downloadTS := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		w.Write([]byte("Hello, World!"))
	}))
	defer downloadTS.Close()

	// Create an uploader
	uploader := NewHTTPUploader(uploadTS.URL)

	// Create a test file
	testFile := []byte("Hello, World!")
	reader := bytes.NewReader(testFile)

	// Upload the file
	err := uploader.UploadFile(context.Background(), reader)
	if err != nil {
		t.Errorf("UploadFile failed: %v", err)
	}

	// Finalize the upload
	resp, err := uploader.Finalize(context.Background())
	if err != nil {
		t.Errorf("Finalize failed: %v", err)
	}

	// Check the response
	if !resp.GetSuccess() {
		t.Errorf("Expected success, got failure")
	}

	// Create a downloader
	downloader := NewHTTPDownloader(downloadTS.URL)

	// Download the file
	downloadReader, err := downloader.DownloadFile(context.Background())
	if err != nil {
		t.Errorf("DownloadFile failed: %v", err)
	}
	defer downloadReader.(io.Closer).Close()

	// Read the downloaded content
	buf := new(bytes.Buffer)
	buf.ReadFrom(downloadReader)

	// Check the content
	if buf.String() != "Hello, World!" {
		t.Errorf("Expected 'Hello, World!', got %s", buf.String())
	}

	// Finalize the download
	err = downloader.Finalize(context.Background(), resp)
	if err != nil {
		t.Errorf("Finalize failed: %v", err)
	}
}
