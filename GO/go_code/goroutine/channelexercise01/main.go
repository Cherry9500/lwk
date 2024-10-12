package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 0; i < 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	// var num int
	var flag bool
	for {
		// time.Sleep(time.Millisecond * 10)
		num, ok := <-intChan
		if !ok {
			break
		}

		flag = true //假设是素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("有一个primeNum因为取不到数据退出")
	exitChan <- true

}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 8000)
	exitChan := make(chan bool, 4)

	start := time.Now().Unix()

	go putNum(intChan)

	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//开始处理主线程
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		end := time.Now().Unix()
		fmt.Println("使用协程耗时=", end-start)

		close(primeChan)
	}()

	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		// fmt.Printf("素数=%d\n", res)
	}
	fmt.Println("main线程已推出")
}
