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
	//d2 = a        // 不能使用值类型(引发panic)， 存储在接口中的值是无法寻址的，因此编译器无法自动获取指针地址引发panic
	                // 赋值的对象必须实现对应的接口方法，而a没有实现对应的方法，而是a的指针实现了方法
	d2.Describe()   // go会把指针进行隐式转换得到value，反过来不行，因为只通过值不能找到唯一的指针地址，如都为整数2可能有多个地址
	a.Describe()    // 指针变量调用指针方法
}



















