package main

import (
	"fmt"
)

var ap *int

func main() {
	var i1 = 5
	fmt.Printf("An interger: %d, its location in memory:%p\n", i1, &i1)
	var intP *int
	intP = &i1
	// *intP 指针的值
	// intP  指针地址
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
	
	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p) // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s) // prints same string
	//  *p 赋另一个值来更改“对象”，这样 s 也会随之更改
	
	
}

