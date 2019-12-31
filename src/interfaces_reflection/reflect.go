package main

import (
	"reflect"
	"fmt"
)
// 反射 可以在运行时动态获取变量的相关信息

/*
- reflect.TypeOf  获取变量的类型
- reflect.ValueOf 获取变量的值
- reflect.Value.Kind 获取变量的类别
- reflect.Value.Interface() 转换成interface{}类型
 */
type student struct {
	Name  string
	Age   int
	Score float32
}

func test(b interface{}) {
	t := reflect.TypeOf(b)
	fmt.Println(t)
	v := reflect.ValueOf(b)
	k := v.Kind()
	fmt.Println(v, k)
	iv := v.Interface()
	stu, ok := iv.(student)
	if ok {
		fmt.Printf("%v %T", stu, stu)
	}
}
func testInt(b interface{}) {
	var1 := reflect.ValueOf(b)
	var1.Elem().SetInt(100) // 操作指针，指针变量前加*， 用来获取指针指向的变量
	c := var1.Elem().Int()
	fmt.Println("%v %T", c, c)
	
}

func Teststruct(a interface{}) {
	// 分析结构体操作
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("Expace struct")
		return
	}
	num := val.NumField()   // 字段
	numeth := val.NumMethod()   // 方法
	fmt.Println(num, numeth)
	
}
func main() {
	var b int = 200
	var a student = student{
		Name:"ttxsgoto",
		Age:20,
		Score:34.3,
	}
	test(a)
	testInt(&b)
	fmt.Println(b)
	Teststruct(a)
	
}


