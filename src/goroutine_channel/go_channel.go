package main

import (
	"time"
	"fmt"
)

func write(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Println("put data", i)
	}
}

func read(ch chan int) {
	for {
		//var b int
		//b = <-ch
		//fmt.Println("xxxxx", b)
		fmt.Println("---->", <-ch)
		time.Sleep(time.Second)
	}
	
}
func main() { // 主协程
	intChan := make(chan int, 10)
	go write(intChan)   // 启动新的一个协程
	go read(intChan)
	
	time.Sleep(10 * time.Second)
	
}
