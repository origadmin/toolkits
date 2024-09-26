package queue

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/yireyun/go-queue"
	"golang.design/x/lockfree"
)

// Running tool: go test -v -race -count=5 -benchmem -run=^$ -bench ^BenchmarkQueue$ github.com/origadmin/toolkits/queue
// go test -v -race -count=5 -benchmem -run=^$ -bench ^BenchmarkQueue$ -benchtime=5s -cpuprofile=cpu.out -memprofile=mem.out github.com/origadmin/toolkits/queue
// go tool pprof --http=:8080 cpu.out
// go tool pprof --http=:8081 mem.out
func BenchmarkQueue(b *testing.B) {
	fmt.Println("BenchmarkQueue")
	benchmarks := []struct {
		name string
		q    Queue[int]
	}{
		{"LockFreeQueue", NewLockFreeQueue[int]()},
		{"MutexQueue", NewMutexQueue[int]()},
		{"WrapQueue", NewWrapQueue[int]()},
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
		{"PoolQueueTest", NewPoolQueue[int]()},
		{"DesignQueueTest", NewDesignQueue[int]()},
		{"WrapQueueTest", NewWrapQueue[int]()},
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
	return rand.Uint32()
}

type wrapQueue[E any] struct {
	q *queue.EsQueue
}

func (w *wrapQueue[E]) Offer(e E) bool {
	ok, _ := w.q.Put(e)
	return ok
}

func (w *wrapQueue[E]) Poll() (E, bool) {
	var zero E
	val, ok, _ := w.q.Get()
	if !ok {
		return zero, false
	}
	return val.(E), ok
}

func (w *wrapQueue[E]) Peek() (E, bool) {
	//TODO implement me
	panic("implement me")
}

func (w *wrapQueue[E]) Size() int64 {
	return int64(w.q.Capaciity())
}

func (w *wrapQueue[E]) IsEmpty() bool {
	return w.Size() == 0
}

func (w *wrapQueue[E]) Clear() {
	//TODO implement me
	panic("implement me")
}

func (w *wrapQueue[E]) ToSlice() []E {
	//TODO implement me
	panic("implement me")
}

func (w *wrapQueue[E]) Iterator() Iterator[E] {
	//TODO implement me
	panic("implement me")
}

func NewWrapQueue[E any]() *wrapQueue[E] {
	return &wrapQueue[E]{
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

var _ Queue[int] = (*wrapQueue[int])(nil)
var _ Queue[int] = (*designQueue[int])(nil)
