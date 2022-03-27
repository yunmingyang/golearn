package main

import "fmt"




func main() {
	f := func(i int) int {
		return i
	}

	fmt.Println(f(3))
}