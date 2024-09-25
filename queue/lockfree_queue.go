package queue

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	segmentSize = 1 << 2 // The size of each segment, here set to 1024
	segmentMask = segmentSize - 1
	maxRetries  = 3
)

type segment[E any] struct {
	buffer []E
	index  int32
	next   unsafe.Pointer // *segment[E]
}

func (s *segment[E]) reset() {
	clear(s.buffer)
	s.next = nil
}

// LockFreeQueue Is a lock-free implementation of concurrent queues
type LockFreeQueue[E any] struct {
	segmentPool sync.Pool
	head        unsafe.Pointer // *segment[E]
	tail        unsafe.Pointer // *segment[E]
	producer    int64
	consumer    int64
}

// NewLockFreeQueue Create a new LockFreeQueue
func NewLockFreeQueue[E any]() *LockFreeQueue[E] {
	queue := &LockFreeQueue[E]{
		segmentPool: sync.Pool{
			New: func() any {
				return &segment[E]{
					buffer: make([]E, segmentSize),
				}
			},
		},
	}
	seg := queue.getSegment()
	queue.head = unsafe.Pointer(seg)
	queue.tail = unsafe.Pointer(seg)
	return queue
}

// Offer Adds the element to the end of the queue
func (q *LockFreeQueue[E]) Offer(e E) bool {
	producer := atomic.LoadInt64(&q.producer)
	//consumer := atomic.LoadInt64(&q.consumer)
	tailSeg := (*segment[E])(atomic.LoadPointer(&q.tail))
	if producer%segmentSize == segmentMask {
		//index := atomic.LoadInt32(&tailSeg.index)
		//if index >= segmentSize {
		// The current segment is full. Try adding a new segment
		if atomic.LoadPointer(&tailSeg.next) == nil {
			newSeg := q.getSegment()
			if atomic.CompareAndSwapPointer(&tailSeg.next, nil, unsafe.Pointer(newSeg)) {
				atomic.CompareAndSwapPointer(&q.tail, unsafe.Pointer(tailSeg), unsafe.Pointer(newSeg))
			} else {
				q.freeSegment(newSeg)
			}
		} else {
			atomic.CompareAndSwapPointer(&q.tail, unsafe.Pointer(tailSeg), atomic.LoadPointer(&tailSeg.next))
		}
		// Clear the current index because it has been moved to a new segment
		//atomic.StoreInt32(&tailSeg.index, 0)
		//}
	}

	if atomic.CompareAndSwapInt64(&q.producer, producer, producer+1) {
		tailSeg.buffer[producer%segmentSize] = e
		//q.tryResetTailCounters() // Attempt to reset the counter
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
	//fmt.Println("atomic.CompareAndSwapInt64(&q.consumer, consumer, consumer+1)", consumer)
	if atomic.CompareAndSwapInt64(&q.consumer, consumer, consumer+1) {
		//fmt.Println("atomic.CompareAndSwapInt64(&q.consumer, consumer, consumer+1) success", consumer)
		headSeg := (*segment[E])(atomic.LoadPointer(&q.head))
		if consumer%segmentSize == segmentMask && headSeg.next != nil {
			//fmt.Println("consumer % segmentSize == segmentMask && headSeg.next != nil")
			// The current segment has been consumed, move to the next segment
			if atomic.CompareAndSwapPointer(&q.head, unsafe.Pointer(headSeg), headSeg.next) {

			}
		}
		q.tryResetHeadCounters() // Attempt to reset the counter
		//fmt.Println("return headSeg.buffer[consumer % segmentSize]", headSeg.buffer[consumer%segmentSize])
		return headSeg.buffer[consumer%segmentSize], true
	}
	//fmt.Println("Failed to remove element from the queue")
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

	headSeg := (*segment[E])(atomic.LoadPointer(&q.head))
	// Secure reads with atomic operations
	e := headSeg.buffer[consumer%segmentSize]
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
	head := atomic.LoadPointer(&q.head)
	tail := atomic.LoadPointer(&q.tail)

	// 2. Create a new empty queue
	seg := q.getSegment()
	newHead := seg
	newTail := seg

	// 3. Updates the status of the queue atomically
	if atomic.CompareAndSwapPointer(&q.head, head, unsafe.Pointer(newHead)) &&
		atomic.CompareAndSwapPointer(&q.tail, tail, unsafe.Pointer(newTail)) {
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
		currSeg: (*segment[E])(atomic.LoadPointer(&q.head)),
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
func (q *LockFreeQueue[E]) releaseSegment(head, tail unsafe.Pointer) {
	if head == nil {
		return
	}

	seg := (*segment[E])(head)
	if seg.next != nil {
		q.releaseSegment(seg.next, tail)
	}

	q.freeSegment(seg)
}

func (q *LockFreeQueue[E]) tryTailSegment() *segment[E] {
	producer := atomic.LoadInt64(&q.producer)
	//consumer := atomic.LoadInt64(&q.consumer)
	tailSeg := (*segment[E])(atomic.LoadPointer(&q.tail))
	if producer%segmentSize == segmentMask {
		//index := atomic.LoadInt32(&tailSeg.index)
		//if index >= segmentSize {
		// The current segment is full. Try adding a new segment
		if tailSeg.next == nil {
			newSeg := q.getSegment()
			if atomic.CompareAndSwapPointer(&tailSeg.next, nil, unsafe.Pointer(newSeg)) {
				atomic.CompareAndSwapPointer(&q.tail, unsafe.Pointer(tailSeg), unsafe.Pointer(newSeg))
			} else {
				q.freeSegment(newSeg)
			}
		} else {
			atomic.CompareAndSwapPointer(&q.tail, unsafe.Pointer(tailSeg), tailSeg.next)
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
		if consumer%segmentSize == 0 {
			// Try to reset the consumer and producer
			newConsumer := int64(0)
			newProducer := producer - consumer
			if atomic.CompareAndSwapInt64(&q.consumer, consumer, newConsumer) {
				if atomic.CompareAndSwapInt64(&q.producer, producer, newProducer) {
					//// Update the index of the current segment
					//headSeg := (*segment[E])(atomic.LoadPointer(&q.head))
					////atomic.StoreInt32(&headSeg.index, 0)
					//
					//// Swap next pointer
					//if headSeg.next != nil {
					//	atomic.CompareAndSwapPointer(&q.head, unsafe.Pointer(headSeg), headSeg.next)
					//}
					return
				} else {
					// If the consumer update fails, the producer is rolled back
					atomic.StoreInt64(&q.producer, producer)
				}
			}
			continue
		}
		return
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
			if it.currSeg.next == nil {
				it.valid = false
				return false
			}
			it.currSeg = (*segment[E])(it.currSeg.next)
		}

		it.currElem = it.currSeg.buffer[elementIndex]
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
