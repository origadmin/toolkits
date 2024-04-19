// Copyright (c) 2024 OrigAdmin. All rights reserved.

//go:build sqlite3 && !cgo

// Package database is the database client register package.
package database

import (
	_ "github.com/sqlite3ent/sqlite3"
)
