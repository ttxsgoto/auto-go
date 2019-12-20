package main

import "fmt"

func main() {
	ifelse()
	switch_case()
	for1()
	for2()
	for3()
	forbreak()
	forcontinue()
	forlable()

}

func ifelse() {
	bool1 := true
	if bool1 {
		fmt.Printf("The value is true\n")
	} else {
		fmt.Printf("The value is false\n")
	}

}

func switch_case() {
	var num1 int = 100

	switch num1 {
	case 98, 99:
		fmt.Println("It's equal to 98")
	case 100:
		fmt.Println("It's equal to 100")
		fallthrough
	case 101:
		fmt.Println("It's equal to 101")
	default:
		fmt.Println("It's not equal to 98 or 100")
	}
}

func for1() {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}

}

func for2() {
	var i int = 2
	for i >= 1 {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
	}
}

func for3() {
	str := "Go is a beautiful language!"
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}

}

func forbreak() {
	var i int = 3
	for {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
		if i < 0 {
			break
		}
	}
}

func forcontinue() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		print(i)
		print(" ")
	}

}

func forlable() {
	LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

}