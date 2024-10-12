package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b) //获取指针变量

	rVal.Elem().SetInt(20) //Elem()获取指针指向的变量，SetInt()更新变量的值
}

func main() {
	var num int = 10
	reflect01(&num)
	fmt.Println("num=", num)
}
