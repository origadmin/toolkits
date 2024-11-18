package database

import (
	"database/sql/driver"

	"github.com/origadmin/toolkits/context"
)

// Tx is a transaction aliased to driver.Tx
type Tx = driver.Tx

// ExecFunc is a function that can be executed within a transaction
type ExecFunc = func(context.Context) error

// TxFunc is a function that can be executed within a transaction
type TxFunc = func(tx Tx) error

// Trans is a transaction wrapper
type Trans = interface {
	Tx(ctx context.Context, fn ExecFunc) error
	InTx(ctx context.Context, fn TxFunc) error
}

type transCtx struct{}

func FromContext(ctx context.Context) Trans {
	if trans, ok := ctx.Value(transCtx{}).(Trans); ok {
		return trans
	}
	return nil
}

func NewTrans(ctx context.Context, trans Trans) context.Context {
	return context.WithValue(ctx, transCtx{}, trans)
}
