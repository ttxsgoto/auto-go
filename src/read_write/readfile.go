package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func main()  {
	inputfile, error := os.Open("/Users/mac/PycharmProjects/ygomi_go/src/read_write/2019-05-02_T_16-07-24.367_UTC.DGPS.gps.bak")
	
	if error != nil {
		fmt.Println("read file error")
		return
	}
	defer inputfile.Close()
	
	inputread := bufio.NewReader(inputfile)
	//fmt.Println(len(inputread))
	for {
		//inputstring, isPrefix, error := inputread.ReadLine()
		//inputstring, error := inputread.ReadString('\n')
		inputstring, error := inputread.ReadBytes('\n')
		if error == io.EOF {
			return
		}
		fmt.Println("line-->", string(inputstring))
	}
}
