package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "d:/test.txt"
	content, err := ioutil.ReadFile(file) //打开和关闭操作都隐藏在readfile里面了
	if err != nil {
		fmt.Println("open file err=", err)
	}
	// fmt.Printf("%v", content)
	fmt.Printf("%v", string(content))

}
