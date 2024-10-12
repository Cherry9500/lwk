package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "10.213.232.207:8888") //拨号函数Dial
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	fmt.Println("conn suc=", conn)

	//客户端可以输入单行数据,然后再退出
	reader := bufio.NewReader(os.Stdin) //os.Stdin表示标准输入【终端】

	for {
		//从终端读入一行用户输入，并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readerString err=", err)
		}

		//如果用户输入的是exit就退出
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出")
			break
		}
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
		// fmt.Printf("客户端发送了 %d 字节的数据,并退出", n)
	}
}
