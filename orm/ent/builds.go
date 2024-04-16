package ent

import (
	"entgo.io/ent/dialect/sql"

	"github.com/origadmin/toolkits/orm"
)

type Builder func(drv *sql.Driver, config orm.Config) (any, error)
