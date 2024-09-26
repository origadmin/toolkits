package queue

import (
	"sync"
	"sync/atomic"
)

const (
	segmentSize = 1 << 2 // The size of each segment, here set to 1024
	segmentMask = segmentSize - 1
)

type segment[E any] struct {
	buffer []E
	next   atomic.Pointer[segment[E]]
}

func (s *segment[E]) reset() {
	clear(s.buffer)
	s.next.Store(nil)
}

func (s *segment[E]) get(index int64) E {
	return s.buffer[index]
}

func (s *segment[E]) set(index int64, e E) {
	s.buffer[index] = e
}

// LockFreeQueue Is a lock-free implementation of concurrent queues
type LockFreeQueue[E any] struct {
	segmentPool *sync.Pool
	head        atomic.Pointer[segment[E]]
	tail        atomic.Pointer[segment[E]]
	producer    int64
	consumer    int64
}

func newSegmentPoll[E any]() *sync.Pool {
	return &sync.Pool{
		New: func() any {
			return &segment[E]{
				buffer: make([]E, segmentSize),
				next:   atomic.Pointer[segment[E]]{},
			}
		},
	}
}

// NewLockFreeQueue Create a new LockFreeQueue
func NewLockFreeQueue[E any]() *LockFreeQueue[E] {
	queue := &LockFreeQueue[E]{
		segmentPool: newSegmentPoll[E](),
	}
	seg := queue.getSegment()
	queue.head.Store(seg)
	queue.tail.Store(seg)
	return queue
}

// Offer Adds the element to the end of the queue
func (q *LockFreeQueue[E]) Offer(e E) bool {
	producer := atomic.LoadInt64(&q.producer)
	tailSeg := q.tail.Load()
	index := producer % segmentSize
	if index == segmentMask {
		// The current segment is full. Try adding a new segment
		if tailSeg.next.Load() == nil {
			newSeg := q.getSegment()
			if tailSeg.next.CompareAndSwap(nil, newSeg) {
				q.tail.CompareAndSwap(tailSeg, newSeg)
			} else {
				q.freeSegment(newSeg)
				return false
			}
		} else {
			return false
		}
	}

	if atomic.CompareAndSwapInt64(&q.producer, producer, producer+1) {
		tailSeg.set(index, e)
		return true
	}
	return false
}

// Poll Removes an element from the queue header and returns it
func (q *LockFreeQueue[E]) Poll() (E, bool) {
	var zero E
	consumer := atomic.LoadInt64(&q.consumer)
	producer := atomic.LoadInt64(&q.producer)
	if consumer >= producer {
		return zero, false // Queue empty
	}
	headSeg := q.head.Load()
	index := consumer % segmentSize
	if atomic.CompareAndSwapInt64(&q.consumer, consumer, consumer+1) {
		e := headSeg.get(index) // Get data first
		if index == segmentMask {
			// The current segment has been consumed, move to the next segment
			if headSeg.next.Load() != nil {
				q.head.CompareAndSwap(headSeg, headSeg.next.Load())
			}
		}
		return e, true
	}
	return zero, false
}

// Peek returns the element at the head of the queue without removing it
func (q *LockFreeQueue[E]) Peek() (E, bool) {
	var zero E
	consumer := atomic.LoadInt64(&q.consumer)
	producer := atomic.LoadInt64(&q.producer)
	if consumer >= producer {
		return zero, false // queue empty
	}

	headSeg := q.head.Load()
	// Secure reads with atomic operations
	e := headSeg.get(consumer % segmentSize)
	return e, true
}

// Size returns the number of elements in the queue
func (q *LockFreeQueue[E]) Size() int64 {
	for {
		producer := atomic.LoadInt64(&q.producer)
		consumer := atomic.LoadInt64(&q.consumer)
		if producer >= consumer {
			return producer - consumer
		}
		// If producer < consumer, a reset occurred during the read
		// Try reading again
	}
}

// IsEmpty returns true if the queue is empty
func (q *LockFreeQueue[E]) IsEmpty() bool {
	return q.Size() == 0
}

// Clear removes all elements from the queue
func (q *LockFreeQueue[E]) Clear() {
	// 1. Saves the status of the current queue
	head := q.head.Load()
	tail := q.tail.Load()

	// 2. Create a new empty queue or reuse the old one
	seg := q.getSegment()

	// 3. Updates the status of the queue atomically
	if q.head.CompareAndSwap(head, seg) && q.tail.CompareAndSwap(tail, seg) {
		atomic.StoreInt64(&q.producer, 0)
		atomic.StoreInt64(&q.consumer, 0)

		// 4. Release old queue resources
		q.releaseSegment(head, tail)
	} else {
		// If the setting fails, try again
		q.Clear()
	}
}

