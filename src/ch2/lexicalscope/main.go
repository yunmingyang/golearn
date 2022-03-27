package main

import (
	"fmt"
	"os"
)

var cwd string
var p = &cwd

func init() {
	// Print global variable address
	fmt.Println(&cwd)
	// No extra symbol, such as * or &, it can print the address directly, that the pointer point
	fmt.Println(p)
}


func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Getwd failed with: %d", err)
		os.Exit(1)
	}
	fmt.Println(cwd, "and its address is: ", &cwd)
}