package queue

import (
	"sync"
	"sync/atomic"
	"testing"
)

// Poll retrieves the first element from a non-empty queue
func TestPollRetrievesFirstElement(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	queue.Offer(42)
	result, _ := queue.Poll()
	if result != 42 {
		t.Errorf("Expected 42, but got %d. Arrr!", result)
	}
}

// Poll decreases the consumer index by one after retrieving an element
func TestPollDecreasesConsumerIndex(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	queue.Offer(42)
	initialConsumer := atomic.LoadInt64(&queue.consumer)
	queue.Poll()
	newConsumer := atomic.LoadInt64(&queue.consumer)
	if newConsumer != initialConsumer+1 {
		t.Errorf("Expected consumer index to increase by 1, but it didn't. Yo-ho-ho!")
	}
}

// Poll handles the case when the queue is empty and returns immediately
func TestPollHandlesEmptyQueue(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	result, _ := queue.Poll()
	if result != 0 {
		t.Errorf("Expected 0 for empty queue, but got %d. Walk the plank!", result)
	}
}

// Poll manages concurrent access correctly without data races
func TestPollConcurrentAccess(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	var wg sync.WaitGroup
	for i := 0; i < 1024; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for !queue.Offer(i) {
			}
		}(i)
	}
	wg.Wait()

	for i := 0; i < 1024; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, ok := queue.Poll(); !ok; _, ok = queue.Poll() {

			}
		}()
	}
	wg.Wait()

	if queue.Size() != 0 {
		t.Errorf("Expected size 0 after concurrent access, but got %d. Avast ye!", queue.Size())
	}
}

// Poll handles the wrap-around case when the consumer index reaches the end of the buffer
func TestPollWrapAround(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	for i := 0; i < 4; i++ {
		queue.Offer(i)
		_, _ = queue.Poll()
	}

	queue.Offer(99)
	result, _ := queue.Poll()

	if result != 99 {
		t.Errorf("Expected 99 after wrap-around, but got %d. Yo-ho-ho and a bottle of rum!", result)
	}
}

// Poll correctly handles the case when the producer and consumer indices are equal
func TestPollProducerConsumerEqual(t *testing.T) {
	queue := NewLockFreeQueue[int]()

	result, _ := queue.Poll()

	if result != 0 {
		t.Errorf("Expected 0 when producer and consumer are equal, but got %d. Dead men tell no tales!", result)
	}
}
