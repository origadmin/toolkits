/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package database implements the functions, types, and interfaces for the module.
package database

import (
	"context"
	"database/sql/driver"
)

type (
	// Tx is a transaction aliased to driver.Tx
	Tx = driver.Tx
	// ExecFunc is a function that can be executed within a transaction
	ExecFunc = func(context.Context) error
	// TxFunc is a function that can be executed within a transaction
	TxFunc = func(tx Tx) error
)

// Trans is a transaction wrapper
type Trans interface {
	Tx(ctx context.Context, fn ExecFunc) error
	InTx(ctx context.Context, fn TxFunc) error
}
