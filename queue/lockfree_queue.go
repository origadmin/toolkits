package queue

import (
	"sync"
	"sync/atomic"
)

const (
	segmentSize = 1 << 1 // The cursor of each segment, here set to 1024
	segmentMask = segmentSize - 1
)

// LockFreeQueue Is a lock-free implementation of concurrent queues
type LockFreeQueue[E any] struct {
	segmentPool *SegmentPool[E]
	head        *segment[E]
	headLock    sync.Mutex
	tail        *segment[E]
	tailLock    sync.Mutex
	cursor      int64
	producer    int64
	consumer    int64
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
		head:        seg,
		tail:        seg,
		cursor:      0,
	}
}

func (q *LockFreeQueue[E]) WithRetry(retry int64) *LockFreeRetryQueue[E] {
	return &LockFreeRetryQueue[E]{LockFreeQueue: q, retries: retry}
}

// Offer Adds the element to the end of the queue
func (q *LockFreeQueue[E]) Offer(e E) bool {
	tailSeg := q.tryTailSegment()
	if tailSeg == nil {
		return false
	}

	producer := q.getProducer()
	// lock current segment insert position
	if q.moveProducer(producer) {
		tailSeg.set(producer%segmentSize, e)
		// update cursor for consumer
		q.incrementCursor()
		return true
	}
	return false
}

// Poll Removes an element from the queue header and returns it
func (q *LockFreeQueue[E]) Poll() (E, bool) {
	var zero E
	if q.Size() < 1 {
		return zero, false // queue empty
	}

	headSeg := q.tryHeadSegment()
	if headSeg == nil {
		return zero, false
	}

	consumer := q.getConsumer()
	// lock current segment index
	if q.moveConsumer(consumer) {
		e := headSeg.get(consumer % segmentSize) // Get data first
		return e, true
	}

	return zero, false
}

// Peek returns the element at the head of the queue without removing it
func (q *LockFreeQueue[E]) Peek() (E, bool) {
	var zero E
	if q.Size() < 1 {
		return zero, false // queue empty
	}

	consumer := q.getConsumer()
	headSeg := q.getHeadSegment()
	if headSeg == nil {
		return zero, false
	}
	// Secure reads with atomic operations
	e := headSeg.get(consumer % segmentSize)
	return e, true
}

// Size returns the number of elements in the queue
func (q *LockFreeQueue[E]) Size() int64 {
	return atomic.LoadInt64(&q.cursor) - atomic.LoadInt64(&q.consumer)
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
		atomic.StoreInt64(&q.cursor, 0)
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
	seg := q.segmentPool.getSegment()
	//seg.reset()
	return seg
}

func (q *LockFreeQueue[E]) freeSegment(seg *segment[E]) {
	seg.reset()
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

func (q *LockFreeQueue[E]) incrementProducer() int64 {
	return atomic.AddInt64(&q.producer, 1)
}

func (q *LockFreeQueue[E]) getConsumer() int64 {
	return atomic.LoadInt64(&q.consumer)
}

func (q *LockFreeQueue[E]) incrementConsumer() int64 {
	return atomic.AddInt64(&q.consumer, 1)
}

func (q *LockFreeQueue[E]) getCursor() int64 {
	return atomic.LoadInt64(&q.cursor)
}

func (q *LockFreeQueue[E]) incrementCursor() int64 {
	return atomic.AddInt64(&q.cursor, 1)
}

func (q *LockFreeQueue[E]) decrementCursor() int64 {
	return atomic.AddInt64(&q.cursor, -1)
}

func (q *LockFreeQueue[E]) tryHeadSegment() *segment[E] {
	consumer := q.getConsumer()
	headSeg := q.getHeadSegment()
	if consumer%segmentSize < segmentMask {
		return headSeg
	}

	// The current segment has been consumed, move to the next segment
	if next := headSeg.nextSegment(); next != nil {
		if q.updateHeadSegment(headSeg, next) {
			return headSeg
		}
		return nil
	} else {
		return nil
	}
}

func (q *LockFreeQueue[E]) tryTailSegment() *segment[E] {
	producer := q.getProducer()
	tailSeg := q.getTailSegment()
	if producer%segmentSize < segmentMask {
		// try to create nSeg segment
		if nSeg := tailSeg.nextSegment(); nSeg == nil {
			nSeg := q.newSegment()
			if !tailSeg.storeNext(nSeg) {
				q.freeSegment(nSeg)
			}
		}
		return tailSeg
	}

	// The current segment is full. Try adding a new segment
	if nSeg := tailSeg.nextSegment(); nSeg == nil {
		nSeg = q.newSegment()
		if tailSeg.storeNext(nSeg) {
			if q.updateTailSegment(tailSeg, nSeg) {
				return tailSeg
			}
			return nil
		}
		q.freeSegment(nSeg)
		return nil
	} else {
		if q.updateTailSegment(tailSeg, nSeg) {
			return tailSeg
		}
		return nil
	}
}

func (q *LockFreeQueue[E]) getTailSegment() *segment[E] {
	q.tailLock.Lock()
	tail := q.tail
	q.tailLock.Unlock()
	return tail
	//return (*segment[E])(atomic.LoadPointer(&q.tail))
}

func (q *LockFreeQueue[E]) updateTailSegment(old, new *segment[E]) bool {
	q.tailLock.Lock()
	q.tail = new
	q.tailLock.Unlock()
	//return atomic.CompareAndSwapPointer(&q.tail, unsafe.Pointer(old), unsafe.Pointer(new))
	return true
}

//func (q *LockFreeQueue[E]) storeTailSegment(new *segment[E]) {
//	atomic.StorePointer(&q.tail, unsafe.Pointer(new))
//}

func (q *LockFreeQueue[E]) getHeadSegment() *segment[E] {
	q.headLock.Lock()
	head := q.head
	q.headLock.Unlock()
	return head

}

func (q *LockFreeQueue[E]) updateHeadSegment(old, new *segment[E]) bool {
	//return atomic.CompareAndSwapPointer(&q.head, unsafe.Pointer(old), unsafe.Pointer(new))
	q.headLock.Lock()
	q.head = new
	q.headLock.Unlock()
	return true
}

//func (q *LockFreeQueue[E]) storeHeadSegment(new *segment[E]) {
//	atomic.StorePointer(&q.head, unsafe.Pointer(new))
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
		it.currElem = it.currSeg.buffer[elementIndex].data
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
