package queue

import (
	"sync"
	"sync/atomic"
)

type PoolQueue[E any] struct {
	private atomic.Pointer[element[E]]
	size    int64
	pool    *sync.Pool
}

func (p *PoolQueue[E]) Offer(e E) bool {
	p.pool.Put(&element[E]{
		data: e,
	})
	atomic.AddInt64(&p.size, 1)
	return true
}

func (p *PoolQueue[E]) Poll() (E, bool) {
	v := p.private.Load()
	if v != nil {
		p.private.Store(nil)
		atomic.AddInt64(&p.size, -1)
		return v.data, true
	}
	pv := p.pool.Get()
	if pv == nil {
		var zero E
		return zero, false
	}
	p.private.Store(nil)
	atomic.AddInt64(&p.size, -1)
	return pv.(*element[E]).data, true
}

func (p *PoolQueue[E]) Peek() (E, bool) {
	v := p.private.Load()
	if v != nil {
		return v.data, true
	}
	pv := p.pool.Get()
	if pv == nil {
		var zero E
		return zero, false
	}
	pvv := pv.(*element[E])
	p.private.Store(pvv)
	return pv.(*element[E]).data, true
}

func (p *PoolQueue[E]) Size() int64 {
	return atomic.LoadInt64(&p.size)
}

func (p *PoolQueue[E]) IsEmpty() bool {
	return p.Size() == 0
}

func (p *PoolQueue[E]) Clear() {
	p.pool = &sync.Pool{
		New: func() any {
			return nil
		},
	}
	return
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
