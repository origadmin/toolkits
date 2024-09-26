package queue

import (
	"sync"
)

type PoolQueue[E any] struct {
	pool *sync.Pool
}

func (p *PoolQueue[E]) Offer(e E) bool {
	p.pool.Put(e)
	return true
}

func (p *PoolQueue[E]) Poll() (E, bool) {
	v := p.pool.Get()
	if v == nil {
		var zero E
		return zero, false
	}
	return v.(E), true
}

func (p *PoolQueue[E]) Peek() (E, bool) {
	//TODO implement me
	panic("implement me")
}

func (p *PoolQueue[E]) Size() int64 {
	//TODO implement me
	panic("implement me")
}

func (p *PoolQueue[E]) IsEmpty() bool {
	//TODO implement me
	panic("implement me")
}

func (p *PoolQueue[E]) Clear() {
	//TODO implement me
	panic("implement me")
}

func (p *PoolQueue[E]) ToSlice() []E {
	//TODO implement me
	panic("implement me")
}

func (p *PoolQueue[E]) Iterator() Iterator[E] {
	//TODO implement me
	panic("implement me")
}

func NewPoolQueue[E any]() *PoolQueue[E] {
	return &PoolQueue[E]{
		pool: &sync.Pool{
			New: func() any {
				return nil
			},
		},
	}
}

var _ Queue[any] = (*PoolQueue[any])(nil)
