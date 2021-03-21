package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

/*
=== RUN   TestTypeAndValue
    reflect_test.go:16: int64 10
    reflect_test.go:17: int64
*/
func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}

/*
=== RUN   TestBasicType
Unknown *float64
*/
func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(&f)

}

/*
=== RUN   TestDeepEqual
    reflect_test.go:54: a==b? true
    reflect_test.go:60: s1 == s2? true
    reflect_test.go:61: s1 == s3? false
true
true
*/
func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	//	t.Log(a == b)
	t.Log("a==b?", reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}

	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3))

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	fmt.Println(c1 == c2)
	fmt.Println(reflect.DeepEqual(c1, c2))
}

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

/*
=== RUN   TestInvokeByName
    reflect_test.go:94: Name: value(Mike), Type(reflect.Value) 
    reflect_test.go:98: Tag:format normal
    reflect_test.go:102: Updated Age: &{1 Mike 1}
*/
func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	//按名字获取成员
	t.Logf("Name: value(%[1]v), Type(%[1]T) ", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field.")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").
		Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Updated Age:", e)
}