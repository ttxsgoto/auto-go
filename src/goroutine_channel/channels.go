package main

import (
	"fmt"
	//"time"
)

func main() {
	ch := make(chan string, 4)
	ch <- "abc"
	
	go sendData(ch)
	getData(ch)
	// 运行时（runtime）会检查所有的协程（像本例中只有一个）是否在等待
	// 着什么东西（可从某个通道读取或者写入某个通道），这意味着程序将无法继续执行
	//time.Sleep(1e9) //等待了 1 秒让两个协程完成，如果不这样，sendData() 就没有机会输出
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
	close(ch)
}

func getData(ch1 chan string) {
	//var input string
	 //time.Sleep(2e9)
	// 方式一
	for {
		input, ok := <-ch1   // 判断是否成功接收信道发送的数据
		if ok==false {
			break
		}
		fmt.Printf("%s ", input)
	}
	// 方式二
	for v := range ch1 {
		fmt.Printf("%s ", v)
	}
}