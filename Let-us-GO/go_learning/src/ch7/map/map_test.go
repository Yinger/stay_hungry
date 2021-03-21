package map_test

import (
	"testing"
)

/*
=== RUN   TestInitMap
    map_test.go:16: 4
    map_test.go:17: len m1=3
    map_test.go:21: len m2=1
    map_test.go:24: len m3=0
*/
func TestInitMap(t *testing.T){
	m1:=map[int]int{1:1,2:4,3:9}
	t.Log(m1[2])
	t.Logf("len m1=%d",len(m1)) //3

	m2:=map[int]int{}
	m2[4]=16
	t.Logf("len m2=%d",len(m2)) //1

	m3:=make(map[int]int,10)
	t.Logf("len m3=%d",len(m3)) //0
}

/*
=== RUN   TestAccessNotExistingKey
    map_test.go:35: 0
    map_test.go:37: 0
    map_test.go:42: key 3 is not existing
*/
func TestAccessNotExistingKey(t *testing.T){
	m1 := map[int]int{}
	t.Log(m1[1]) //0
	m1[2] = 0
	t.Log(m1[2]) //0

	if v,ok := m1[3]; ok{
		t.Logf("Key 3's value is %d",v)
	}else{
		t.Log("key 3 is not existing")
	}
}

/*
=== RUN   TestTravelMap
    map_test.go:55: 1 1
    map_test.go:55: 2 4
    map_test.go:55: 3 9
*/
func TestTravelMap(t *testing.T){
	m1:=map[int]int{1:1,2:4,3:9}
	for k,v:=range m1 {
		t.Log(k, v)
	}
}