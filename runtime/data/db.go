/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package database implements the functions, types, and interfaces for the module.
package data

import (
	"database/sql"

	"github.com/origadmin/toolkits/errors"
	configv1 "github.com/origadmin/toolkits/runtime/gen/go/config/v1"
)

const (
	ErrConfigNil = errors.String("database: config is nil")
)

func Open(data *configv1.Data) (*sql.DB, error) {
	database := data.GetDatabase()
	if database == nil {
		return nil, ErrConfigNil
	}

	db, err := sql.Open(database.Driver, database.Source)
	if err != nil {
		return nil, err
	}
	return db, nil
}
