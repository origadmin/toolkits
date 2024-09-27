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
	size        int64 // Keep size as a local variable
}

// NewLockFreeQueue Create a new LockFreeQueue
func NewLockFreeQueue[E any]() *LockFreeQueue[E] {
	return NewLockFreeQueueWithPool(&SegmentPool[E]{})
}

// NewLockFreeQueueWithPool Create a new LockFreeQueue with a pre-created segment pool
func NewLockFreeQueueWithPool[E any](pool *SegmentPool[E]) *LockFreeQueue[E] {
	if pool == nil {
		panic("NewLockFreeQueueWithPool: pool cannot be nil")
	}
	pool.NewPool()
	seg := pool.getSegment()
	return &LockFreeQueue[E]{
		segmentPool: pool,
		head:        unsafe.Pointer(seg),
		tail:        unsafe.Pointer(seg),
		size:        0,
	}
}

// Offer Adds the element to the end of the queue
func (q *LockFreeQueue[E]) Offer(e E) bool {
	producer := q.getProducer()
	tailSeg := q.getTailSegment()
	position := producer % segmentSize
	// If position to the end, create next for next offer
	if position == segmentMask {
		// The current segment is full. Try adding a new segment
		if tailSeg.nextSegment() == nil {
			newSeg := q.newSegment()
			if tailSeg.storeNext(newSeg) {
				// update tail pointer to new segment for next offer
				q.storeTailSegment(newSeg)
			} else {
				// next will alway nil for create,otherwise return
				q.freeSegment(newSeg)
				return false
			}
		} else {
			return false
		}
	}
	// lock current segment insert position
	if q.moveProducer(producer) {
		tailSeg.set(position, e)
		// update cursor for consumer
		q.incrementSize()
		return true
	}
	return false
}

// Poll Removes an element from the queue header and returns it
func (q *LockFreeQueue[E]) Poll() (E, bool) {
	var zero E
	consumer := q.getConsumer()
	if q.Size() < 1 {
		return zero, false // queue empty
	}
	// Secure reads with atomic operations
	headSeg := q.getHeadSegment()
	position := consumer % segmentSize
	// lock current segment index
	if q.moveConsumer(consumer) {
		q.decrementSize() // Decrement size
		if position == segmentMask {
			// The current segment has been consumed, move to the next segment
			if next := headSeg.nextSegment(); next != nil {
				q.storeHeadSegment(next)
			}
		}
		e := headSeg.get(position) // Get data first
		return e, true
	}
	return zero, false
}

// Peek returns the element at the head of the queue without removing it
func (q *LockFreeQueue[E]) Peek() (E, bool) {
	var zero E
	consumer := q.getConsumer()
	if q.Size() < 1 {
		return zero, false // queue empty
	}
	position := consumer % segmentSize
	headSeg := q.getHeadSegment()
	// Secure reads with atomic operations
	e := headSeg.get(position)
	return e, true
}

// Size returns the number of elements in the queue
func (q *LockFreeQueue[E]) Size() int64 {
	return atomic.LoadInt64(&q.size)
	//cursor := q.getCursor()
	//consumer := q.getConsumer()
	//capacity := cursor - consumer
	//if capacity > 0 {
	//	return capacity
	//}
	//// If cursor < consumer, a reset occurred during the read
	//// Try reading again
	//return 0
}

// IsEmpty returns true if the queue is empty
func (q *LockFreeQueue[E]) IsEmpty() bool {
	return q.Size() == 0
}

// Clear removes all elements from the queue
func (q *LockFreeQueue[E]) Clear() {
	// 1. Saves the status of the current queue
	head := q.getHeadSegment()
	tail := q.getTailSegment()

	// 2. Create a new empty queue or reuse the old one
	seg := q.newSegment()

	// 3. Updates the status of the queue atomically
	if q.updateHeadSegment(head, seg) && q.updateTailSegment(tail, seg) {
		atomic.StoreInt64(&q.producer, 0)
		atomic.StoreInt64(&q.consumer, 0)
		atomic.StoreInt64(&q.size, 0)
		// 4. Release old queue resources
		q.releaseSegment(head, tail)
	} else {
		// If the setting fails, try again
		q.Clear()
	}
}

// ToSlice returns a slice representation of the queue
func (q *LockFreeQueue[E]) ToSlice() []E {
	size := q.Size()
	result := make([]E, 0, size)
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

func (q *LockFreeQueue[E]) newSegment() *segment[E] {
	return q.segmentPool.getSegment()
}

func (q *LockFreeQueue[E]) freeSegment(seg *segment[E]) {
	q.segmentPool.putSegment(seg)
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

func (q *LockFreeQueue[E]) moveProducer(producer int64) bool {
	return atomic.CompareAndSwapInt64(&q.producer, producer, producer+1)
}

func (q *LockFreeQueue[E]) addProducer(delta int64) int64 {
	return atomic.AddInt64(&q.producer, delta)
}

func (q *LockFreeQueue[E]) moveConsumer(consumer int64) bool {
	return atomic.CompareAndSwapInt64(&q.consumer, consumer, consumer+1)
}

func (q *LockFreeQueue[E]) addConsumer(delta int64) int64 {
	return atomic.AddInt64(&q.consumer, delta)
}

func (q *LockFreeQueue[E]) getProducer() int64 {
	return atomic.LoadInt64(&q.producer)
}

func (q *LockFreeQueue[E]) getConsumer() int64 {
	return atomic.LoadInt64(&q.consumer)
}

func (q *LockFreeQueue[E]) incrementSize() int64 {
	return atomic.AddInt64(&q.size, 1)
}

func (q *LockFreeQueue[E]) decrementSize() int64 {
	return atomic.AddInt64(&q.size, -1)
}

func (q *LockFreeQueue[E]) tryTailSegment(tailSeg *segment[E], producer int64) {
	if producer%segmentSize == segmentMask {
		// The current segment is full. Try adding a new segment
		if next := tailSeg.nextSegment(); next == nil {
			newSeg := q.newSegment()
			if tailSeg.storeNext(newSeg) {
				q.storeTailSegment(newSeg)
			} else {
				q.freeSegment(newSeg)
			}
		} else {
			// next will alway nil for create,otherwise return
		}
	}
}

func (q *LockFreeQueue[E]) getTailSegment() *segment[E] {
	return (*segment[E])(atomic.LoadPointer(&q.tail))
}

func (q *LockFreeQueue[E]) updateTailSegment(old, new *segment[E]) bool {
	return atomic.CompareAndSwapPointer(&q.tail, unsafe.Pointer(old), unsafe.Pointer(new))
}
func (q *LockFreeQueue[E]) storeTailSegment(new *segment[E]) {
	atomic.StorePointer(&q.tail, unsafe.Pointer(new))
}

func (q *LockFreeQueue[E]) getHeadSegment() *segment[E] {
	return (*segment[E])(atomic.LoadPointer(&q.head))
}

func (q *LockFreeQueue[E]) updateHeadSegment(old, new *segment[E]) bool {
	return atomic.CompareAndSwapPointer(&q.head, unsafe.Pointer(old), unsafe.Pointer(new))
}

func (q *LockFreeQueue[E]) storeHeadSegment(new *segment[E]) {
	atomic.StorePointer(&q.head, unsafe.Pointer(new))
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
