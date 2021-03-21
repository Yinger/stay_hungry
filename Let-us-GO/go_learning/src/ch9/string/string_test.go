package string_test

import (
	"testing"
)

/*
=== RUN   TestString
    string_test.go:20: 
    string_test.go:22: 5
    string_test.go:26: 亻�
    string_test.go:27: 4
    string_test.go:29: 3
    string_test.go:32: 1
    string_test.go:34: 中 unicode 4e2d
    string_test.go:35: 中 UTF8 e4b8ad
*/
func TestString(t *testing.T) {
	var s string
	t.Log(s) //初始化为默认零值“”
	s = "hello"
	t.Log(len(s))
	//s[1] = '3' //string是不可变的byte slice
	//s = "\xE4\xB8\xA5" //可以存储任何二进制数据
	s = "\xE4\xBA\xBB\xFF"
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(len(s)) //是byte数

	c := []rune(s)
	t.Log(len(c))
	//	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

/*
=== RUN   TestStringToRune
    string_test.go:51: 中 4e2d
    string_test.go:51: 华 534e
    string_test.go:51: 人 4eba
    string_test.go:51: 民 6c11
    string_test.go:51: 共 5171
    string_test.go:51: 和 548c
    string_test.go:51: 国 56fd
*/
func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}