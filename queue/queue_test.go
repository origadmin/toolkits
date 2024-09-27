package queue

import (
	"fmt"
	"math/rand/v2"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/yireyun/go-queue"
	"golang.design/x/lockfree"
)

// Running tool: go test -v -race -count=5 -benchmem -run=^$ -bench ^BenchmarkQueue$ github.com/origadmin/toolkits/queue
// go test -v -race -count=5 -benchmem -run=^$ -bench ^BenchmarkQueue$ -benchtime=5s -cpuprofile=cpu.out -memprofile=mem.out github.com/origadmin/toolkits/queue
// go tool pprof --http=:8080 cpu.out
// go tool pprof --http=:8081 mem.out
func BenchmarkQueue(b *testing.B) {
	benchmarks := []struct {
		name string
		q    Queue[int]
	}{
		{"LockFreeQueue", NewLockFreeQueue[int]()},
		{"MutexQueue", NewMutexQueue[int]()},
		{"WrapQueue", NewYireyunQueue[int]()},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name+"/Offer", func(b *testing.B) {
			q := bm.q
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for !q.Offer(i) {
				}
			}
		})

		b.Run(bm.name+"/Poll", func(b *testing.B) {
			q := bm.q
			for i := 0; i < b.N; i++ {
				for !q.Offer(i) {
				}
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, ok := q.Poll(); !ok; _, ok = q.Poll() {
				}
			}
		})

		b.Run(bm.name+"/Mixed", func(b *testing.B) {
			q := bm.q

			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					if fastrand()%2 == 0 {
						// Write operation
						for !q.Offer(int(fastrand())) {
						}
					} else {
						// Read operation
						_, _ = q.Poll()
					}
				}
			})
		})
	}
}

// Running tool: go test -v -race -cpu=8 -benchmem -run=^$ -bench ^BenchmarkQueueMix$ github.com/origadmin/toolkits/queue
func BenchmarkQueueMix(b *testing.B) {
	benchmarks := []struct {
		name string
		q    Queue[int]
	}{
		//{"PoolQueueTest", NewPoolQueue[int]()}, // result value is not correct
		{"DesignQueueTest", NewDesignQueue[int]()},
		//{"YireyunQueueTest", NewYireyunQueue[int]()}, // too slow removed
		{"ChanQueueTest", NewChannelQueue[int]()},
		{"LockFreeQueue", NewLockFreeQueue[int]()},
		{"MutexQueue", NewMutexQueue[int]()},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name+"/Mixed", func(b *testing.B) {
			q := bm.q

			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					if fastrand()%2 == 0 {
						// Write operation
						cnt := 0
						for !q.Offer(int(fastrand())) && cnt < 1000 {
							cnt++
						}
					} else {
						// Read operation
						_, _ = q.Poll()
					}
				}
			})
		})
	}
}

func fastrand() uint32 {
	time.Sleep(1)
	return rand.Uint32()
}

type yireyunQueue[E any] struct {
	q *queue.EsQueue
}

func (w *yireyunQueue[E]) Offer(e E) bool {
	ok, _ := w.q.Put(e)
	return ok
}

func (w *yireyunQueue[E]) Poll() (E, bool) {
	var zero E
	val, ok, _ := w.q.Get()
	if !ok {
		return zero, false
	}
	return val.(E), ok
}

func (w *yireyunQueue[E]) Peek() (E, bool) {
	//TODO implement me
	panic("implement me")
}

func (w *yireyunQueue[E]) Size() int64 {
	return int64(w.q.Capaciity())
}

func (w *yireyunQueue[E]) IsEmpty() bool {
	return w.Size() == 0
}

func (w *yireyunQueue[E]) Clear() {
	//TODO implement me
	panic("implement me")
}

func (w *yireyunQueue[E]) ToSlice() []E {
	//TODO implement me
	panic("implement me")
}

func (w *yireyunQueue[E]) Iterator() Iterator[E] {
	//TODO implement me
	panic("implement me")
}

func NewYireyunQueue[E any]() *yireyunQueue[E] {
	return &yireyunQueue[E]{
		q: queue.NewQueue(segmentSize),
	}
}

type designQueue[E any] struct {
	queue *lockfree.Queue
}

func (d *designQueue[E]) Offer(e E) bool {
	d.queue.Enqueue(e)
	return true
}

func (d *designQueue[E]) Poll() (E, bool) {
	dequeue := d.queue.Dequeue()
	if dequeue == nil {
		var zero E
		return zero, false
	}
	return dequeue.(E), true
}

func (d *designQueue[E]) Peek() (E, bool) {
	//TODO implement me
	panic("implement me")
}

