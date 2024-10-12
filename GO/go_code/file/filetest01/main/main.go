package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//文件是指针类型
	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	defer file.Close() //要及时关闭，否则会有内存泄漏

	//创建一个带缓存的 *reader
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //读到表示文件的末尾
			break
		}
		fmt.Println(str)
	}
	fmt.Println("文件读取结束...")
}
