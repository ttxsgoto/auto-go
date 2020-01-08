package main

import (
	"fmt"
	"time"
)

func cale(taskChan chan int, resultChan chan int) (resultChan1 chan int) {
	for v := range taskChan {
		flag := true
		for i := 2; i < v; i++ {
			if (v % i == 0) {
				flag = false
				break
			}
		}
		if flag {
			resultChan <- v
		}
	}
	return resultChan
}

func main() {
	intChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)
	go func() {
		for i := 0; i < 10000; i++ {
			intChan <- i
		}
		close(intChan)
	}()
	
	for i := 0; i < 8; i++ {
		go cale(intChan, resultChan)
	}
	go func() {
		for v := range resultChan {
			fmt.Println("--->", v)
		}
	}()
	//for v := range resultChan {
	//	fmt.Println("--->", v)
	//}
	
	time.Sleep(2 * time.Second)
}