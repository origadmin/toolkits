package queue

import (
	"sync/atomic"
	"unsafe"
)

const (
	segmentSize = 1 << 16 // The size of each segment, here set to 1024
	segmentMask = segmentSize - 1
)

// LockFreeQueue Is a lock-free implementation of concurrent queues
type LockFreeQueue[E any] struct {
	segmentPool *SegmentPool[E]
	head        unsafe.Pointer
	tail        unsafe.Pointer
	producer    int64
	consumer    int64
}

// NewLockFreeQueue Create a new LockFreeQueue
func NewLockFreeQueue[E any]() *LockFreeQueue[E] {
	//queue := &LockFreeQueue[E]{
	//	segmentPool: newSegmentPool[E](),
	//}
	//seg := queue.getSegment()
	//queue.head = unsafe.Pointer(seg)
	//queue.tail = unsafe.Pointer(seg)
	return NewLockFreeQueueWithPool(&SegmentPool[E]{})
}

// NewLockFreeQueueWithPool Create a new LockFreeQueue with a pre-created segment pool
func NewLockFreeQueueWithPool[E any](pool *SegmentPool[E]) *LockFreeQueue[E] {
	if pool == nil {
		panic("NewLockFreeQueueWithPool: pool cannot be nil")
	}
	pool.NewPool()
	seg := pool.Get()
	return &LockFreeQueue[E]{
		segmentPool: pool,
		head:        unsafe.Pointer(seg),
		tail:        unsafe.Pointer(seg),
	}
}

// Offer Adds the element to the end of the queue
func (q *LockFreeQueue[E]) Offer(e E) bool {
	producer := q.getProducer()
	tailSeg := q.getTailSegment()
	index := producer % segmentSize
	if index == segmentMask {
		// The current segment is full. Try adding a new segment
		if tailSeg.nextSegment() == nil {
			newSeg := q.getSegment()
			if tailSeg.storeNext(newSeg) {
				q.updateTailSegment(tailSeg, newSeg)
			} else {
				q.freeSegment(newSeg)
				return false
			}
		} else {
			return false
		}
	}

	if q.updateProducer(producer, producer+1) {
		tailSeg.set(index, e)
		return true
	}
	return false
}

// Poll Removes an element from the queue header and returns it
func (q *LockFreeQueue[E]) Poll() (E, bool) {
	var zero E
	consumer := q.getConsumer()
	producer := q.getProducer()
	if consumer-producer < 1 {
		return zero, false // Queue empty
	}
	headSeg := q.getHeadSegment()
	index := consumer % segmentSize
	e := headSeg.get(index) // Get data first
	if q.updateConsumer(consumer, consumer+1) {
		if index == segmentMask {
			// The current segment has been consumed, move to the next segment
			if next := headSeg.nextSegment(); next != nil {
				q.updateHeadSegment(headSeg, next)
			}
		}
		return e, true
	}
	return zero, false
}

// Peek returns the element at the head of the queue without removing it
func (q *LockFreeQueue[E]) Peek() (E, bool) {
	var zero E
	consumer := q.getConsumer()
	producer := q.getProducer()
	if consumer >= producer {
		return zero, false // queue empty
	}

	headSeg := q.getHeadSegment()
	// Secure reads with atomic operations
	e := headSeg.get(consumer % segmentSize)
	return e, true
}

// Size returns the number of elements in the queue
func (q *LockFreeQueue[E]) Size() int64 {
	producer := q.getProducer()
	consumer := q.getConsumer()
	if producer > consumer {
		return producer - consumer
	}
	// If producer < consumer, a reset occurred during the read
	// Try reading again
	return 0
}

// IsEmpty returns true if the queue is empty
func (q *LockFreeQueue[E]) IsEmpty() bool {
	return q.Size() == 0
}

