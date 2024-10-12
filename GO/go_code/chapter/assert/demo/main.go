package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	var b Point
	// b = a
	b = a.(Point) //类型断言
	fmt.Println(b)

	// var x interface{}
	// var c float32 = 1.1
	// x = c
	// y := x.(float32) //不能是float64,类型一定要匹配
	// fmt.Printf("y的类型是%T 值是=%v",y,y)

	//带检测的类型断言
	var x interface{}
	var c float32 = 1.1
	x = c

	// y, ok := x.(float32)
	if y, ok := x.(float32); ok {
		fmt.Println("convert success")
		fmt.Printf("y的类型是%T 值是=%v", y, y)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println("继续执行...")
}
