package main

import "fmt"

type Desc interface {
	Describe1()
}
type Person struct {
	name string
	age  int
}

type St string

func (s St) Describe1() {
	fmt.Println("s is ", s)
}

func (p Person) Describe1() {
	fmt.Printf("%s is %d years.", p.name, p.age)
	
}

func findtype(i interface{}) {
	fmt.Println()
	
	switch v := i.(type) {
	case string:
		fmt.Println("string....")
	case Desc:
		fmt.Println("Desc ", v)
		v.Describe1()
	default:
		fmt.Println("default ", v)
		fmt.Printf("Unknow type\n")
		
	}
}
func main() {
	findtype("ttxs")
	p := Person{
		name:"ttxs",
		age:25,
	}
	findtype(p)
	fmt.Println()
	st := St("xiaojiuyue")
	findtype(st)
	
}
