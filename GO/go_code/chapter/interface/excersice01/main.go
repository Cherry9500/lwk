package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero //Hero结构体的切片类型

func (hs HeroSlice) Len() int {
	return len(hs)
}

// 决定使用什么标准排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age //按年龄
	// return hs[i].Name < hs[j].Name
}

func (hs HeroSlice) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
	// hs[i],hs[j] = hs[j],hs[i] // 和上面等价
}

func main() {
	var intSlice = []int{0, -1, 10, 7, 90}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
	for _, v := range heroes {
		fmt.Println(v)
	}

	sort.Sort(heroes)
	fmt.Println("排序后----------")
	for _, v := range heroes {
		fmt.Println(v)
	}
}
