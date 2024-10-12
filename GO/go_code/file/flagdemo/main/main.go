package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port int

	//&user 就是接受用户命令行中输入的 -u 后面的参数值
	//"u"，就是 -u 指定参数
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "locahost", "主机名,默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号,默认为3306")

	flag.Parse() //转换方法，必须调用

	fmt.Printf("user=%v pwd=%v host=%v port=%v", user, pwd, host, port)
}
