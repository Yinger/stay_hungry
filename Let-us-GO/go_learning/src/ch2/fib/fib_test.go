package fib

import (
	"fmt"
	"testing"
)

/*
=== RUN   TestFibList
1 2 3 5 8 13
*/
func TestFibList(t *testing.T){
	var a int = 1
	var b int = 2

	fmt.Print(a)
	for i:=0; i < 5; i++{
		fmt.Print(" ",b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}

/*
=== RUN   TestExchange
    fib_test.go:37: 2 1
    fib_test.go:40: 1 2
*/
func TestExchange(t *testing.T){
	a := 1
	b := 2
	tmp := a
	a = b
	b = tmp
	t.Log(a,b)

	a,b = b,a
	t.Log(a,b)
}