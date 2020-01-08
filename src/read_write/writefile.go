package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	outputfile, error := os.OpenFile("/Users/mac/PycharmProjects/ygomi_go/src/read_write/test.log", os.O_WRONLY | os.O_CREATE| os.O_TRUNC, 0666)
	
	if error != nil {
		fmt.Println("read file error")
		return
	}
	defer outputfile.Close()
	
	outputwriter := bufio.NewWriter(outputfile)
	//fmt.Println(len(inputread))
	outstring := "hello world\n"
	fmt.Fprintf(outputwriter, "Some test data.\n")
	for i :=0; i<10; i++ {
		outputwriter.WriteString(outstring) // 直接将内容写入文件,不使用缓冲
	}
	
	outputwriter.Flush()
}
