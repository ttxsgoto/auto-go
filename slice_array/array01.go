package main

import "fmt"

func main() {
	//a := [...]string{"a", "b", "c", "d", "e"}
	//for _, i := range a {
	//	fmt.Println("Array item, ", i)
	//}
	var ar [3] int
	ar[0] = 1
	ar[1] =2
	arrAge := [5]int{18, 20, 15, 22, 16}
	var arrLazy = [...]int{5, 6, 7, 8, 22}
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	
	f(ar)
	fp(&ar)
	fmt.Println(ar)
	fmt.Println(arrAge)
	fmt.Println(arrLazy)
	fmt.Println(arrKeyValue)
	
}

func f(a [3]int) {
	a[1] =1
	fmt.Println(a)
}
func fp(a *[3]int) {
	//&a[1] = 3
	fmt.Println(a)
	
}