package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[0:] {
		fmt.Println("index is: ", index, "arg is: ", arg)
	}
}
