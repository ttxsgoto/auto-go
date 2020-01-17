package main

import "fmt"
/*
	- 定义方法传递值类型，没有指定为指针类型， 在方法调用时，也可以传递指针过去
	  因为指针类型可以访问其关联的值类型的方法，如果定义的指针类型，则不能用值传递过去
	  指针传递可以使用&xxx{} or new(xxx)结果一样
	- Go中的所有内容都通过值传递，每次调用函数时，传递给它的数据都会被复制。 对于带有
	  值接收器的方法，在调用该方法时将复制该值
	- 指针类型可以调用其关联的值类型的方法，反之亦然
	- 一切都是通过值传递的，甚至是方法的接收者
	- https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
	
 */
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
	
	var areaIntf Shaper // 接口是一种类型
	
	fmt.Println("areaIntf: ", areaIntf)
	fmt.Printf("areaIntf type of %T\n", areaIntf)
	
	areaIntf = sq1  // 将一个 Square 类型的变量赋值给一个接口类型的变量
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
	
	// 接口类型值  静态类型和动态类型
	// 接口类型的静态类型就是接口本身,接口没有静态值,它指向的是动态值
	// 接口类型的变量存的是实现接口的类型的值，该值就是接口的动态值，实现接口的类型就是接口的动态类型
	Dointerface()
	
	// 空接口,一个不包含任何方法的接口，称之为空接口,因为空接口不包含任何方法，所以任何类型都默认实现了空接口
	type Empty interface {
		
	}
	names := []string{"stanley", "david", "oscar"}
	//PrintAll(names) // error:cannot use names (type []string) as type []interface {} in argument to PrintAll
	vals := make([]interface{}, len(names))
	for i, v := range names {
		vals[i] = v
	}
	fmt.Println("names---", names)
	fmt.Println("vals---", vals)
	PrintAll(vals)
	
}

type incame interface {
	name()
}

type St1 struct {
	
}

type St2 struct {
	
}

func (St1) name() {
	
}
func (St2) name() {
	
}

func Dointerface() {
	var na incame = St1{}
	fmt.Printf("Type is %T\n", na)
	fmt.Printf("Value is %v\n", na)
	na = St2{}
	fmt.Printf("Type is %T\n", na)
	fmt.Printf("Value is %v\n", na)
	
}

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}










