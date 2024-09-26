package queue

type ChannelQueue[E any] struct {
	ch chan E // 添加一个通道字段
}

// Clear implements Queue.
func (q *ChannelQueue[E]) Clear() {
	for !q.IsEmpty() {
		<-q.ch // 清空通道中的所有元素
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
	case q.ch <- item: // 尝试将元素发送到通道
		return true
	default:
		return false // 如果通道已满，返回 false
	}
}

// Peek implements Queue.
func (q *ChannelQueue[E]) Peek() (E, bool) {
	var zero E
	select {
	case item := <-q.ch: // 尝试从通道中取出元素
		q.ch <- item // 重新放回通道
		return item, true
	default:
		return zero, false // 如果通道为空，返回 nil
	}
}

// Poll implements Queue.
func (q *ChannelQueue[E]) Poll() (E, bool) {
	var zero E
	select {
	case item := <-q.ch: // 从通道中取出元素
		return item, true
	default:
		return zero, false // 如果通道为空，返回 nil
	}
}

// Size implements Queue.
func (q *ChannelQueue[E]) Size() int64 {
	return int64(len(q.ch)) // 返回通道中的元素数量
}

// ToSlice implements Queue.
func (q *ChannelQueue[E]) ToSlice() []E {
	panic("unimplemented")
}

// NewChannelQueue 创建一个新的 ChannelQueue 实例
func NewChannelQueue[E any]() *ChannelQueue[E] {
	return &ChannelQueue[E]{
		ch: make(chan E, segmentSize), // 初始化通道
	}
}

// IsEmpty 检查队列是否为空
func (q *ChannelQueue[E]) IsEmpty() bool {
	return len(q.ch) == 0 // 检查通道长度
}

var _ Queue[any] = (*ChannelQueue[any])(nil)
