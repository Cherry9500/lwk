package main

import (
	"fmt"
)

type A struct {
	Name string
	age  int
}

func (a *A) Sayok() {
	fmt.Println("A Sayok", a.Name, a.age)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name, a.age)
}

type B struct {
	A
	Name string
}

func (b *B) Sayok() {
	fmt.Println("B Sayok", b.Name, b.age)
}

func main() {
	// var b B
	// b.A.Name = "tom"
	// b.A.age = 78
	// b.A.Sayok()
	// b.A.hello()

	// b.Name = "smith"
	// b.age =20
	// b.Sayok()
	// b.hello()

	var b B
	b.Name = "jack"
	b.A.Name = "scott"
	b.age = 20
	b.Sayok()
	b.A.Sayok()
	b.hello()
}
