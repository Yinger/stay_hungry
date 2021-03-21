package empty_interface_test

import (
	"testing"
	"fmt"
)

func DoSomething(p interface{}) {
	// if i,ok  := p.(int) ; ok {
	// 	fmt.Println("Integer ", i)
	// 	return
	// }
	// if s,ok := p.(string); ok{
	// 	fmt.Println("string",s)
	// 	return
	// }
	// fmt.Println("Unkonw Type")
	switch v := p.(type) {
	case int :
		fmt.Println("Integer ", v)
	case string :
		fmt.Println("string",v)
	default:
		fmt.Println("Unkonw Type")
	}
}

/*
=== RUN   TestEmptyInterfaceAssertion
Integer  10
string Hi
Unkonw Type
*/
func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("Hi")
	DoSomething(3.14)
}

