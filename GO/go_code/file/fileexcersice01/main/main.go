package main

import (
	"fmt"
	"io/ioutil"
)

// 将D盘的abc导入到C盘的lwk
func main() {
	file1Path := "d:/abc.txt"
	file2Path := "c:/lwk.txt"
	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Println("read file err=%v", err)
		return
	}
	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Printf("read file err=%v\n", err)
	}

}
