package client

import (
	"ch15/series"
	"testing"
)

/*
init1
init2
=== RUN   TestPackage
    package_test.go:16: [1 1 2 3 5]
    package_test.go:17: 25
*/
func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSerie(5))
	t.Log(series.Square(5))
}