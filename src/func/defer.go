package main

import "fmt"

func main() {
	function1()
	out()
}

func function1() {
	var a int = 3
	fmt.Printf("In function1 at the top\n")
	defer function2(a) // defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数
	fmt.Printf("In function1 at the bottom!\n")
	fmt.Printf("It's Over!\n")
}

func function2(num int) {
	fmt.Printf("Function2: %d Deferred until the end of the calling function!\n", num)
}

func out() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

}