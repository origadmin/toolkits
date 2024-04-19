// Copyright (c) 2024 OrigAdmin. All rights reserved.

//go:build sqlite3 && cgo

// Package database is the database client wrapper
package database

import (
	_ "github.com/mattn/go-sqlite3"
)
