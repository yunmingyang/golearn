package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("cmd name: ", os.Args[0])

	fmt.Println("parameters is: ", strings.Join(os.Args[1:], " "))
}
