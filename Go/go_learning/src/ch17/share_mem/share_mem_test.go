package share_mem_test

import (
	"sync"
	"testing"
	"time"
)

/*
=== RUN   TestCounter
    share_mem_test.go:22: counter = 4991
*/
func TestCounter(t *testing.T) {

	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)

}

/*
=== RUN   TestCounterThreadSafe
    share_mem_test.go:43: counter = 4988
*/
func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	// time.Sleep(1 * time.Second) //去掉這行代碼結果不等5000 因爲外層的執行速度比内部協程執行的快
	t.Logf("counter = %d", counter)

}

/*
=== RUN   TestCounterWaitGroup
    share_mem_test.go:67: counter = 5000
*/
func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)

}