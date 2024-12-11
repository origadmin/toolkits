package fileupload

import (
	"github.com/origadmin/toolkits/errors"
)

var (
	ErrNoFile         = errors.String("no file provided")
	ErrInvalidRequest = errors.String("invalid request")
	ErrUploadFailed   = errors.String("upload failed")
)
