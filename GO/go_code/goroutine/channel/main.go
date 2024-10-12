package main

import "fmt"

func main() {
	// var chan2 chan<- int  //声明为只写管道
	// var chan3 <-chan int  //声明为只读管道
	var intChan chan int
	intChan = make(chan int, 3) //初始化

	//intChan是引用类型
	fmt.Printf("intChan 的值=%v\n  intChan本身的地址%v\n", intChan, &intChan)

	intChan <- 10
	num := 211
	intChan <- num

	fmt.Printf("channel len=%v  cap=%v\n", len(intChan), cap(intChan))

	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len=%v  cap=%v\n", len(intChan), cap(intChan))

}
