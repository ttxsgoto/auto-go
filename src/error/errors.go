package main

import (
	"errors"
	"fmt"
)

var errNotFound error = errors.New("not found error")

func main() {
	fmt.Printf("error: %v\n", errNotFound)
	fmt.Printf("error: %s\n", errNotFound)
	fmt.Println("Starting the program")
	panic("A severe error occurred: stopping the program!")
	fmt.Println("Ending the program")
}
