package dorm

import (
	"github.com/origadmin/toolkits/context"
)

// ExecFunc is a function that can be executed within a transaction
type ExecFunc = func(context.Context) error

// Trans is a transaction wrapper
type Trans interface {
	Tx(ctx context.Context, fn func(context.Context) error) error
	InTx(ctx context.Context, fn func(tx Tx) error) error
}

type Tx interface {
	Commit() error
	Rollback() error
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
