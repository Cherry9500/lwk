package main

import (
	"bufio"
	"fmt"
	"os"
)

// 在原基础上更改10行hello world
func main() {
	filePath := "d:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	defer file.Close()

	str := "hello world\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	//因为writer是带缓存的，因此在调用WriteStrin方法时，其实内容是先写入到缓存的
	//需要用Flush方法将缓存的数据真正写入到文件中，否则将没有数据
	writer.Flush()
}
