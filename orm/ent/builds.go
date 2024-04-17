package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"

	"github.com/origadmin/toolkits/orm"
)

type Builder func(drv *sql.Driver, config orm.Config) (any, error)
type Before func(ctx context.Context, drv *sql.Driver) (err error)
