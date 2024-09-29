package queue

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type SegmentPool[E any] struct {
	pool *sync.Pool
}

func (s *SegmentPool[E]) getSegment() *segment[E] {
	return s.pool.Get().(*segment[E])
}

func (s *SegmentPool[E]) putSegment(seg *segment[E]) {
	s.pool.Put(seg)
}

func (s *SegmentPool[E]) NewPool() {
	if s.pool == nil {
		s.pool = &sync.Pool{
			New: func() any {
				return &segment[E]{
					buffer: make([]element[E], segmentSize),
					next:   nil,
				}
			},
		}
	}
}

type element[E any] struct {
	mu   sync.Mutex
	data E
}

type segment[E any] struct {
	buffer []element[E]
	next   unsafe.Pointer
}

func (s *segment[E]) reset() {
	s.buffer = make([]element[E], segmentSize)
	s.next = nil
}

func (s *segment[E]) get(cursor int64) E {
	//s.buffer[cursor].mu.Lock()
	e := s.buffer[cursor].data
	//s.buffer[cursor].mu.Unlock()
	return e
}

func (s *segment[E]) set(cursor int64, e E) {
	//s.buffer[cursor].mu.Lock()
	s.buffer[cursor].data = e
	//s.buffer[cursor].mu.Unlock()
}

func (s *segment[E]) newNextSegment(pool *SegmentPool[E]) *segment[E] {
	if s.nextSegment() != nil {
		return nil
	}
	seg := pool.getSegment()
	if s.storeNext(seg) {
		return seg
	}
	pool.putSegment(seg)
	return nil
}

func (s *segment[E]) nextSegment() *segment[E] {
	return (*segment[E])(atomic.LoadPointer(&s.next))
}

func (s *segment[E]) nextSegmentPointer() unsafe.Pointer {
	return atomic.LoadPointer(&s.next)
}

func (s *segment[E]) storeNext(seg *segment[E]) bool {
	return atomic.CompareAndSwapPointer(&s.next, nil, unsafe.Pointer(seg))
}

func (s *segment[E]) storePointer(seg unsafe.Pointer) bool {
	return atomic.CompareAndSwapPointer(&s.next, nil, seg)
}

func newSegmentPool[E any]() *SegmentPool[E] {
	return &SegmentPool[E]{
		pool: &sync.Pool{
			New: func() any {
				return &segment[E]{
					buffer: make([]element[E], segmentSize),
					next:   nil,
				}
			},
		},
	}
}
