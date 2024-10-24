// Copyright (c) 2024 OrigAdmin. All rights reserved.

//go:build sqlite3 && !cgo

// Package support is the database client register package.
package support

import (
	_ "github.com/sqlite3ent/sqlite3"
)

type SQLite3Go struct{}
