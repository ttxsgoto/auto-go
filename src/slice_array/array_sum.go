package main

import "fmt"

func main() {
	array := [3] float64{4.4, 3.5, 5.6}
	x := Sum(&array)
	
	fmt.Printf("The sum of the array is: %f", x)
	
}

func Sum(a *[3] float64) (sum float64) {
	// 接收也为指针地址
	for _, v := range a {
		sum += v
	}
	fmt.Println(&a, *a, a)
	return
}
