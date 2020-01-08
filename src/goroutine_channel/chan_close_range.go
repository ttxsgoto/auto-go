package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)   // 从第11个数据开始， 都是为0， 已经关掉channel了
	
	for v := range ch {
		fmt.Println(v)
	}
	//for {
	//	var b int
	//	b, ok := <- ch  // 判断channel是否被关闭了
	//	if ok == false {
	//		fmt.Println("channel is closed")
	//		break
	//	}
	//	fmt.Println(b)
	//}
	
}

