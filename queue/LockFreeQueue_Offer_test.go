package queue

import (
	"sync/atomic"
	"testing"
)

// Successfully adds an element to the queue when there is space
func TestOfferAddsElementWhenSpaceAvailable(t *testing.T) {
	queue := NewLockFreeQueue[int]()

	success := queue.Offer(42)
	if !success {
		t.Error("Arrr! Failed to add element when there be space!")
	}
}

// Correctly updates the producer index after adding an element
func TestOfferUpdatesProducerIndex(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	queue.Offer(42)
	if atomic.LoadInt64(&queue.producer) != 1 {
		t.Error("Shiver me timbers! getProducer index not updated correctly!")
	}
}

// Returns true when an element is successfully added
func TestOfferReturnsTrueOnSuccess(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	success := queue.Offer(42)
	if !success {
		t.Error("Blimey! Offer didn't return true when it should have!")
	}
}

// Correctly handles the scenario when the queue is empty
func TestOfferHandlesEmptyQueue(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	success := queue.Offer(42)
	if !success || queue.Size() != 1 {
		t.Error("Ahoy! Offer didn't handle empty queue correctly!")
	}
}

// Properly manages the wrap-around of the buffer using the mask
func TestOfferManagesWrapAround(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	for i := 0; i < 4096; i++ {
		queue.Offer(i + 1)
	}
	if queue.Size() != 4096 {
		t.Error("Yo-ho-ho! Offer didn't manage buffer wrap-around correctly!")
	}

	for i := 0; i < 4096; i++ {
		v, ok := queue.Poll()
		if !ok || v != i+1 {
			t.Error("Yo-ho-ho! Offer didn't manage buffer wrap-around correctly!", v, queue.consumer, queue.producer)
		}
	}

	success := queue.Offer(42)
	if p, _ := queue.Poll(); !success || p != 42 {
		t.Error("Walk the plank! Offer didn't manage buffer wrap-around correctly!")
	}
}
