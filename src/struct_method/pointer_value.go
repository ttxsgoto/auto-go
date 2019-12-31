package main

import "fmt"

type B struct {
	thing int
}

func (self *B) change() {
	self.thing = 1
}

func (self B) write() string {
	//self.thing = 45
	return fmt.Sprint(self)
	
}

func main() {
	var b1 B    // b1 值
	b1.change()
	fmt.Println(b1.write())
	
	b2 := new(B)    // 指针
	b2.change()
	fmt.Println(b2.write())
	
}
