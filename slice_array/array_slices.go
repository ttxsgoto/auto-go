package main

import "fmt"

func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5]
	
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	
	fmt.Println(arr1, slice1)
	
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice at %d is %d\n", i, slice1[i])
	}
	
	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
	
	slice1 = slice1[0:4]
	fmt.Println(slice1)
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
		
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
	
	var arr = [5]int{0, 1, 2, 3, 4}
	fmt.Println(sum(arr[:]))
	
}

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

