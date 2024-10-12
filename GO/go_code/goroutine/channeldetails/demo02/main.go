package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}

func test() {
	//捕获panic保证其它协程不受报错协程的影响
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test()发生错误", err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "golang" //没有make直接使用报错了
}

func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}
}
