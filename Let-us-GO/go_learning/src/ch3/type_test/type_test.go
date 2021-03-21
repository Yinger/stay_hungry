package type_test

import (
	"testing"
)

type MyInt int

/*
=== RUN   TestImplicit
    type_test.go:22: 1 1 1
*/
func TestImplicit(t *testing.T){
	var a int = 1
	var b int64 
	// b = a cannot use a (type int) as type int64 in assignment
	b = int64(a)

	var c MyInt
	// c = b cannot use a (type int) as type int64 in assignment
	c = MyInt(b)
	t.Log(a,b,c)
}

/*
=== RUN   TestPoint
    type_test.go:34: 1 0xc000016108
    type_test.go:35: int *int
*/
func TestPoint(t *testing.T){
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1 invalid operation (go 不支持指针运算)
	t.Log(a, aPtr)
	t.Logf("%T %T",a,aPtr)
}

/*
=== RUN   TestString
    type_test.go:45: **
    type_test.go:46: 0
*/
func TestString(t *testing.T){
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))

	if s == ""{

	}

	// if s == nil{

	// }
}