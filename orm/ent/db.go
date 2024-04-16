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

type Builder func(driver *sql.Driver, config orm.Config) (any, error)

func New(ctx context.Context, config orm.Config) (*orm.ORM, error) {
	var drv *sql.Driver
	var err error
	dialect := strings.ToLower(config.Dialect)
	switch dialect {
	case "mysql", "tidb":
		if err = helpers.CreateMySQLDatabase(config.DSN); err != nil {
			return nil, fmt.Errorf("failed to create database: %v", err)
		}
		dialect = "mysql"
	case "postgres", "cockroachdb":
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

	if config.Connector == nil {
		config.Connector = &Connect{
			driver: drv,
		}
	}

	db := orm.New(config)
	var name string
	for _, name = range config.Names {
		err = db.Open(name)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

type Connect struct {
	driver *sql.Driver
	once   *sync.Once
	build  func(driver *sql.Driver, config orm.Config) (any, error)
}

// Open builds ORM instances.
func (c *Connect) Open(config orm.Config) (any, error) {
	if c.once == nil && config.Once {
		c.once = new(sync.Once)
	}
	return c.build(c.driver, config)
}

func (c *Connect) Close() error {
	var err error
	if c.once != nil {
		c.once.Do(func() {
			err = c.driver.Close()
		})
	}
	return err
}

// Connector builds ORM instances.
func Connector(driver *sql.Driver, builds ...func(driver *sql.Driver, config orm.Config) (any, error)) *Connect {
	if len(builds) == 0 {
		builds = append(builds, builder)
	}
	return &Connect{
		driver: driver,
		build:  builds[0],
	}
}

func builder(driver *sql.Driver, config orm.Config) (any, error) {
	return nil, errors.New("not implemented")
}
