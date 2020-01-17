package main

import "fmt"

type struct1 struct {
	i1  int
	f1  float32
	str string
}

type person struct {
	name   string
	age    byte
	isDead bool
}
func main() {
	//type myStruct struct {
	//	i int
	//}
	//var v myStruct
	//var p *myStruct
	//*p.i = 10
	//v.i = 2
	//fmt.Println(v.i, p.i)
	
	
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 1.23
	ms.str = "abc"
	
	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Println(*ms, ms)
	p1 := person{name: "zzy", age: 100}
	p2 := p1
	p2.name = "changed"
	fmt.Println(p1.name, p2.name)
	
}
