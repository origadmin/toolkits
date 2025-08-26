/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package errors provides a means to return detailed information
// for a request error, typically encoded in JSON format.
// This package wraps standard library error handling features and
// offers additional stack trace capabilities.
package errors

import (
	_ "errors"

	_ "github.com/go-kratos/kratos/v2/errors"
	_ "github.com/hashicorp/go-multierror"
	_ "github.com/pkg/errors"
)

//go:generate adptool .
//go:adapter:package github.com/go-kratos/kratos/v2/errors kerrors
//go:adapter:package:func *
//go:adapter:package:func:prefix Kratos
//go:adapter:package github.com/pkg/errors perr
//go:adapter:package:func *
//go:adapter:package:func:prefix Pkg
//go:adapter:package github.com/hashicorp/go-multierror merr
//go:adapter:package:type *
//go:adapter:package:type:prefix Multi
//go:adapter:package errors stderr
