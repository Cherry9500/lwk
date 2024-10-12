package main

import "fmt"

//接口里的方法必须全部实现，接口是引用类型，接收的是地址
type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}
func (p Phone) Call() {
	fmt.Println("开始打电话...")
}

type Camera struct {
	Name string
}

func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

type Computer struct {
}

//接收一个usb接口类型的变量
func (computer Computer) Working(usb Usb) { //usb这个接口变量体现了多态的特性：多态参数
	usb.Start()
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	var usbArr [3]Usb
	usbArr[0] = Phone{"小米"}
	usbArr[1] = Phone{"华为"}
	usbArr[2] = Camera{"尼康"}

	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}
}
