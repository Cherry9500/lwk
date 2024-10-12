package main

import "fmt"

type Student struct {
}

func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型,值是%v\n", index, x)
		case float32:
			fmt.Printf("第%v个参数是float32类型,值是%v\n", index, x)
		case float64:
			fmt.Printf("第%v个参数是float64类型,值是%v\n", index, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数是int类型,值是%v\n", index, x)
		case string:
			fmt.Printf("第%v个参数是string类型,值是%v\n", index, x)
		case Student:
			fmt.Printf("第%v个参数是-结构体-类型,值是%v\n", index, x)
		case *Student:
			fmt.Printf("第%v个参数是-指针-类型,值是%v\n", index, x)
		default:
			fmt.Printf("第%v个参数的类型不确定,值是%v\n", index, x)
		}
	}
}

func main() {
	var n1 float32 = 1.1
	n2 := 2.3
	n3 := 30
	name := "tom"
	address := "北京"
	n4 := 300

	stu1 := Student{}
	stu2 := &Student{}

	TypeJudge(n1, n2, n3, name, address, n4, stu1, stu2)
}
