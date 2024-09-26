package queue

import (
	"fmt"
	"math/rand/v2"
	"testing"
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

func BenchmarkQueueMix(b *testing.B) {
	benchmarks := []struct {
		name string
		q    Queue[int]
	}{
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

func fastrand() uint32 {
	return rand.Uint32()
}
