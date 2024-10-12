package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	fileName := "d:/test.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}

	defer file.Close()

	var count CharCount
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough // 穿透处理
			case v >= 'A' && v <= 'z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Printf("字符的个数为=%v 数字的个数为%v 空格的个数为%v 其他字符的个数%v",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
