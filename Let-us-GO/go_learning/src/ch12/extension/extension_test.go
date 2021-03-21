package extension

import (
	"testing"
	"fmt"
)

type Pet struct {

}

func (p *Pet) Speak() {
	fmt.Print("...")
}


func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ",host)
}

type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	fmt.Print("Wang!")
}


func (d *Dog) SpeakTo(host string) {
	d.Speak()
	fmt.Println(" ",host)
}

/*
=== RUN   TestDog
Wang!  Lin
*/
func TestDog (t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Lin")
}

type Cat struct {
	Pet
}

func (c *Cat) Speak() {
	fmt.Print("Miao!")
}

/*
=== RUN   TestCat
...  Lin
*/
func TestCat (t *testing.T) {
	cat := new(Cat)
	cat.SpeakTo("Lin")
}


