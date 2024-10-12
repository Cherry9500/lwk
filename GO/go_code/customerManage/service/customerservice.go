package service

import (
	"goproject/go_code/customerManage/model"
)

// 增删改查的操作
type CustomerService struct {
	customers []model.Customer // 一个切片

	customerNum int //+1做为id
}

// 编写一个方法，返回*CustomerService一个指针
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@sohu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 返回客户切片
func (lwk *CustomerService) List() []model.Customer {
	return lwk.customers
}

func (lwk *CustomerService) Add(customer model.Customer) bool { //一定要要用指针来绑定
	//确定一个分配id的规则为添加的顺序
	lwk.customerNum++
	customer.Id = lwk.customerNum
	lwk.customers = append(lwk.customers, customer)
	return true
}

// 根据id查早在切片中对应的下标,没有返回-1
func (lwk *CustomerService) FindById(id int) int {
	index := -1
	//遍历
	for i := 0; i < len(lwk.customers); i++ {
		if lwk.customers[i].Id == id {
			index = i
		}
	}
	return index
}

func (lwk *CustomerService) Delete(id int) bool {
	index := lwk.FindById(id)
	if index == -1 {
		return false
	}
	//从切片中删除一个元素
	lwk.customers = append(lwk.customers[:index], lwk.customers[index+1:]...)
	return true

}

// 修改客户信息
// func (lwk *CustomerService) Update(id int) bool {
// 	index := lwk.FindById(id)
// 	if index == -1 {
// 		return false
// 	}
// 	lwk.customers = append(lwk.customers, )
// 	return true
// }
