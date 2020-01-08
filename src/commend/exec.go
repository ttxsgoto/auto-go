package main

import (
	"os"
	"fmt"
	"os/exec"
)

func main() {
	env := os.Environ()
	fmt.Println(env)
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	pid, err := os.StartProcess("ls", []string{"ls", "-l"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err)  //
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", pid)
	cmd := exec.Command("pwd")  // this opens a gedit-window
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error %v executing command!", err)
		os.Exit(1)
	}
	fmt.Printf("The command is %v", cmd)
	
}

