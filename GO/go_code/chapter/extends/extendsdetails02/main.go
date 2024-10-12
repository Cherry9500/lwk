package main

import (
	"fmt"
)

type A struct {
	Name string
	age  int
}

type B struct {
	Name  string
	score float64
}

type C struct {
	A
	B
}

// 组合关系
type D struct {
	a A // 有名结构体
}

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

func main() {
	var c C
	c.A.Name = "tom"
	fmt.Println("c")

	var d D
	d.a.Name = "jack"

	tv1 := TV{Goods{"电视机01", 5000.99}, Brand{"海尔", "山东"}}
	tv2 := TV{
		Goods{
			Price: 10000.99,
			Name:  "电视剧02",
		},
		Brand{
			Name:    "夏普",
			Address: "北京",
		},
	}
	fmt.Println("tv1", tv1.Brand, tv1.Goods)
	fmt.Println("tv2", tv2.Brand, tv2.Goods)

	tv3 := TV2{&Goods{"电视机03", 7000.99}, &Brand{"创维", "河南"}}
	tv4 := TV2{
		&Goods{
			Name:  "电视剧04",
			Price: 10000.99,
		},
		&Brand{
			Name:    "长虹",
			Address: "四川",
		},
	}
	fmt.Println("tv3", *tv3.Brand, *tv3.Goods)
	fmt.Println("tv4", *tv4.Brand, *tv4.Goods)
}
