package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//TODO: know how the strings implement that
	fmt.Println(strings.Join(os.Args[1:], " "))
}
