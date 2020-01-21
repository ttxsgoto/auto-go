package main

import (
	"sync"
	"fmt"
	"time"
)

// WaitGroup 用于等待一批 Go 协程执行结束。程序控制会一直阻塞，直到这些协程全部执行完毕
// 工作池,通常为一组线程,它们正等待分配给它们的任务,一旦完成分配的任务，他们就会再次为下一个任务提供服务


func process(i int, wg *sync.WaitGroup)  {
	fmt.Println("start Goroutine, ", i)
	time.Sleep(time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
	
}

func main()  {
	
	num :=3
	var wg sync.WaitGroup
	for i:=0;i<num;i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	
	wg.Wait()
	fmt.Println("all done.")
}


