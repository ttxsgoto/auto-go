package main

import (
	"os/exec"
	"fmt"
	"os"
	"strings"
	"github.com/dghubble/gologin"
)

func main() {
	var (
		date1, ls []byte
		err error
		cmd *exec.Cmd
	)
	gologin.DebugOnlyCookieConfig
	
	cmd = exec.Command("date")
	if date1, err = cmd.Output(); err != nil {
		fmt.Println(err, date1)
		os.Exit(1)
	}
	fmt.Println(string(date1))
	cmd = exec.Command("ls", "-l")
	if ls, err = cmd.Output(); err != nil {
		fmt.Println(err, ls)
		os.Exit(1)
	}
	fmt.Println(string(ls))
	fmt.Println(strings.Trim(string(ls), "\n"))
}
