# Queue

```
名称                                         次数(26550976)       平均时间(202.1 ns/op)     平均内存使用量 (7 B/op) 平均分配次数 (0 allocs/op)
BenchmarkQueue/LockFreeQueue/Offer-20           26550976               202.1 ns/op             7 B/op          0 allocs/op
```
# BenchmarkQueue/LockFreeQueue/Offer-20
- 名称：表示当前基准测试的名称。
- BenchmarkQueue 是基准测试的顶级名称。
- LockFreeQueue 表示使用的队列类型。
- Offer 表示具体的测试操作（这里是入队操作）。
- -20 表示运行时的 GOMAXPROCS 设置（即并行度）。
# 26550976
- 次数：表示该基准测试运行了多少次。
- 这个数字表示在给定时间内（默认为 1 秒），该测试能够执行的操作次数。
# 202.1 ns/op
- 平均时间：表示每次操作的平均耗时。
- ns 表示纳秒（1 纳秒 = 10^-9 秒）。
- op 表示一次操作（operation）。
- 因此，202.1 ns/op 表示每次入队操作的平均耗时为 202.1 纳秒。
# 7 B/op
- 平均内存使用量：表示每次操作分配的平均字节数。
- B 表示字节（byte）。
- op 表示一次操作。
- 因此，7 B/op 表示每次入队操作平均分配了 7 字节的内存。
# 0 allocs/op
- 平均分配次数：表示每次操作分配的内存次数。
- allocs 表示分配次数。
- op 表示一次操作。
- 因此，0 allocs/op 表示每次入队操作没有进行任何内存分配。
# 总结
- 次数 (26550976)：表示该测试运行了多少次。
- 平均时间 (202.1 ns/op)：表示每次入队操作的平均耗时。
- 平均内存使用量 (7 B/op)：表示每次入队操作平均分配的字节数。
- 平均分配次数 (0 allocs/op)：表示每次入队操作的内存分配次数。
- 这些指标可以帮助你理解不同操作的性能表现，从而进一步优化代码。例如，如果发现某个操作的平均时间较长或内存使用较多，可以针对性地进行优化。

# Benchmarking
```
$>go test -v -race -benchmem -cpu="8,16,32" -benchtime=5s -run=^$ -bench ^BenchmarkQueue$ -cpuprofile=cpu -memprofile=mem github.com/origadmin/toolkits/queue
goos: windows
goarch: amd64
pkg: github.com/origadmin/toolkits/queue
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkQueue
BenchmarkQueue
BenchmarkQueue/LockFreeQueue/Offer
BenchmarkQueue/LockFreeQueue/Offer-8            31714936               199.1 ns/op             8 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Offer-16           26100540               205.9 ns/op             8 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Offer-32           31186132               226.5 ns/op             8 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Poll
BenchmarkQueue/LockFreeQueue/Poll-8             23684212               279.5 ns/op             0 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Poll-16            24003744               250.0 ns/op             0 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Poll-32            25795556               260.0 ns/op             0 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Mixed
BenchmarkQueue/LockFreeQueue/Mixed-8             7564692               858.1 ns/op             9 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Mixed-16            6900072               922.3 ns/op            12 B/op          0 allocs/op
BenchmarkQueue/LockFreeQueue/Mixed-32            6511975               956.4 ns/op            10 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Offer
BenchmarkQueue/MutexQueue/Offer-8               13904630               440.2 ns/op            41 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Offer-16              14502745               462.0 ns/op            37 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Offer-32              10905456               643.8 ns/op            77 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Poll
BenchmarkQueue/MutexQueue/Poll-8                21881016               774.1 ns/op             0 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Poll-16               18971287               668.8 ns/op             0 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Poll-32                9894604               874.3 ns/op             0 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Mixed
BenchmarkQueue/MutexQueue/Mixed-8                5919752              1240 ns/op              81 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Mixed-16               5668646              1077 ns/op               0 B/op          0 allocs/op
BenchmarkQueue/MutexQueue/Mixed-32               5773432               935.0 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/origadmin/toolkits/queue     220.216s
```

# Benchmarking
```
$>go test -v -race -cpu=8 -benchmem -run=^$ -benchtime=5s -bench ^BenchmarkQueueMix$ github.com/origadmin/toolkits/queue
goos: windows
goarch: amd64
pkg: github.com/origadmin/toolkits/queue
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkQueueMix
BenchmarkQueueMix/DesignQueue/Mixed
BenchmarkQueueMix/DesignQueue/Mixed-8            4249044              1578 ns/op              20 B/op          1 allocs/op
BenchmarkQueueMix/WrapQueue/Mixed
BenchmarkQueueMix/WrapQueue/Mixed-8              1000000             39670 ns/op            1497 B/op         93 allocs/op
BenchmarkQueueMix/ChanQueue/Mixed
BenchmarkQueueMix/ChanQueue/Mixed-8              8007457               653.5 ns/op             0 B/op          0 allocs/op
BenchmarkQueueMix/LockFreeQueue/Mixed
BenchmarkQueueMix/LockFreeQueue/Mixed-8          5856188              1031 ns/op               9 B/op          0 allocs/op
BenchmarkQueueMix/MutexQueue/Mixed
BenchmarkQueueMix/MutexQueue/Mixed-8             4987773              1212 ns/op              20 B/op          0 allocs/op
PASS
ok      github.com/origadmin/toolkits/queue     69.761s
```

// Running tool: go test -v -race -count=5 -benchmem -run=^$ -bench ^BenchmarkQueue$ -parallel=32 github.com/origadmin/toolkits/queue
// Running tool: go test -v -race -benchmem -cpu="8,16,32" -benchtime=5s -run=^$ -bench ^BenchmarkQueue$ -cpuprofile=cpu -memprofile=mem github.com/origadmin/toolkits/queue
```shell
go test -v -race -benchmem -run=^$ -bench ^BenchmarkQueue$ -benchtime=30s -cpuprofile=cpu.out -memprofile=mem.out -parallel=32 github.com/origadmin/toolkits/queue
go test -v -race -benchmem -cpu="8" -run=^$ -bench ^BenchmarkQueueMix$ -benchtime=30s -cpuprofile=cpu -memprofile=mem github.com/origadmin/toolkits/queue

go tool pprof --http=:8080 cpu.out
go tool pprof --http=:8081 mem.out
```


