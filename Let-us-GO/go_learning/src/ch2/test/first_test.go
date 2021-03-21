package try_test

import (
	"testing"
)

/*
=== RUN   TestFirstTry
    first_test.go:12: first try
*/
func TestFirstTry(t *testing.T){
	t.Log("first try")
}

//add go.testFlags": ["-v"] to settings.json
//add go.coverOnSave true to settings.json