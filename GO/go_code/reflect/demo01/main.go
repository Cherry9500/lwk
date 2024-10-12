package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	//先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//再获取reflect.Value
	rVal := reflect.ValueOf(b)

	//获取变量对应的kind
	fmt.Printf("kind=%v kind=%v\n",rVal.Kind(),rTyp.Kind())

	// n2 := 2 + rVal.Int()
	// fmt.Println("n2=",n2)

	fmt.Printf("rVal=%v type=%T\n", rVal, rVal)

	iV := rVal.Interface() //将反射对象还原成接口类型的变量
	num2 := iV.(int) //通过断言转换
	fmt.Println("num2", num2)
}

func reflectTest02(b interface{}) {
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	rVal := reflect.ValueOf(b)
	fmt.Printf("kind=%v kind=%v\n",rVal.Kind(),rTyp.Kind())

	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T\n", iV, iV)

	stu,ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n",stu.Name)
	}
}

type Student struct {
	Name string
	Age  int
}

func main() {
	var num int = 100
	reflectTest01(num)

	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)
}
