package main

// 定义方法

import "fmt"

type TwoInts struct {
	a int
	b int
}

type IntVector []int

func (self *TwoInts) AddItem() int {
	return self.a + self.b
	
}

func (self *TwoInts) AddToNum(param int) int {
	return self.a + self.b + param
	
}

func (self IntVector) Sum() (s int) {
	for _, item := range self {
		s += item
	}
	return
	
}

func main() {
	two1 := new(TwoInts)
	two1.a = 10
	two1.b = 12
	sum := two1.AddItem()
	sumnum := two1.AddToNum(34)
	fmt.Printf("The sum is %d %d\n", sum, sumnum)
	
	fmt.Println(IntVector{1, 2, 3}.Sum())
}