package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

/*
=== RUN   TestGroutine
2
0
9
3
1
5
8
6
7
4
*/
func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			//time.Sleep(time.Second * 1)
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 50)
}