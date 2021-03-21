package loop_test

import (
	"testing"
)

/*
=== RUN   TestWhileLoop
    loop_test.go:18: 0
    loop_test.go:18: 1
    loop_test.go:18: 2
    loop_test.go:18: 3
    loop_test.go:18: 4
*/
func TestWhileLoop(t *testing.T){
	n := 0
	for n < 5 {
		t.Log(n)
		n ++
	}
}