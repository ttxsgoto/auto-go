package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	int
	float32
	innerS
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 3.2
	outer.int = 60
	outer.float32 = 3.4
	outer.in1 = 24
	outer.in2 = 25
	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.float32 is: %d\n", outer.float32)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)
	
	outer2 := outerS{6, 3.4, 79, 3.4, innerS{2, 3}}
	fmt.Println("Outer2 is: ", outer2)
}
