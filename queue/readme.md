

# Benchmarking
```
PS D:\workspace\project\golang\sugoitech\toolkits> go test -v -race -count=5 -benchmem -run=^$ -bench ^BenchmarkQueue$ github.com/origadmin/toolkits/queue
goos: windows
goarch: amd64
pkg: github.com/origadmin/toolkits/queue
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkQueue
BenchmarkQueue
BenchmarkQueue/LockFreeQueue/Offer
BenchmarkQueue/LockFreeQueue/Offer-20            1826245               650.4 ns/op            16 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Offer-20            2007847               611.0 ns/op            16 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Offer-20            2014292               627.6 ns/op            16 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Offer-20            1970569               627.9 ns/op            16 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Offer-20            2009721               641.5 ns/op            16 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Poll
BenchmarkQueue/LockFreeQueue/Poll-20             3332791               308.4 ns/op             0 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Poll-20             4130342               302.4 ns/op             0 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Poll-20             4178796               288.6 ns/op             0 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Poll-20             4407406               292.1 ns/op             0 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Poll-20             4437904               289.0 ns/op             0 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Mixed
BenchmarkQueue/LockFreeQueue/Mixed-20            2069326               622.4 ns/op             7 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Mixed-20            2027505               675.1 ns/op             7 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Mixed-20            2016572               684.5 ns/op             7 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Mixed-20            2084904               605.3 ns/op             7 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Mixed-20            1881807               579.5 ns/op             7 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Offer
BenchmarkQueue/MutexQueue/Offer-20               3260104              1201 ns/op              48 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Offer-20               3141787               356.5 ns/op            45 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Offer-20               3273115               687.0 ns/op            67 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Offer-20               3546742               715.6 ns/op            54 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Offer-20               2894822               516.6 ns/op            83 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Poll
BenchmarkQueue/MutexQueue/Poll-20                2590372               422.3 ns/op             0 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Poll-20                2624344               414.1 ns/op             0 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Poll-20                2528770               418.6 ns/op             0 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Poll-20                2498575               426.7 ns/op             0 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Poll-20                2630511               853.6 ns/op             0 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Mixed
BenchmarkQueue/MutexQueue/Mixed-20               1337218              1021 ns/op               0 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Mixed-20               1379344               965.8 ns/op             0 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Mixed-20               1385858              1085 ns/op             238 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Mixed-20               1376232               975.0 ns/op             0 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Mixed-20               1393788               959.2 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/origadmin/toolkits/queue     112.831s
```