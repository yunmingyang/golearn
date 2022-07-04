// Since there should be only one main function is in the package and
// one package in a directory, define this function and pass it to the main function,
// to avoid the error
package main

import (
	"fmt"
	"os"
)

func main() {
	// base on the mechanism which is defualt initialization to ""
	// var s = "" not use too much
	// var s string = "" when the variable type is the same with the initial type
	// the "string" is excess, but if they are not the same, the type is necessary
	// We use " var s string " and " s := "" "(only in function). If the initial type is important,
	// specify the type of variable explicitly
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
