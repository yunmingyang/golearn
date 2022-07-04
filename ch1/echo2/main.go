// Since there should be only one main function is in the package and
// one package in a directory, define this function and pass it to the main function,
// to avoid the error
package main

import (
	"fmt"
	"os"
)

func main() {
	// most simple, but only function inside, can not use for package variable
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
