package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 要求, 服务端在注册rpc对象是, 能让编译期检测出 注册对象是否合法.

// 创建接口, 在接口中定义方法原型
// rpc封装---服务端封装 1.定义接口
type MyInterface interface {
	HelloWorld(string, *string) error
}

// 调用该方法时, 需要给 i 传参, 参数应该是实现了 HelloWorld 方法的类对象!
// rpc封装---服务端封装 2.封装注册服务方法
func RegisterService(i MyInterface) {
	rpc.RegisterName("hello", i)
}

// --------------------------------------------------------------------
// 定义类对象
type World struct {
}

// 绑定类方法
func (lwk *World) HelloWorld(name string, resp *string) error {
	*resp = name + " 你好!"
	return nil
	//return errors.New("未知的错误!")
}

func main() {

	// 1. 注册RPC服务, 绑定对象方法
	RegisterService(new(World))
	/*	err := rpc.RegisterName("hello", new(World))
		if err != nil {
			fmt.Println("注册 rpc 服务失败!", err)
			return
		}*/

	// 2. 设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	fmt.Println("开始监听 ...")
	// 3. 建立链接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Accept() err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("链接成功...")

	// 4. 绑定服务
	jsonrpc.ServeConn(conn)
	//rpc.ServeConn(conn)
}
