package orm

import "context"

type Config struct {
	Context      context.Context
	Debug        bool
	Once         bool                 // use once for close db,when connector used same driver
	Connector    Connector            // use for single database
	Connectors   map[string]Connector // use for multi database, if needed
	Dialect      string               // mysql/postgres/sqlite3/...
	ORM          string               // gorm,ent,xorm,...
	DSN          string
	MaxLifetime  int
	MaxIdleTime  int
	MaxOpenConns int
	MaxIdleConns int
	Names        []string
}
