package queue

import (
	"runtime"
)

type LockFreeRetryQueue[E any] struct {
	*LockFreeQueue[E]
	retries int64
}

// Offer Adds the element to the end of the queue
func (q *LockFreeRetryQueue[E]) Offer(e E) bool {
	for retries := q.retries; retries > 0; retries-- {
		if q.LockFreeQueue.Offer(e) {
			return true
		}
		runtime.Gosched()
	}
	return false
}

// Poll Removes an element from the queue header and returns it
func (q *LockFreeRetryQueue[E]) Poll() (E, bool) {
	var zero E
	for retries := q.retries; retries > 0; retries-- {
		if e, ok := q.LockFreeQueue.Poll(); ok {
			return e, true
		}
		runtime.Gosched()
	}
	return zero, false
}

// Peek returns the element at the head of the queue without removing it
func (q *LockFreeRetryQueue[E]) Peek() (E, bool) {
	var zero E
	for retries := q.retries; retries > 0; retries-- {
		if e, ok := q.LockFreeQueue.Peek(); ok {
			return e, true
		}
		runtime.Gosched()
	}
	return zero, false
}

// Size returns the number of elements in the queue
func (q *LockFreeRetryQueue[E]) Size() int64 {
	return q.Size()
}

// IsEmpty returns true if the queue is empty
func (q *LockFreeRetryQueue[E]) IsEmpty() bool {
	return q.Size() == 0
}

// Clear removes all elements from the queue
func (q *LockFreeRetryQueue[E]) Clear() {
	q.Clear()
}

// ToSlice returns a slice representation of the queue
func (q *LockFreeRetryQueue[E]) ToSlice() []E {
	return q.ToSlice()
}

// Iterator returns an iterator for the queue
func (q *LockFreeRetryQueue[E]) Iterator() Iterator[E] {
	return q.Iterator()
}
