package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

type Rectangle struct {
	length, width float32
}

func (self *Square) Area() float32 {
	return self.side * self.side
}

func (self Rectangle) Area() float32 {
	return self.length * self.width
	
}

func main() {
	sq1 := new(Square)
	sq1.side = 5
	
	r := Rectangle{5, 3}
	q := &Square{5}
	
	shapes := []Shaper{r, q}
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
	
	var areaIntf Shaper
	
	areaIntf = sq1  // 将一个 Square 类型的变量赋值给一个接口类型的变量
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
	
}











