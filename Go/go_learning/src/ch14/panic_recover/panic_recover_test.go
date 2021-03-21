package panic_recover_test

import (
	"errors"
	"fmt"
	"testing"
)

/*
=== RUN   TestPanicVxExit
Start
recovered from  Something wrong!
*/
func TestPanicVxExit(t *testing.T) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
	//os.Exit(-1)
	//fmt.Println("End")
}