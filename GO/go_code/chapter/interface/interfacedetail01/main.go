package main

import (
	"fmt"
)

type AInterface interface {
	Say()
}

type Stu struct {
	Name string
}

// type integer int //非结构体类型，不如int也可以实现接口
// func (i integer) Say() {
// 	fmt.Println("integer Say i =", i)
// }


func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}

func main() {
	var stu Stu
	var a AInterface = stu
	a.Say()

	// var i integer = 10
	// var b AInterface = i
	// b.Say()
}
