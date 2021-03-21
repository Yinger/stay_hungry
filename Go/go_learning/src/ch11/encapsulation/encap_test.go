package encap

import (
	"unsafe"
	"fmt"
	"testing"
)

type Employee struct {
	ID string
	Name string
	Age int
}

/*
=== RUN   TestCreateEmployeeObj
Address is c0000923a0    encap_test.go:32: ID:0-Name:Bob-Age:20
Address is c000092430    encap_test.go:33: ID:-Name:Mike-Age:30
    encap_test.go:34: 
Address is c000092460    encap_test.go:35: ID:2-Name:Rose-Age:22
    encap_test.go:36: e is encap.Employee
    encap_test.go:37: e2 is *encap.Employee
*/
func TestCreateEmployeeObj(t *testing.T){
	e := Employee{"0","Bob",20}
	e1 := Employee{Name:"Mike",Age:30}
	e2 := new(Employee) //返回指针
	e2.ID = "2"
	e2.Age = 22
	e2.Name = "Rose"

	t.Log(e)
	t.Log(e1)
	t.Log(e1.ID)
	t.Log(e2)
	t.Logf("e is %T",e)
	t.Logf("e2 is %T",e2)
}

func (e Employee) String() string {
	fmt.Printf("Address is %x",unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.ID,e.Name,e.Age)
}

func (e *Employee) ToString() string {
	fmt.Printf("Address is %x",unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.ID,e.Name,e.Age)
}

/*
=== RUN   TestStructOperations
Address is c0000924f0    encap_test.go:60: ID:0-Name:Bob-Age:20
Address is c0000924c0    encap_test.go:61: ID:0-Name:Bob-Age:20
Address is c0000924c0Address is c000092550    encap_test.go:65: ID:0-Name:Bob-Age:20
Address is c000092520    encap_test.go:66: ID:0-Name:Bob-Age:20
Address is c000092520--- PASS: TestStructOperations (0.00s)
*/
func TestStructOperations(t *testing.T) {
	e := Employee{"0","Bob",20}
	t.Log(e.String())
	t.Logf(e.ToString())
	fmt.Printf("Address is %x",unsafe.Pointer(&e.Name))

	e2 := &Employee{"0","Bob",20} //不用再拷贝，推荐
	t.Log(e2.String())
	t.Logf(e2.ToString())
	fmt.Printf("Address is %x",unsafe.Pointer(&e2.Name))
}