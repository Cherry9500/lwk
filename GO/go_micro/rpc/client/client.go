package main

import (
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// rpc封装---客户端封装 1.定义类
type Myclient struct {
	c *rpc.Client //client可以调用Call方法
}

// 要求, 服务端在注册rpc对象是, 能让编译期检测出 注册对象是否合法.

// 由于使用了 c 调用 Call, 因此需要初始化 c
// rpc封装---客户端封装 3.初始化客户端
func InitClient(addr string) Myclient {
	conn, _ := jsonrpc.Dial("tcp", addr)

	return Myclient{c: conn}
}

// 实现函数, 原型参照上面的 Interface来实现.
// rpc封装---客户端封装 2.绑定类方法
func (lwk *Myclient) HelloWorld(a string, b *string) error {

	// 参数1, 参照上面的 Interface , RegisterName 而来.  a :传入参数  b:传出参数.
	return lwk.c.Call("hello.HelloWorld", a, b)
}

//--------------------------------------------------------------------

// func main01() {
// 	// 1. 用 rpc 链接服务器 --Dial()
// 	//conn, err := rpc.Dial("tcp", "127.0.0.1:8800")
// 	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8800")
// 	if err != nil {
// 		fmt.Println("Dial err:", err)
// 		return
// 	}
// 	defer conn.Close()

// 	// 2. 调用远程函数
// 	var reply string // 接受返回值 --- 传出参数

// 	err = conn.Call("hello.HelloWorld", "李白", &reply)
// 	if err != nil {
// 		fmt.Println("Call:", err)
// 		return
// 	}

// 	fmt.Println(reply)
// }

// 结合 03-design.go 测试
func main() {
	myClient := InitClient("127.0.0.1:8800")

	var resp string

	err := myClient.HelloWorld("杜甫", &resp)
	if err != nil {
		fmt.Println("HelloWorld err:", err)
		return
	}

	fmt.Println(resp, err)
}
