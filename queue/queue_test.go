package queue

import (
	"fmt"
	"sync/atomic"
	"testing"
)

// Running tool: go test -v -race -count=5 -benchmem -run=^$ -bench ^BenchmarkQueue$ github.com/origadmin/toolkits/queue
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
				q.Offer(i)
			}
		})

		b.Run(bm.name+"/Poll", func(b *testing.B) {
			q := bm.q
			for i := 0; i < b.N; i++ {
				q.Offer(i)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				q.Poll()
			}
		})

		b.Run(bm.name+"/Mixed", func(b *testing.B) {
			q := bm.q
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					if fastrand()%2 == 0 {
						q.Offer(int(fastrand()))
					} else {
						q.Poll()
					}
				}
			})
		})
	}
}

var r uint32 = 1

func fastrand() uint32 {
	return atomic.AddUint32(&r, 1664525)
}
