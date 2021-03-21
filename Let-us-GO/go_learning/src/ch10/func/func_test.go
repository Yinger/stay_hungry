package func_test

import (
	"fmt"
	"time"
	"math/rand"
	"testing"
)

func returnMultiValues()(int,int){
	return rand.Intn(10),rand.Intn(20)
}

func timeSpent(inner func(op int) int) func (op int) int{
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
    func_test.go:38: 1 7
    func_test.go:41: 7
time spent: 1.000792736
    func_test.go:44: 10
*/
func TestFn(t *testing.T){
	a,b := returnMultiValues()
	t.Log(a,b)

	a,_ = returnMultiValues()
	t.Log(a)

	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func Sum(ops ...int) int{
	ret := 0
	for _, op := range ops{
		ret += op
	}

	return ret
}

/*
=== RUN   TestVarParam
    func_test.go:62: 10
    func_test.go:63: 15
*/
func TestVarParam(t *testing.T){
	t.Log(Sum(1,2,3,4))
	t.Log(Sum(1,2,3,4,5))
}

func Clear(){
	fmt.Println("Clear resources.")
}

/*
=== RUN   TestDefer
Start
Clear resources.
--- FAIL: TestDefer (0.00s)
panic: err [recovered]
	panic: err
*/
func TestDefer(t *testing.T){
	defer Clear()
	fmt.Println("Start")
	panic("err")
}