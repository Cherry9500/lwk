package main

import (
	"fmt"
	"goproject/go_code/customerManage/model"
	"goproject/go_code/customerManage/service"
)

type customerView struct {
	key  string //接收用户的输入
	loop bool

	customerService *service.CustomerService
}

// 显示所有的客户信息
func (lwk *customerView) list() {
	customers := lwk.customerService.List() //
	fmt.Println("-------------客户列表-------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n-------------客户列表完成-------------\n\n")
}

// 添加用户
func (lwk *customerView) add() {
	fmt.Println("----------添加客户---------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)

	//构建一个新Customer实例
	//id号需要系统自动分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if lwk.customerService.Add(customer) {
		fmt.Println("-----------添加完成----------")
	} else {
		fmt.Println("-----------添加失败----------")
	}
}

// 删除客户
func (lwk *customerView) delete() {
	fmt.Println("-----------删除客户----------")
	fmt.Println("请选择删除客户的编号(-1退出):")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除
	}

	fmt.Println("确认删除:(Y/N):")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if lwk.customerService.Delete(id) {
			fmt.Println("-----------删除完成------------")
		} else {
			fmt.Println("-----------输入的id不存在------------")
		}
	}
}

// 退出
func (lwk *customerView) exit() {
	fmt.Println("确认是否退出(Y/N):")
	for {
		fmt.Scanln(&lwk.key)
		if lwk.key == "y" || lwk.key == "n" || lwk.key == "Y" || lwk.key == "N" {
			break
		}
		fmt.Println("你的输入有误，确认是否退出(Y/N):")
	}
	if lwk.key == "Y" || lwk.key == "y" {
		lwk.loop = false
	}
}

func (lwk *customerView) mainMenu() {
	for {
		fmt.Println("----------客户信息管理软件---------")
		fmt.Println("          	1.添加用户")
		fmt.Println("          	2.修改用户")
		fmt.Println("          	3.删除用户")
		fmt.Println("          	4.客户列表")
		fmt.Println("          	5.退    出")
		fmt.Print("请选择(1-5):")

		fmt.Scanln(&lwk.key)
		switch lwk.key {
		case "1":
			lwk.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			lwk.delete()
		case "4":
			lwk.list()
		case "5":
			lwk.exit()
		default:
			fmt.Println("输入有误，请重新输入：")
		}
		if !lwk.loop {
			break
		}
	}
	fmt.Println("系统已退出")
}

func main() {
	customerView := customerView{
		key:  "",
		loop: true,
	}
	customerView.customerService = service.NewCustomerService() //初始化customerService字段
	customerView.mainMenu()
}
