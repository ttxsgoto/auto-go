package main

import "fmt"

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age int
}

func (p Person) Describe()  {
	fmt.Printf("%s %d years\n", p.name, p.age)
	
}

type Address struct {
	state string
	country string
}

func (a *Address) Describe()  { //指针实现方法，需要传递指针过来
	fmt.Printf("State %s Country %s\n", a.state, a.country)
}

func main()  {
	
	var d1 Describer
	p1 := Person{name:"ttxsgoto", age:23}
	d1 = p1
	d1.Describe()
	
	p2 := Person{name: "test", age:25}
	d1 = &p2
	d1.Describe()
	
	var d2 Describer
	a :=Address{state:"chengdu", country:"sichuan"}
	d2 = &a
	//d2 = a      // 不能使用值类型(引发panic)， 存储在接口中的值是无法寻址的，因此编译器无法自动获取指针地址引发panic
	d2.Describe()
	a.Describe()    // 指针变量调用指针方法
	
}



















