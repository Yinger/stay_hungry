package condition_test

import (
	"testing"
)

/*
=== RUN   TestIfMultiSec
    condition_test.go:13: 1 == 1
*/
func TestIfMultiSec(t *testing.T){
	if a := 1 == 1; a{
		t.Log("1 == 1")
	}
}

/*
=== RUN   TestSwitchMultiCase
    condition_test.go:29: Even
    condition_test.go:31: Odd
    condition_test.go:29: Even
    condition_test.go:31: Odd
    condition_test.go:33: it is not 0-3
*/
func TestSwitchMultiCase(t *testing.T){
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2 :
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

/*
=== RUN   TestSwitchCaseCondition
    condition_test.go:50: Even
    condition_test.go:52: Odd
    condition_test.go:50: Even
    condition_test.go:52: Odd
    condition_test.go:50: Even
*/
func TestSwitchCaseCondition(t *testing.T){
	for i := 0; i < 5; i++ {
		switch {
		case i % 2 == 0 :
			t.Log("Even")
		case i % 2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknow")
		}
	}
}