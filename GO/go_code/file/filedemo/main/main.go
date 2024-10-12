package main

import (
	"fmt"
	"os"
)

func main() {
	//文件是指针类型
	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	fmt.Printf("file=%v", file)
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
}
