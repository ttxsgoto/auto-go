package main

import (
	"./pack1"
	"fmt"
)

func main() {
	var test1 string
	test1 = xxx.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", xxx.Pack1Int)
	// fmt.Printf("Float from package1: %f\n", pack1.pack1Float)
	
}


