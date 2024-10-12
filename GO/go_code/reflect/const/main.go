package main

import "fmt"

func main() {
	var num int
	//常量在声明的时候必须给值,常量不能修改，只能修饰bool，数值(int、float系列)类型，string类型
	const tax int = 0
	fmt.Println(num, tax)
}
