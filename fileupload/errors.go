package fileupload

import "errors"

var (
	ErrNoFile         = errors.New("no file provided")
	ErrInvalidRequest = errors.New("invalid request")
	ErrUploadFailed   = errors.New("upload failed")
)
