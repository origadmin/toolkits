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
		{"LockFreeRetryQueue", NewLockFreeQueue[int]().WithRetry(1024)},
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
							runtime.Gosched()
						}
					} else {
						// Read operation
						_, ok := q.Poll()
						if !ok {
							runtime.Gosched()
						}
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
func TestPollReturnsElement(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	writes := int64(0)
	reads := int64(0)

	per := 4096
	mm := map[int]int{}
	for j := 0; j < per; j++ {
		v := j
		for !queue.Offer(v + 1) {
		}
		atomic.AddInt64(&writes, 1)
		mm[v+1] = v + 1
	}

	func(sta int) {
		for i := 0; i < per; i++ {
			v, ok := queue.Poll()
			for !ok {
				v, ok = queue.Poll()
			}
			atomic.AddInt64(&reads, 1)
			if v != mm[v] {
				t.Errorf("Expected %d, but got %d. Ye-ho!", mm[v], v)
			}
		}
	}(0)

	if reads != writes {
		t.Errorf("Expected reads(%d) and writes(%d) to be equal, but they aren't. Ye-ho! ", reads, writes)
	}
	var duplicates []int
	var missings []int
	for i := 0; i < per; i++ {
		if _, ok := mm[i+1]; !ok {
			missings = append(missings, i+1)
		}
	}

	fmt.Println("All done!", reads, writes, queue.Size())
	fmt.Println("Duplicates: ", duplicates)
	fmt.Println("Missings: ", missings)
}

// Poll returns the correct element when multiple elements are in the queue
func TestPollReturnsCorrectElement(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	wg := sync.WaitGroup{}
	writes := int64(0)
	reads := int64(0)

	max := 64
	per := 64
	wg.Add(max)
	mm := map[int]int{}
	for i := 0; i < max*per; i++ {
		mm[i+1] = 0
	}
	for i := 0; i < max; i++ {
		go func(sta int) {
			defer wg.Done()
			for j := 0; j < per; j++ {
				v := sta*per + j
				for !queue.Offer(v + 1) {
					runtime.Gosched()
				}
				atomic.AddInt64(&writes, 1)
			}
		}(i)
	}

	wg.Wait()

	for i := 0; i < max*per; i++ {
		v, ok := queue.Poll()
		mm[v] = 0
		if !ok {
			t.Error("Yo-ho-ho! Offer didn't manage buffer wrap-around correctly!", v, queue.consumer, queue.producer)
		}
	}

	//if reads != writes {
	//	t.Errorf("Expected reads(%d) and writes(%d) to be equal, but they aren't. Ye-ho! ", reads, writes)
	//}

	ti := map[int]struct{}{}
	poolMax := 0
	var duplicates []int
	var missings []int
	for k, v := range mm {
		if v != 0 {
			missings = append(missings, k)
		}
	}

	fmt.Println("All done!", poolMax, reads, writes, queue.Size())
	fmt.Println("cursor:", len(ti))
	fmt.Println("Duplicates: ", duplicates)
	fmt.Println("Missings: ", missings)
}

// Poll returns the correct element when multiple elements are in the queue
func TestPollReturnsCorrectOffer(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	wg := sync.WaitGroup{}
	writes := int64(0)

	max := 1024
	per := 1024

	mm := sync.Map{}
	for i := 0; i < max*per; i++ {
		mm.Store(i+1, 1)
	}
	wg.Add(max)
	for i := 0; i < max; i++ {
		go func(sta int) {
			defer wg.Done()
			for j := 0; j < per; j++ {
				v := sta*per + j
				for !queue.Offer(v + 1) {
					runtime.Gosched()
				}
				atomic.AddInt64(&writes, 1)
			}
		}(i)
	}

	wg.Wait()

	ti := map[int]struct{}{}
	poolMax := 0
	var duplicates []int
	mm.Range(func(k, v any) bool {
		poolMax++
		vv := k.(int)
		if _, ok := ti[vv]; ok {
			duplicates = append(duplicates, vv)
			return true
		}
		if vvv := v.(int); vvv == 1 {
			ti[vv] = struct{}{}
		}
		return true
	})

	var missings []int
	var indexes []int
	seg := queue.getHeadSegment()
	consumer := int64(0)
	producer := queue.getProducer()

Loop:
	for seg != nil {
		for idx, v := range seg.buffer {
			if v.data == 0 {
				indexes = append(indexes, idx)
				_ = idx
				break
			}
			if consumer >= producer {
				break Loop
			}
			consumer++
			if _, ok := ti[v.data]; !ok {
				missings = append(missings, v.data)
			} else {
				delete(ti, v.data)
			}
		}
		if seg.nextSegment() == nil {
			break Loop
		}
		seg = seg.nextSegment()
	}

	//for i := 0; i < max*per; i++ {
	//	_, truth := mm.Load(i + 1)
	//	if !truth {
	//		missings = append(missings, i+1)
	//	}
	//}

	fmt.Println("All done!", poolMax, writes, queue.Size())
	fmt.Println("Duplicates: ", duplicates)
	//fmt.Println("Missings: ", ti)
	fmt.Println("Missings: ", missings)
	//fmt.Println("Indexes: ", indexes)
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
