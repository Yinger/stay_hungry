package objectpool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

/*
=== RUN   TestSyncPool
Create a new object.
100
Create a new object.
100
*/
func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	runtime.GC() //GC 会清除sync.pool中缓存的对象
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
}

/*
=== RUN   TestSyncPoolInMultiGroutine
100
100
100
Create a new object.
10
Create a new object.
10
Create a new object.
10
Create a new object.
10
Create a new object.
10
Create a new object.
10
Create a new object.
10
*/
func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}