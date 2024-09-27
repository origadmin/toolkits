package queue

import (
	"sync/atomic"
)

type ChannelQueue[E any] struct {
	private atomic.Pointer[element[E]]
	size    int64
	ch      chan E // Add a channel field
}

// Clear implements Queue.
func (q *ChannelQueue[E]) Clear() {
	for !q.IsEmpty() {
		q.private.Store(nil)
		<-q.ch // Empty the channel of all elements
	}
}

// Iterator implements Queue.
func (q *ChannelQueue[E]) Iterator() Iterator[E] {
	panic("implements")
	// 返回一个迭代器，遍历通道中的元素
	// return func() (E, bool) {
	// 	if q.IsEmpty() {
	// 		return nil, false // 如果通道为空，返回 nil
	// 	}
	// 	return <-q.ch, true // 从通道中取出元素
	// }
}

// Offer implements Queue.
func (q *ChannelQueue[E]) Offer(item E) bool {
	select {
	case q.ch <- item: // Try to send elements to a channel
		atomic.AddInt64(&q.size, 1)
		return true
	default:
		return false // If the channel is full, return false
	}
}

// Peek implements Queue.
func (q *ChannelQueue[E]) Peek() (E, bool) {
	var zero E
	v := q.private.Load()
	if v != nil {
		return v.data, true
	}
	select {
	case item := <-q.ch: // Try to extract the element from the channel
		// q.ch <- item // Put it back in the channel
		q.private.Store(&element[E]{
			data: item,
		})
		return item, true
	default:
		return zero, false // If the channel is empty, return nil
	}
}

// Poll implements Queue.
func (q *ChannelQueue[E]) Poll() (E, bool) {
	v := q.private.Load()
	if v != nil {
		q.private.Store(nil)
		atomic.AddInt64(&q.size, -1)
		return v.data, true
	}
	var zero E
	select {
	case item := <-q.ch: // Remove the element from the channel
		q.private.Store(nil)
		atomic.AddInt64(&q.size, -1)
		return item, true
	default:
		return zero, false // If the channel is empty, return nil
	}
}

// Size implements Queue.
func (q *ChannelQueue[E]) Size() int64 {
	return atomic.LoadInt64(&q.size) // Returns the number of elements in the channel
}

// ToSlice implements Queue.
func (q *ChannelQueue[E]) ToSlice() []E {
	panic("unimplemented")
}

// NewChannelQueue Create a new instance of ChannelQueue
func NewChannelQueue[E any]() *ChannelQueue[E] {
	return &ChannelQueue[E]{
		ch: make(chan E, segmentSize), // Initialize channel
	}
}

// IsEmpty Check whether the queue is empty
func (q *ChannelQueue[E]) IsEmpty() bool {
	return q.Size() == 0 // Check channel length
}

var _ Queue[any] = (*ChannelQueue[any])(nil)
