package main

import "fmt"

type BInterface interface {
	test01()
}

type CInterface interface {
	test02()
}

// 必须要实现接口里的所有方法
type AInterface interface {
	BInterface
	CInterface
	test03()
}

type Stu struct {
}

func (stu Stu) test01() {

}
func (stu Stu) test02() {

}
func (stu Stu) test03() {

}

//可以把任何一个变量赋给空接口
type T interface {
}

func main() {
	var stu Stu
	var a AInterface = stu
	a.test02()

	var t T = stu
	fmt.Println(t)
	var t2 interface{} = stu
	var num1 float64 =8.8
	t2 = num1 
	t = num1
	fmt.Println(t2,t)
}
