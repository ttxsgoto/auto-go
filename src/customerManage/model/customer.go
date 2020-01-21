package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// 工厂模式
func NewCustomer(id int, name string, gender string,
age int, phone string, email string) Customer {
	return Customer{
		Id:id,
		Name:name,
		Gender:gender,
		Age:age,
		Phone:phone,
		Email:email,
	}
}

func NewNotIDCustomer(name string, gender string,
age int, phone string, email string) Customer {
	return Customer{
		Name:name,
		Gender:gender,
		Age:age,
		Phone:phone,
		Email:email,
	}
}



func (this *Customer) GetInfo() (info string) {
	info = fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",
		this.Id, this.Name, this.Gender,
		this.Age, this.Phone, this.Email)
	return
}






