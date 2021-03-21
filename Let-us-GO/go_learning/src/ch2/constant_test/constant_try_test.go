package constant_test

import (
	"testing"
)

const(
	Monday = iota + 1
	Tuesday
	Wednesday
)

const(
	Readable = 1 << iota
	Writable
	Executable
)

/*
=== RUN   TestConstantTry
    constant_try_test.go:24: 1 2
*/
func TestConstantTry(t *testing.T){
	t.Log(Monday,Tuesday)
}

/*
=== RUN   TestConstantTry1
    constant_try_test.go:33: true true trues
*/
func TestConstantTry1(t *testing.T){
	a:=7 //0111
	t.Log(a&Readable == Readable,a&Writable == Writable,a&Executable == Executable)
}