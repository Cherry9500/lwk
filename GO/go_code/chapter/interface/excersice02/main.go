package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct {
	Name  string
	Age   int
	Score float64
}

type StudentSlice []Student //结构体的切片类型

func (ss StudentSlice) Len() int {
	return len(ss)
}

// 决定使用什么标准排序
func (ss StudentSlice) Less(i, j int) bool {
	return int(ss[i].Score) < int(ss[j].Score)
	// return ss[i].Name < ss[j].Name
}

func (ss StudentSlice) Swap(i, j int) {
	temp := ss[i]
	ss[i] = ss[j]
	ss[j] = temp
	// ss[i],ss[j] = ss[j],ss[i] // 和上面等价
}

func main() {
	var intSlice = []int{10, 90, 23, 54, 89, 60}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	var stu StudentSlice
	for i := 0; i < 10; i++ {
		student := Student{
			Name:  fmt.Sprintf("学生~%d", rand.Intn(100)),
			Age:   rand.Intn(100),
			Score: float64(rand.Intn(120)),
		}
		stu = append(stu, student)
	}
	for _, v := range stu {
		fmt.Println(v)
	}

	sort.Sort(stu)
	fmt.Println("排序后----------")
	for _, v := range stu {
		fmt.Println(v)
	}
}
