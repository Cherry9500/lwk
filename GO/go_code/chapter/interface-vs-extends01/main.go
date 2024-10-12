package main

import "fmt"

type Monkey struct {
	Name string
}

type Littlemonkey struct {
	Monkey
}

type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (mon *Monkey) Climbing() {
	fmt.Println(mon.Name, "爬树...")
}

func (mon *Littlemonkey) Flying() {
	fmt.Println(mon.Name, "会飞...")
}

func (mon *Littlemonkey) Swimming() {
	fmt.Println(mon.Name, "会游泳...")
}

func main() {
	monkey := Littlemonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.Climbing()
	monkey.Flying()
	monkey.Swimming()
}