func (d *designQueue[E]) Size() int64 {
	//TODO implement me
	panic("implement me")
}

func (d *designQueue[E]) IsEmpty() bool {
	//TODO implement me
	panic("implement me")
}

func (d *designQueue[E]) Clear() {
	//TODO implement me
	panic("implement me")
}

func (d *designQueue[E]) ToSlice() []E {
	//TODO implement me
	panic("implement me")
}

func (d *designQueue[E]) Iterator() Iterator[E] {
	//TODO implement me
	panic("implement me")
}

func NewDesignQueue[E any]() *designQueue[E] {
	queue := lockfree.NewQueue()
	return &designQueue[E]{
		queue: queue,
	}
}

// Poll returns the correct element when multiple elements are in the queue
func TestPollReturnsCorrectElement(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	wg := sync.WaitGroup{}
	writes := int64(0)
	reads := int64(0)
	for i := 0; i < 128; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			for i := 0; i < 4096; i++ {
				for !queue.Offer(i) {
				}
				atomic.AddInt64(&writes, 1)
			}
		}()
	}

	for i := 0; i < 128; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 4096; i++ {
				for {
					_, ok := queue.Poll()
					if ok {
						atomic.AddInt64(&reads, 1)
						break
					}
				}
			}
		}()
	}

	wg.Wait()
	if reads != writes {
		t.Errorf("Expected reads(%d) and writes(%d) to be equal, but they aren't. Ye-ho! ", reads, writes)
	}
	fmt.Println("All done!", reads, writes)
}

// Poll returns the correct element when multiple elements are in the queue
func TestPollReturnsCorrectLockFreeQueue(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	wg := sync.WaitGroup{}
	writes := int64(0)
	reads := int64(0)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 2048; i++ {
			for !queue.Offer(i) {
			}
			atomic.AddInt64(&writes, 1)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 2048; i++ {
			var v int
			var ok bool
			for {
				v, ok = queue.Poll()
				if ok {
					if v != i {
						t.Errorf("Expected %d, but got %d. Ye-ho!", i, v)
					}
					atomic.AddInt64(&reads, 1)
					break
				}
				runtime.Gosched()
			}
		}
	}()

	wg.Wait()
	if reads != writes {
		t.Errorf("Expected reads(%d) and writes(%d) to be equal, but they aren't. Ye-ho! ", reads, writes)
	}
	fmt.Println("All done!")
}

// Poll returns the correct element when multiple elements are in the queue
func TestPollReturnsCorrectChanQueue(t *testing.T) {
	queue := NewChannelQueue[int]()
	wg := sync.WaitGroup{}
	writes := int64(0)
	reads := int64(0)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 2048; i++ {
			for !queue.Offer(i) {
			}
			atomic.AddInt64(&writes, 1)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 2048; i++ {
			var v int
			var ok bool
			for {
				v, ok = queue.Peek()
				if ok {
					if v != i {
						t.Errorf("Expected %d, but got %d. Ye-ho!", i, v)
					}
				}
				v, ok = queue.Poll()
				if ok {
					if v != i {
						t.Errorf("Expected %d, but got %d. Ye-ho!", i, v)
					}
					atomic.AddInt64(&reads, 1)
					break
				}
				runtime.Gosched()
			}
		}
	}()

	wg.Wait()
	if reads != writes {
		t.Errorf("Expected reads(%d) and writes(%d) to be equal, but they aren't. Ye-ho! ", reads, writes)
	}
	fmt.Println("All done!")
}

// Poll returns the correct element when multiple elements are in the queue
func TestPollReturnsCorrectPoolQueue(t *testing.T) {
	queue := NewPoolQueue[int]()
	wg := sync.WaitGroup{}
	writes := int64(0)
	reads := int64(0)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 2048; i++ {
			for !queue.Offer(i) {
			}
			atomic.AddInt64(&writes, 1)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 2048; i++ {
			var v int
			var ok bool
			for {
				//v, ok = queue.Peek()
				//if ok {
				//	if v != i {
				//		t.Errorf("Peek Expected %d, but got %d. Ye-ho!", i, v)
				//	}
				//}
				v, ok = queue.Poll()
				if ok {
					if v != i {
						t.Errorf("Expected %d, but got %d. Ye-ho!", i, v)
					}
					atomic.AddInt64(&reads, 1)
					break
				}
				runtime.Gosched()
			}
		}
	}()

	wg.Wait()
	if reads != writes {
		t.Errorf("Expected reads(%d) and writes(%d) to be equal, but they aren't. Ye-ho! ", reads, writes)
	}
	fmt.Println("All done!")
}

var _ Queue[int] = (*yireyunQueue[int])(nil)
var _ Queue[int] = (*designQueue[int])(nil)
