package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	// var allChan chan interface{}
	allChan := make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{"花猫", 4}
	allChan <- cat

	//希望获得到管道中的三个元素，需要先将前两个推出
	<-allChan
	<-allChan

	newCat := <-allChan

	fmt.Printf("newCat=%T , newCat=%v\n",newCat,newCat)
	// fmt.Printf("newCat.Name=%v",newCat.Name) //编译器不识别,要使用类型断言
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v",a.Name)



}