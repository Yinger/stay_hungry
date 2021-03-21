package string_test

import (
	"strconv"
	"strings"
	"testing"
)

/*
=== RUN   TestStringFn
    string_fun_test.go:20: A
    string_fun_test.go:20: B
    string_fun_test.go:20: C
    string_fun_test.go:22: A-B-C
*/
func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

/*
=== RUN   TestConv
    string_fun_test.go:32: str10
    string_fun_test.go:34: 20
*/
func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}