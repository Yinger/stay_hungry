package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

func cancel2(cancelChan chan struct{}) {
	close(cancelChan)
}

/*
=== RUN   TestCancel
4 Cancelled
0 Cancelled
1 Cancelled
2 Cancelled
3 Cancelled
*/
func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}
	// cancel1(cancelChan) //4 Cancelled
	cancel2(cancelChan) 
	time.Sleep(time.Second * 1)
}