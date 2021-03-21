package customer_type_test

import (
	"testing"
	"fmt"
	"time"
)

type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv{
	return func(n int) int{
		start := time.Now()
		ret := inner(n)

		fmt.Println("time spent:",time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int{
	time.Sleep(time.Second*1)
	return op
}

/*
=== RUN   TestFn
time spent: 1.000347025
    customer_type_test.go:33: 10
*/
func TestFn(t *testing.T){
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}