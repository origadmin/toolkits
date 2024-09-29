package queue

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

func TestConcurrentWrite(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	wg := sync.WaitGroup{}
	writes := int64(0)
	reads := int64(0)
	numWriters := 1000
	numReaders := 1000
	itemsPerWriter := 10000

	// 初始化等待组
	wg.Add(numWriters)

	// 使用 sync.Map 记录所有写入的数据
	allItems := sync.Map{}

	for i := 0; i < numWriters; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < itemsPerWriter; j++ {
				item := id*itemsPerWriter + j + 1
				for !queue.Offer(item) {
					runtime.Gosched()
				}
				atomic.AddInt64(&writes, 1)
				//allItems.Store(item, struct{}{})
			}
		}(i)
	}

	// 等待所有写入完成
	wg.Wait()

	for i := 0; i < numReaders; i++ {
		for j := 0; j < itemsPerWriter; j++ {
			var v int
			var ok bool
			for {
				v, ok = queue.Poll()
				if ok {
					atomic.AddInt64(&reads, 1)
					allItems.Store(v, struct{}{})
					break
				}
				_ = v
				runtime.Gosched()
			}
		}
	}

	// 检查写入数量是否正确
	if writes != int64(numWriters*itemsPerWriter) {
		t.Errorf("Expected %d writes, but got %d", numWriters*itemsPerWriter, writes)
	}

	// 检查所有数据是否已写入
	itemCount := 0
	var missingItems []int
	allItems.Range(func(key, _ interface{}) bool {
		itemCount++
		if _, ok := allItems.Load(key); !ok {
			missingItems = append(missingItems, key.(int))
		}
		return true
	})

	if itemCount != int(writes) {
		t.Errorf("Expected %d unique items, but got %d", writes, itemCount)
	}
	t.Logf("Missing items: %v", missingItems)
}

func TestConcurrentRead(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	wg := sync.WaitGroup{}
	writes := int64(0)
	reads := int64(0)
	numWriters := 1000
	numReaders := 1000
	itemsPerWriter := 10000

	// 先写入数据
	for i := 0; i < numWriters; i++ {
		for j := 0; j < itemsPerWriter; j++ {
			item := i*itemsPerWriter + j + 1
			for !queue.Offer(item) {
				runtime.Gosched()
			}
			atomic.AddInt64(&writes, 1)
		}
	}

	// 初始化等待组
	wg.Add(numReaders)

	// 使用 sync.Map 记录所有读取的数据
	readItems := sync.Map{}

	for i := 0; i < numReaders; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < itemsPerWriter; j++ {
				var v int
				var ok bool
				for {
					v, ok = queue.Poll()
					if ok {
						atomic.AddInt64(&reads, 1)
						readItems.Store(v, struct{}{})
						break
					}
					runtime.Gosched()
				}
			}
		}()
	}

	// 等待所有读取完成
	wg.Wait()

	// 检查读取数量是否正确
	if reads != writes {
		t.Errorf("Expected %d reads, but got %d", writes, reads)
	}

	// 检查所有数据是否已读取
	itemCount := 0
	var missingItems []int
	readItems.Range(func(key, _ interface{}) bool {
		itemCount++
		if _, ok := readItems.Load(key); !ok {
			missingItems = append(missingItems, key.(int))
		}
		return true
	})

	if itemCount != int(writes) {
		t.Errorf("Expected %d unique items, but got %d", writes, itemCount)
	}
	t.Logf("Missing items: %v", missingItems)
}

func TestMixedConcurrentOperations(t *testing.T) {
	queue := NewLockFreeQueue[int]()
	wg := sync.WaitGroup{}
	writes := int64(0)
	reads := int64(0)
	numWriters := 100
	numReaders := 100
	itemsPerWorker := 1000

	// 初始化等待组
	wg.Add(numWriters + numReaders)

	// 使用 sync.Map 记录所有写入的数据
	allItems := sync.Map{}

	// 写入数据
	for i := 0; i < numWriters; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < itemsPerWorker; j++ {
				item := id*itemsPerWorker + j
				for !queue.Offer(item) {
					runtime.Gosched()
				}
				atomic.AddInt64(&writes, 1)
				allItems.Store(item, struct{}{})
			}
		}(i)
	}

	// 使用 sync.Map 记录所有读取的数据
	readItems := sync.Map{}

	// 读取数据
	for i := 0; i < numReaders; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < itemsPerWorker; j++ {
				var v int
				var ok bool
				for {
					v, ok = queue.Poll()
					if ok {
						atomic.AddInt64(&reads, 1)
						readItems.Store(v, struct{}{})
						break
					}
					runtime.Gosched()
				}
			}
		}()
	}

	// 等待所有操作完成
	wg.Wait()

	// 检查写入数量是否正确
	if writes != int64(numWriters*itemsPerWorker) {
		t.Errorf("Expected %d writes, but got %d", numWriters*itemsPerWorker, writes)
	}

	// 检查读取数量是否正确
	if reads != writes {
		t.Errorf("Expected %d reads, but got %d", writes, reads)
	}

	// 检查所有数据是否已读取
	itemCount := 0
	readItems.Range(func(_, _ interface{}) bool {
		itemCount++
		return true
	})

	if itemCount != int(writes) {
		t.Errorf("Expected %d unique items, but got %d", writes, itemCount)
	}

	// 检查所有写入的数据是否都被读取
	allItems.Range(func(key, _ interface{}) bool {
		if _, ok := readItems.Load(key); !ok {
			t.Errorf("Item %v was not read", key)
		}
		return true
	})
}
