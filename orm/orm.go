package orm

import (
	"context"
	"sync"
)

type Connector interface {
	Open(config Config) (any, error)
	Close() error
}

type ORM struct {
	sync.RWMutex
	dbs    map[string]any
	conns  map[string]Connector
	closes []func() error
	cfg    Config
}

func New(config Config) *ORM {
	if config.Context == nil {
		config.Context = context.Background()
	}
	orm := &ORM{
		cfg: config,
		dbs: make(map[string]any),
	}
	return orm
}

func (o *ORM) Config() Config {
	return o.cfg
}

func (o *ORM) Close() error {
	var cls func() error
	for _, cls = range o.closes {
		if err := cls(); err != nil {
			return err
		}
	}
	return nil
}

func (o *ORM) Open(name string) error {
	conn, ok := o.conns[name]
	if !ok {
		conn = o.cfg.Connector
	}
	db, err := conn.Open(o.cfg)
	if err != nil {
		return err
	}
	o.Lock()
	o.closes = append(o.closes, conn.Close)
	o.dbs[name] = db
	o.Unlock()
	return nil
}

func (o *ORM) DB(name string) any {
	o.RLock()
	db := o.dbs[name]
	o.RUnlock()
	return db
}
