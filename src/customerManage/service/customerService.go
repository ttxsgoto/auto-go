package service

import "customerManage/model"

// 对customer的操作，customer切片来存储数据
type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

// 返回CustomerService
func NewCustomerService() *CustomerService {
	CustomerService := &CustomerService{}
	CustomerService.customerNum = 1
	customer := model.NewCustomer(
		1, "ttxsgoto", "man", 30, "12312312311", "ttxsgoto@163.com",
	)
	CustomerService.customers = append(CustomerService.customers, customer)
	return CustomerService
}

// list
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

// add
func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum = len(this.customers) + 1
	//this.customerNum ++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

func (this *CustomerService) DeleteById(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	// del index data
	this.customers = append(this.customers[:index], this.customers[index + 1:]...)
	return true
}

// by id find customer
func (this *CustomerService) FindById(id int) int {
	index := -1
	//this.customers
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}








