package orm

import "sync"

type Connector interface {
	Open(config Config) (any, error)
	Close() error
}

type ORM interface {
	Config() Config
	Connect(name string, connector Connector) error
	DB(name string) any
	Close() error
}

type orm struct {
	sync.RWMutex
	config Config
	dbs    map[string]any
	cls    []func() error
}

func (o *orm) Connect(name string, conn Connector) error {
	db, err := conn.Open(o.config)
	if err != nil {
		return err
	}
	o.Lock()
	o.dbs[name] = db
	o.cls = append(o.cls, conn.Close)
	o.Unlock()
	return nil
}

func (o *orm) Config() Config {
	return o.config
}

func (o *orm) DB(name string) any {
	o.RLock()
	db, ok := o.dbs[name]
	o.RUnlock()
	if !ok {
		panic("db not found " + name)
	}
	return db
}

func (o *orm) Close() error {
	for i := range o.cls {
		if err := o.cls[i](); err != nil {
			return err
		}
	}
	return nil
}

func New(config Config) ORM {
	return &orm{config: config, dbs: make(map[string]any)}
}

var _ ORM = (*orm)(nil)
