package main

import (
	"bufio"
	"os"
	"io"
	"fmt"
	"flag"
)

func cat(r *bufio.Reader) {
	for {
		buf, error := r.ReadBytes('\n')
		if error == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
	
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
	
}
