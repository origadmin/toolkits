package ent

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"entgo.io/ent/dialect/sql"

	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/orm"
	"github.com/origadmin/toolkits/orm/internal/helpers"
)

type DB struct {
	config orm.Config
	once   *sync.Once
	drv    *sql.Driver
}

func Open(ctx context.Context, config orm.Config) (*DB, error) {
	if config.Context == nil {
		config.Context = context.Background()
	}
	var drv *sql.Driver
	var err error
	dialect := strings.ToLower(config.Dialect)
	switch dialect {
	case "mysql", "tidb":
		if err = helpers.CreateMySQLDatabase(config.DSN); err != nil {
			return nil, fmt.Errorf("failed to create database: %v", err)
		}
		dialect = "mysql"
	case "postgres":
	case "sqlite3":
		_ = os.MkdirAll(filepath.Dir(config.DSN), os.ModePerm)
	default:
		err = fmt.Errorf("unsupported database type: %s", config.Dialect)
	}
	drv, err = sql.Open(dialect, config.DSN)

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	sqlDB := drv.DB()
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.MaxLifetime) * time.Second)
	sqlDB.SetConnMaxIdleTime(time.Duration(config.MaxIdleTime) * time.Second)

	var once *sync.Once
	if config.Once {
		once = new(sync.Once)
	}

	db := &DB{
		config: config,
		once:   once,
		drv:    drv,
	}
	return db, nil
}

func (db *DB) Close() error {
	var err error
	if db.once != nil {
		db.once.Do(func() {
			err = db.drv.Close()
		})
	}
	return err
}

func (db *DB) Before(ctx context.Context, funcs ...Before) error {
	for _, fn := range funcs {
		err := fn(ctx, db.drv)
		if err != nil {
			return err
		}
	}
	return nil
}

type Connect struct {
	db      *DB
	builder func(driver *sql.Driver, config orm.Config) (any, error)
}

// Open builds DB instances.
func (c *Connect) Open(config orm.Config) (any, error) {
	return c.builder(c.db.drv, config)
}

func (c *Connect) Close() error {
	return c.db.Close()
}

// Connector builds DB instances.
func Connector(db *DB, builder Builder) *Connect {
	return &Connect{
		db:      db,
		builder: builder,
	}
}

func builder(driver *sql.Driver, config orm.Config) (any, error) {
	return nil, errors.New("not implemented")
}

func before(ctx context.Context, drv *sql.Driver) (err error) {
	return
}
