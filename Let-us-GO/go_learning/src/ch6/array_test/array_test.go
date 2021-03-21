package array_test

import (
	"testing"
)

/*
=== RUN   TestArrayInit
    array_test.go:17: 0 0
    array_test.go:18: [1 5 3 4] [1 2 3]
*/
func TestArrayInit(t *testing.T){
	var arr [3]int
	arr1 := [4]int{1,2,3,4}
	arr3 := [...]int{1,2,3}
	arr1[1] = 5
	t.Log(arr[1],arr[2])
	t.Log(arr1,arr3)
}

/*
=== RUN   TestArrayTravel
    array_test.go:39: 1
    array_test.go:39: 3
    array_test.go:39: 4
    array_test.go:39: 5
    array_test.go:43: 0 1
    array_test.go:43: 1 3
    array_test.go:43: 2 4
    array_test.go:43: 3 5
    array_test.go:47: 1
    array_test.go:47: 3
    array_test.go:47: 4
    array_test.go:47: 5
*/
func TestArrayTravel(t *testing.T){
	arr3 := [...]int{1,3,4,5}
	for i :=0;i<len(arr3);i++{
		t.Log(arr3[i])
	}

	for idx,e := range arr3{
		t.Log(idx,e)
	}

	for _,e:= range arr3{
		t.Log(e)
	}
}

/*
=== RUN   TestArraySection
    array_test.go:60: [1 2 3]
    array_test.go:63: [4 5]
    array_test.go:66: [1 2 3 4 5]
*/
func TestArraySection(t *testing.T){
	arr3 := [...]int{1,2,3,4,5}
	arr3_sec := arr3[:3]
	t.Log(arr3_sec)

	arr3_sec = arr3[3:]
	t.Log(arr3_sec)

	arr3_sec = arr3[:]
	t.Log(arr3_sec)
}