// Clear removes all elements from the queue
func (q *LockFreeQueue[E]) Clear() {
	// 1. Saves the status of the current queue
	head := q.getHeadSegment()
	tail := q.getHeadSegment()

	// 2. Create a new empty queue or reuse the old one
	seg := q.getSegment()

	// 3. Updates the status of the queue atomically
	if q.updateHeadSegment(head, seg) && q.updateTailSegment(tail, seg) {
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
		nextIdx: q.getConsumer(),
		endIdx:  q.getProducer(),
		currSeg: q.getHeadSegment(),
	}
}

func (q *LockFreeQueue[E]) getSegment() *segment[E] {
	return q.segmentPool.Get()
}

func (q *LockFreeQueue[E]) freeSegment(seg *segment[E]) {
	q.segmentPool.Put(seg)
}

// releaseSegment recursively releases all segments of the old queue
func (q *LockFreeQueue[E]) releaseSegment(head, tail *segment[E]) {
	if head == nil {
		return
	}
	if next := head.nextSegment(); next != nil {
		q.releaseSegment(next, tail)
	}

	q.freeSegment(head)
}

func (q *LockFreeQueue[E]) updateProducer(producer int64, value int64) bool {
	return atomic.CompareAndSwapInt64(&q.producer, producer, value)
}

func (q *LockFreeQueue[E]) updateConsumer(consumer int64, value int64) bool {
	return atomic.CompareAndSwapInt64(&q.consumer, consumer, value)
}

func (q *LockFreeQueue[E]) getProducer() int64 {
	return atomic.LoadInt64(&q.producer)
}

func (q *LockFreeQueue[E]) getConsumer() int64 {
	return atomic.LoadInt64(&q.consumer)
}

//func (q *LockFreeQueue[E]) tryTailSegment() *segment[E] {
//	producer := q.getProducer()
//	tailSeg := q.getTailSegment()
//	if producer%segmentSize == segmentMask {
//		// The current segment is full. Try adding a new segment
//		if next := tailSeg.nextSegment(); next == nil {
//			newSeg := q.getSegment()
//			if tailSeg.storeNext(newSeg) {
//				q.updateTailSegment(tailSeg, newSeg)
//			} else {
//				q.freeSegment(newSeg)
//			}
//		} else {
//			q.updateTailSegment(tailSeg, next)
//		}
//	}
//	return tailSeg
//}

func (q *LockFreeQueue[E]) getTailSegment() *segment[E] {
	return (*segment[E])(atomic.LoadPointer(&q.tail))
}

func (q *LockFreeQueue[E]) updateTailSegment(old, new *segment[E]) bool {
	return atomic.CompareAndSwapPointer(&q.tail, unsafe.Pointer(old), unsafe.Pointer(new))
}

func (q *LockFreeQueue[E]) getHeadSegment() *segment[E] {
	return (*segment[E])(atomic.LoadPointer(&q.head))
}

func (q *LockFreeQueue[E]) updateHeadSegment(old, new *segment[E]) bool {
	return atomic.CompareAndSwapPointer(&q.head, unsafe.Pointer(old), unsafe.Pointer(new))
}

// Added the tryResetHeadCounters method
//func (q *LockFreeQueue[E]) tryResetHeadCounters() {
//	for {
//		producer := atomic.LoadInt64(&q.producer)
//		consumer := atomic.LoadInt64(&q.consumer)
//		// Checks whether the current segment is fully consumed
//		if consumer%segmentSize == segmentMask {
//			// Try to reset the consumer and producer
//			newConsumer := int64(0)
//			newProducer := producer - consumer
//			if atomic.CompareAndSwapInt64(&q.consumer, consumer, newConsumer) {
//				if !atomic.CompareAndSwapInt64(&q.producer, producer, newProducer) {
//					// If the consumer update fails, the producer is rolled back
//					atomic.StoreInt64(&q.producer, producer)
//				} else {
//					return
//				}
//			} else {
//				// If the consumer update fails, the producer is rolled back
//				atomic.StoreInt64(&q.producer, producer)
//			}
//		} else {
//			atomic.AddInt64(&q.consumer, 1)
//			return
//		}
//	}
//}

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
			if it.currSeg.next == nil {
				it.valid = false
				return false
			}
			it.currSeg = it.currSeg.nextSegment()
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
