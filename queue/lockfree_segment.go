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
	seg.reset()
	s.pool.Put(seg)
}

func (s *SegmentPool[E]) NewPool() {
	if s.pool == nil {
		s.pool = &sync.Pool{
			New: func() any {
				return &segment[E]{
					buffer: make([]E, segmentSize),
					next:   nil,
				}
			},
		}
	}
}

type segment[E any] struct {
	buffer []E
	next   unsafe.Pointer
}

func (s *segment[E]) reset() {
	clear(s.buffer)
	s.next = nil
}

func (s *segment[E]) get(index int64) E {
	return s.buffer[index]
}

func (s *segment[E]) set(index int64, e E) {
	s.buffer[index] = e
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
					buffer: make([]E, segmentSize),
					next:   nil,
				}
			},
		},
	}
}