// ToSlice returns a slice representation of the queue
func (q *LockFreeQueue[E]) ToSlice() []E {
	result := make([]E, 0, q.Size())
	iter := q.Iterator()
	for iter.Next() {
		result = append(result, iter.Value())
	}
	return result
}

// Iterator returns an iterator for the queue
func (q *LockFreeQueue[E]) Iterator() Iterator[E] {
	return &lockFreeQueueIterator[E]{
		queue:   q,
		nextIdx: atomic.LoadInt64(&q.consumer),
		endIdx:  atomic.LoadInt64(&q.producer),
		currSeg: q.head.Load(),
	}
}

func (q *LockFreeQueue[E]) getSegment() *segment[E] {
	return q.segmentPool.Get().(*segment[E])
}

func (q *LockFreeQueue[E]) freeSegment(seg *segment[E]) {
	seg.reset()
	q.segmentPool.Put(seg)
}

// releaseSegment recursively releases all segments of the old queue
func (q *LockFreeQueue[E]) releaseSegment(head, tail *segment[E]) {
	if head == nil {
		return
	}
	next := head.next.Load()
	if next != nil {
		q.releaseSegment(next, tail)
	}

	q.freeSegment(head)
}

func (q *LockFreeQueue[E]) tryTailSegment() *segment[E] {
	producer := atomic.LoadInt64(&q.producer)
	//consumer := atomic.LoadInt64(&q.consumer)
	tailSeg := q.tail.Load()
	if producer%segmentSize == segmentMask {
		//index := atomic.LoadInt32(&tailSeg.index)
		//if index >= segmentSize {
		// The current segment is full. Try adding a new segment
		if tailSeg.next.Load() == nil {
			newSeg := q.getSegment()
			if tailSeg.next.CompareAndSwap(nil, newSeg) {
				q.tail.CompareAndSwap(tailSeg, newSeg)
			} else {
				q.freeSegment(newSeg)
			}
		} else {
			q.tail.CompareAndSwap(tailSeg, tailSeg.next.Load())
		}
		// Clear the current index because it has been moved to a new segment
		//atomic.StoreInt32(&tailSeg.index, 0)
		//}
	}
	return tailSeg
}

// Added the tryResetHeadCounters method
func (q *LockFreeQueue[E]) tryResetHeadCounters() {
	for {
		producer := atomic.LoadInt64(&q.producer)
		consumer := atomic.LoadInt64(&q.consumer)
		// Checks whether the current segment is fully consumed
		if consumer%segmentSize == segmentMask {
			// Try to reset the consumer and producer
			newConsumer := int64(0)
			newProducer := producer - consumer
			if atomic.CompareAndSwapInt64(&q.consumer, consumer, newConsumer) {
				if !atomic.CompareAndSwapInt64(&q.producer, producer, newProducer) {
					// If the consumer update fails, the producer is rolled back
					atomic.StoreInt64(&q.producer, producer)
				} else {
					return
				}
			} else {
				// If the consumer update fails, the producer is rolled back
				atomic.StoreInt64(&q.producer, producer)
			}
		} else {
			atomic.AddInt64(&q.consumer, 1)
			return
		}
	}
}

// lockFreeQueueIterator is an iterator implementation of LockFreeQueue
type lockFreeQueueIterator[E any] struct {
	queue    *LockFreeQueue[E]
	nextIdx  int64
	endIdx   int64
	currSeg  *segment[E]
	currElem E
	valid    bool
}

func (it *lockFreeQueueIterator[E]) Next() bool {
	for {
		if it.nextIdx >= it.endIdx {
			it.valid = false
			return false
		}

		consumer := atomic.LoadInt64(&it.queue.consumer)
		if it.nextIdx < consumer {
			// Elements that have already been consumed, skip
			it.nextIdx = consumer
			continue
		}

		segmentIndex := it.nextIdx / segmentSize
		elementIndex := it.nextIdx % segmentSize

		for i := int64(0); i < segmentIndex; i++ {
			if it.currSeg.next.Load() == nil {
				it.valid = false
				return false
			}
			it.currSeg = it.currSeg.next.Load()
		}

		// it.currSeg.bufferMutex.RLock()
		it.currElem = it.currSeg.buffer[elementIndex]
		// it.currSeg.bufferMutex.RUnlock()
		it.nextIdx++
		it.valid = true
		return true
	}
}

// Value returns the current element
func (it *lockFreeQueueIterator[E]) Value() E {
	if !it.valid {
		panic("Value called before Next or after end of iteration")
	}
	return it.currElem
}

// NewQueue creates a new Queue
func NewQueue[E any]() Queue[E] {
	return NewLockFreeQueue[E]()
}
