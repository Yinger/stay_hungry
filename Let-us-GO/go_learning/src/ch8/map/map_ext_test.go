package map_test

import (
	"testing"
)

/*
=== RUN   TestMapWithFuncValue
    map_ext_test.go:17: 2 4 8
*/
func TestMapWithFuncValue(t *testing.T){
	m:=map[int]func(op int)int{}
	m[1] = func(op int) int {return op}
	m[2] = func(op int) int {return op * op}
	m[3] = func(op int) int {return op * op * op}

	t.Log(m[1](2),m[2](2),m[3](2)) // 2 4 8
}

/*
=== RUN   TestMapForSet
    map_ext_test.go:31: 1 is existing
    map_ext_test.go:37: 2
    map_ext_test.go:43: 1 is not existing
*/
func TestMapForSet(t *testing.T){
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	if mySet[n] {
		t.Logf("%d is existing",n)
	} else {
		t.Logf("%d is not existing",n)
	}

	mySet[3] = true
	t.Log(len(mySet)) //2
	delete(mySet,n)

	if mySet[n] {
		t.Logf("%d is existing",n)
	} else {
		t.Logf("%d is not existing",n)
	}
}

