package main

import (
	"fmt"
	"customerManage/service"
	"customerManage/model"
)

type customerView struct {
	key             string // 输入
	loop            bool   // 是否loop显示菜单
	customerService *service.CustomerService
}

// 显示客户信息
func (this *customerView) list() []model.Customer {
	customerList := this.customerService.List()
	fmt.Println("-------------------------List-------------------------")
	for i := 0; i < len(customerList); i++ {
		fmt.Println(customerList[i].GetInfo())
	}
	fmt.Println()
	return customerList
}

// add客户
func (this *customerView) add() {

	fmt.Println("-------------------------Add-------------------------")
	fmt.Println("name:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("gender:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("age:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("phone:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("email:")
	email := ""
	fmt.Scanln(&email)
	// 构建新的Customer实例
	newNustomer := model.NewNotIDCustomer(name, gender, age, phone, email)
	if this.customerService.Add(newNustomer) {
		fmt.Println("-------------------------Add Over-----------------------")
	}
}

// delete by id
func (this *customerView) delete() {
	fmt.Println("-------------------------Del-------------------------")
	fmt.Println("Delete ID:")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("Sure Delete(y/n):")
	// 这里加入判断直到为yes执行，或者为no退出
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if this.customerService.DeleteById(id) {
			fmt.Println("Delete Success...")
		} else {
			fmt.Println("Delete Fail, ID not exists")
		}
	}
}

func (this *customerView) exit() {
	fmt.Println("Sure exit(y/n).")
	for {
		fmt.Scanln(&this.key)
		if this.key == "y" || this.key == "Y" || this.key == "n" || this.key == "N" {
			break
		}
		fmt.Println("Input Error...")
	}
	if this.key == "y" || this.key == "Y" {
		this.loop = false
	}
}

// display
func (this *customerView) mainMenu() {
	for {
		fmt.Println("-------------------------Info-------------------------")
		fmt.Println("1. create customer")
		fmt.Println("2. update customer")
		fmt.Println("3. delete customer")
		fmt.Println("4. show customer")
		fmt.Println("5. exit")
		fmt.Println("please input num: ")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		//fmt.Println("create")
		case "2":
			fmt.Println("update")
		case "3":
			this.delete()
		//fmt.Println("delete")
		case "4":
			this.list()
		case "5":
			this.exit()
		//this.loop = false
		default:
			fmt.Println("input error, please once input.")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("Exit......")

}

func main() {
	customerView := customerView{
		key:"",
		loop:true,
	}
	// 完成对customerService结构体的初始化
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
	fmt.Println("ok")
}
