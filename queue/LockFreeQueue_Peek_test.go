package queue

import (
	"sync"
	"testing"
)

// Peek returns the correct element when the queue is not empty
func TestPeekReturnsCorrectElement(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	queue.Offer(42)
	if got, _ := queue.Peek(); got != 42 {
		t.Errorf("Peek() = %v, want %v", got, 42)
	}
}

// Peek does not remove the element from the queue
func TestPeekDoesNotRemoveElement(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	queue.Offer(42)
	queue.Peek()
	if got := queue.Size(); got != 1 {
		t.Errorf("Size() = %v, want %v", got, 1)
	}
}

// Peek returns the first element added to the queue
func TestPeekReturnsFirstElement(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	queue.Offer(42)
	queue.Offer(43)
	if got, _ := queue.Peek(); got != 42 {
		t.Errorf("Peek() = %v, want %v", got, 42)
	}
}

// Peek returns zero value when the queue is empty
func TestPeekReturnsZeroWhenEmpty(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	if got, _ := queue.Peek(); got != 0 {
		t.Errorf("Peek() = %v, want %v", got, 0)
	}
}

// Peek handles concurrent access correctly
func TestPeekHandlesConcurrentAccess(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			queue.Offer(i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			queue.Peek()
		}
	}()

	wg.Wait()

	if got := queue.Size(); got != 1000 {
		t.Errorf("Size() = %v, want %v", got, 1000)
	}
}

// Peek works correctly when the queue is full
func TestPeekWorksWhenFull(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	for i := 0; i < 4; i++ {
		queue.Offer(i)
	}
	if got, _ := queue.Peek(); got != 0 {
		t.Errorf("Peek() = %v, want %v", got, 0)
	}
}

// Peek should be tested with different data types
func TestPeekWithDifferentDataTypes(t *testing.T) {
	intQueue := NewLockFreeQueue[int]()
	intQueue.Offer(42)
	if got, _ := intQueue.Peek(); got != 42 {
		t.Errorf("Peek() = %v, want %v", got, 42)
	}

	stringQueue := NewLockFreeQueue[string]()
	stringQueue.Offer("Ahoy!")
	if got, _ := stringQueue.Peek(); got != "Ahoy!" {
		t.Errorf("Peek() = %v, want %v", got, "Ahoy!")
	}
}
