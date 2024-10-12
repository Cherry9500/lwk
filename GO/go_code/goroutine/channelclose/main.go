package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan) //close之后可以读取，但是不能再写入数据

	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}

	//管道的遍历
	close(intChan2) //需要先关闭管道
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}
