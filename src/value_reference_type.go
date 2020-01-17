package main

import "fmt"
/*
- 值类型和引用类型(指针类型)
- 函数参数中使用
	- 当为值类型变量作为参数传递给函数时，值类型作为一个拷贝传递进去，在方法内部的改变都不会
	  影响到外面的变量
	- 当传递的是指针,内存地址没有发生变化，作为参数的方法内部也是一份新的拷贝，在方法内部的
	  修改可以影响到外部变量
	- 值类型和指针类型在方法内部都会产生一份拷贝，指向不同，指针类型的拷贝会产生一个新的地址，
	  这个地址和原数据地址指向的是同一份数据，所以修改后外部的变量会得到新的数据
 */
var arr = [5]int{1, 2, 3, 4, 5 }
var slice_arr = []int{1, 2, 3, 4, 5}
var mystring string = "ttxsgoto"

type Foo struct {
	Name string
}

func main() {
	// 值类型
	arr_copy := arr
	arr_copy[2] = 33
	fmt.Println("arr :", arr)
	fmt.Println("arr_copy :", arr_copy)
	fmt.Println()
	// 引用类型
	slice_arr_copy := slice_arr
	slice_arr_copy[2] = 33
	fmt.Println("slice_arr :", slice_arr)
	fmt.Println("slice_arr_copy :", slice_arr_copy)
	
	// 内存说明
	value := returnvalue()
	pointer := returnpointer()
	fmt.Printf("mystring value: %v, address: %p \n", mystring, &mystring)
	fmt.Printf("returnvalue: %v, address: %p \n", value, &value)
	fmt.Printf("returnpointer: %v, address: %p \n", *pointer, &pointer)
	fmt.Println("=================")
	teststruct()
	
}

func returnvalue() string {
	return mystring
}

func returnpointer() *string {
	return &mystring
}

func teststruct() {
	footest := Foo{Name:"ttxsgoto"}
	
	fmt.Printf("sourceName: %v address: %p \n", footest, &footest)
	footest.setName("xiaojiuyue")
	fmt.Printf("foo setName:%v , address: %p\n", footest, &footest)
	
	//(&footest).setNamePointer("xiaojiuyue...")
	// || 等价于
	footest.setNamePointer("xiaojiuyue......")
	fmt.Printf("foo setNamePointer:%v , address: %p\n", footest, &footest)
	
}

func (foo Foo) setName(name string) {
	fmt.Printf("setName: %v address: %p \n", foo, &foo)
	foo.Name = name
}

func (foo *Foo) setNamePointer(name string) {
	fmt.Printf("setNamePointer: %v address: %p \n", foo, &foo)
	foo.Name = name
}