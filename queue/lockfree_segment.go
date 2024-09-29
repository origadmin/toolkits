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
					buffer: make([]E, segmentSize),
					next:   nil,
				}
			},
		}
	}
}

type segment[E any] struct {
	cursor int64
	buffer []E
	next   unsafe.Pointer
}

func (s *segment[E]) reset() {
	s.cursor = 0
	s.buffer = make([]E, segmentSize)
	s.next = nil
}

func (s *segment[E]) hasSpace() bool {
	return atomic.LoadInt64(&s.cursor) < segmentSize
}

func (s *segment[E]) get(cursor int64) E {
	return s.buffer[cursor]
}

func (s *segment[E]) set(cursor int64, e E) {
	s.buffer[cursor] = e
}

func (s *segment[E]) trySegmentNext(pool *SegmentPool[E]) *segment[E] {
	if s.hasSpace() {
		return s
	}
	if nSeg := s.nextSegment(); nSeg != nil {
		return nSeg
	}
	nSeg := pool.getSegment()
	if s.storeNext(nSeg) {
		return nSeg
	}
	pool.putSegment(nSeg)
	return nil
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
					buffer: make([]E, segmentSize),
					next:   nil,
				}
			},
		},
	}
}
