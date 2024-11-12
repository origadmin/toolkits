/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package database implements the functions, types, and interfaces for the module.
package database

import (
	"database/sql"

	"github.com/origadmin/toolkits/runtime/config"
)

func Open(database *config.Data_Database) (*sql.DB, error) {
	db, err := sql.Open(database.Dialect, database.Source)
	if err != nil {
		return nil, err
	}
	return db, nil
}
