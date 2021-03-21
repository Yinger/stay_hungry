package slice_test

import (
	"testing"
)

/*
=== RUN   TestSliceInit
    slice_test.go:19: 0 0
    slice_test.go:22: 1 1
    slice_test.go:25: 4 4
    slice_test.go:28: 3 5
    slice_test.go:30: 0 0 0
    slice_test.go:32: 0 0 0 1
    slice_test.go:33: 4 5
*/
func TestSliceInit(t *testing.T){
	var s0 []int
	t.Log(len(s0),cap(s0))

	s0 = append(s0,1)
	t.Log(len(s0),cap(s0))

	s1:=[]int{1,2,3,4}
	t.Log(len(s1),cap(s1))

	s2 := make([]int ,3 ,5)
	t.Log(len(s2),cap(s2))
	// t.Log(s2[0],s2[1],s2[2],s2[3])
	t.Log(s2[0],s2[1],s2[2])
	s2 = append(s2,1)
	t.Log(s2[0],s2[1],s2[2],s2[3])
	t.Log(len(s2),cap(s2))
}

/*
=== RUN   TestSliceGrowing
    slice_test.go:53: 1 1
    slice_test.go:53: 2 2
    slice_test.go:53: 3 4
    slice_test.go:53: 4 4
    slice_test.go:53: 5 8
    slice_test.go:53: 6 8
    slice_test.go:53: 7 8
    slice_test.go:53: 8 8
    slice_test.go:53: 9 16
    slice_test.go:53: 10 16
*/
func TestSliceGrowing(t *testing.T){
	s := []int{}
	for i := 0;i<10;i++{
		s = append(s,i)
		t.Log(len(s),cap(s))
	}
}

/*
=== RUN   TestSliceShareMemory
    slice_test.go:67: [Apr May Jun] 3 9
    slice_test.go:69: [Jun Jul Aug] 3 7
    slice_test.go:71: [Apr May Unknow]
    slice_test.go:72: [Jan Feb Mar Apr May Unknow Jul Aug Sep Oct Nov Dec]
*/
func TestSliceShareMemory(t *testing.T){
	year:=[]string{"Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"}
	Q2:=year[3:6]
	t.Log(Q2,len(Q2),cap(Q2)) // [Apr May Jun] 3 9
	summer:=year[5:8]
	t.Log(summer,len(summer),cap(summer)) //[Jun Jul Aug] 3 7
	summer[0] = "Unknow"
	t.Log(Q2) //[Apr May Unknow]
	t.Log(year) //[Jan Feb Mar Apr May Unknow Jul Aug Sep Oct Nov Dec]
}

//slice can only be compared to nil
// func TestSliceComparing(t *testing.T){
// 	a:=[]int{1,2,3,4}
// 	b:=[]int{1,2,3,4}
// 	if a == b {
// 		t.Log("equal")
// 	}
// }