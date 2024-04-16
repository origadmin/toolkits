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

func New(config Config) (ORM, error) {
	var (
		err  error
		name string
		conn Connector
	)
	dbs := make(map[string]any, len(config.Connectors))
	if len(config.Connectors) != 0 {
		for name, conn = range config.Connectors {
			dbs[name], err = conn.Open(config)
			if err != nil {
				return nil, err
			}
		}
	}
	return &orm{config: config, dbs: dbs}, nil
}

var _ ORM = (*orm)(nil)
