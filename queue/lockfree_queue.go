package queue

import (
	"sync/atomic"
	"unsafe"
)

const segmentSize = 1 << 10 // The size of each segment, here set to 1024

type segment[E any] struct {
	buffer [segmentSize]E
	next   unsafe.Pointer // *segment[E]
}

// LockFreeQueue Is a lock-free implementation of concurrent queues
type LockFreeQueue[E any] struct {
	head     unsafe.Pointer // *segment[E]
	tail     unsafe.Pointer // *segment[E]
	producer int64
	consumer int64
}

// NewLockFreeQueue Create a new LockFreeQueue
func NewLockFreeQueue[E any]() *LockFreeQueue[E] {
	seg := &segment[E]{}
	return &LockFreeQueue[E]{
		head: unsafe.Pointer(seg),
		tail: unsafe.Pointer(seg),
	}
}

// Offer Adds the element to the end of the queue
func (q *LockFreeQueue[E]) Offer(e E) bool {
	for {
		producer := atomic.LoadInt64(&q.producer)
		consumer := atomic.LoadInt64(&q.consumer)
		tailSeg := (*segment[E])(atomic.LoadPointer(&q.tail))

		if producer-consumer >= segmentSize {
			// The current segment is full. Try adding a new segment
			if tailSeg.next == nil {
				newSeg := &segment[E]{}
				if atomic.CompareAndSwapPointer(&tailSeg.next, nil, unsafe.Pointer(newSeg)) {
					atomic.CompareAndSwapPointer(&q.tail, unsafe.Pointer(tailSeg), unsafe.Pointer(newSeg))
				}
			} else {
				atomic.CompareAndSwapPointer(&q.tail, unsafe.Pointer(tailSeg), tailSeg.next)
			}
			continue
		}

		if atomic.CompareAndSwapInt64(&q.producer, producer, producer+1) {
			tailSeg.buffer[producer%segmentSize] = e
			q.tryResetCounters() // 尝试重置计数器
			return true
		}
	}
}

// Poll Removes an element from the queue header and returns it
func (q *LockFreeQueue[E]) Poll() (E, bool) {
	for {
		consumer := atomic.LoadInt64(&q.consumer)
		producer := atomic.LoadInt64(&q.producer)
		headSeg := (*segment[E])(atomic.LoadPointer(&q.head))

		if consumer >= producer {
			var zero E
			return zero, false // Queue empty
		}

		e := headSeg.buffer[consumer%segmentSize]
		if atomic.CompareAndSwapInt64(&q.consumer, consumer, consumer+1) {
			if consumer%segmentSize == segmentSize-1 && headSeg.next != nil {
				// The current segment has been consumed, move to the next segment
				atomic.CompareAndSwapPointer(&q.head, unsafe.Pointer(headSeg), headSeg.next)
			}
			q.tryResetCounters() // 尝试重置计数器
			return e, true
		}
	}
}

// 新增 tryResetCounters 方法
func (q *LockFreeQueue[E]) tryResetCounters() {
	for {
		producer := atomic.LoadInt64(&q.producer)
		consumer := atomic.LoadInt64(&q.consumer)

		if producer < segmentSize {
			return
		}

		newProducer := producer - consumer
		newConsumer := int64(0)

		if atomic.CompareAndSwapInt64(&q.producer, producer, newProducer) {
			if atomic.CompareAndSwapInt64(&q.consumer, consumer, newConsumer) {
				return
			}
			// 如果消费者计数器更新失败，回滚生产者计数器
			atomic.StoreInt64(&q.producer, producer)
		}
	}
}

// Peek returns the element at the head of the queue without removing it
func (q *LockFreeQueue[E]) Peek() (E, bool) {
	consumer := atomic.LoadInt64(&q.consumer)
	producer := atomic.LoadInt64(&q.producer)
	if consumer >= producer {
		var zero E
		return zero, false // queue empty
	}
	headSeg := (*segment[E])(atomic.LoadPointer(&q.head))
	return headSeg.buffer[consumer%segmentSize], true
}

// Size returns the number of elements in the queue
func (q *LockFreeQueue[E]) Size() int64 {
	for {
		producer := atomic.LoadInt64(&q.producer)
		consumer := atomic.LoadInt64(&q.consumer)
		if producer >= consumer {
			return int64(producer - consumer)
		}
		// 如果 producer < consumer，说明在读取过程中发生了重置
		// 重新尝试读取
	}
}

// IsEmpty returns true if the queue is empty
func (q *LockFreeQueue[E]) IsEmpty() bool {
	return q.Size() == 0
}

// Clear removes all elements from the queue
func (q *LockFreeQueue[E]) Clear() {
	for {
		if _, ok := q.Poll(); !ok {
			break
		}
	}
}

func (q *LockFreeQueue[E]) ToSlice() []E {
	result := make([]E, 0, q.Size())
	iter := q.Iterator()
	for iter.Next() {
		result = append(result, iter.Value())
	}
	return result
}

func (q *LockFreeQueue[E]) Iterator() Iterator[E] {
	return &lockFreeQueueIterator[E]{
		queue:   q,
		nextIdx: atomic.LoadInt64(&q.consumer),
		endIdx:  atomic.LoadInt64(&q.producer),
		currSeg: (*segment[E])(atomic.LoadPointer(&q.head)),
	}
}

// lockFreeQueueIterator 是 LockFreeQueue 的迭代器实现
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
			// 已经被消费的元素，跳过
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

func (it *lockFreeQueueIterator[E]) Value() E {
	if !it.valid {
		panic("Value called before Next or after end of iteration")
	}
	return it.currElem
}
