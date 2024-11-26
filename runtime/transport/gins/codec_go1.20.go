//go:build go1.20

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package gins

import "net/http"

// ResponseController is type net/http.ResponseController which was added in Go 1.20.
type ResponseController = http.ResponseController
