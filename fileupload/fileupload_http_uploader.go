package fileupload

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type httpUploader struct {
	builder  Builder
	client   *http.Client
	request  *http.Request
	response *http.Response
	header   FileHeader
	uri      string
	buf      []byte
}

func (u *httpUploader) SetFileHeader(ctx context.Context, header FileHeader) error {
	u.header = header

	// Create new request
	req, err := http.NewRequestWithContext(ctx, "POST", u.uri, nil)
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("Content-Type", header.GetContentType())
	req.Header.Set("Content-Length", fmt.Sprintf("%d", header.GetSize()))
	for k, v := range header.GetHeader() {
		req.Header.Set(k, v)
	}

	u.request = req
	return nil
}

func (u *httpUploader) UploadFile(ctx context.Context, rd io.Reader) error {
	if u.request == nil {
		return ErrInvalidRequest
	}

	// Create pipe for request body
	pr, pw := io.Pipe()
	u.request.Body = pr
	if u.buf == nil {
		u.buf = u.builder.NewBuffer()
	}
	// Start uploading in background
	go func() {
		defer func() {
			u.builder.Free(u.buf)
			pw.Close()
		}()

		for {
			select {
			case <-ctx.Done():
				pw.CloseWithError(ctx.Err())
				return
			default:
				n, err := rd.Read(u.buf)
				if n > 0 {
					if _, werr := pw.Write(u.buf[:n]); werr != nil {
						pw.CloseWithError(werr)
						return
					}
				}
				if err == io.EOF {
					return
				}
				if err != nil {
					pw.CloseWithError(err)
					return
				}
			}
		}
	}()
	if u.client == nil {
		u.client = &http.Client{}
	}
	// Send request
	resp, err := u.client.Do(u.request)
	if err != nil {
		return err
	}
	u.response = resp
	return nil
}

func (u *httpUploader) Finalize(ctx context.Context) (UploadResponse, error) {
	var resp httpFileResponse
	decoder := json.NewDecoder(u.response.Body)
	if err := decoder.Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
