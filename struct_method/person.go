package main

import (
	"strings"
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
	
}

func main() {
	var pers1 Person
	pers1.firstName = "ttxs"    // 直接赋值
	pers1.lastName = "goto"
	upPerson(&pers1)
	fmt.Printf("The name %s %s \n", pers1.firstName, pers1.lastName)
	
	pers2 := new(Person)    // new方法分配内存，返回的是指针
	pers2.firstName = "xiao"
	pers2.lastName = "jiuyue"
	(*pers2).lastName = "xx"    // 反解指针方式设置值
	upPerson(pers2)
	fmt.Printf("The Name %s %s \n", pers2.firstName, pers2.lastName)
	
	pers3 := &Person{"x", "jy"}
	upPerson(pers3)
	fmt.Printf("The Name %s %s \n", pers3.firstName, pers3.lastName)
	
}

