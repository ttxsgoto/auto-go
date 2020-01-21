package main

import (
	"fmt"
	"time"
)

func write1(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("channel ", i)
	}
	close(ch)
}

func main() {
	ch := make(chan int, 2)
	go write1(ch)
	
	time.Sleep(2 * time.Second)
	
	for v := range ch {
		fmt.Println("read ch :", v)
		time.Sleep(2 * time.Second)
	}
	
}
